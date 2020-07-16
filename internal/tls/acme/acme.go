package acme

import (
	"context"
	"crypto/tls"
	"fmt"
	"path/filepath"
	"sync"

	"github.com/caddyserver/certmagic"
	"github.com/foxcpp/maddy/framework/config"
	"github.com/foxcpp/maddy/framework/hooks"
	"github.com/foxcpp/maddy/framework/module"
	"github.com/go-acme/lego/v3/challenge"
)

var (
	magicInit sync.Once

	storage certmagic.Storage
	cache   *certmagic.Cache

	providers = make(map[string]func() (challenge.Provider, error))
)

type Loader struct {
	instName    string
	dnsProvider string
	serverNames []string

	magic   *certmagic.Config
	manager *certmagic.ACMEManager

	cancelACME context.CancelFunc
}

func New(_, instName string, _, inlineArgs []string) (module.Module, error) {
	l := &Loader{
		instName: instName,
	}

	if len(inlineArgs) > 0 {
		l.dnsProvider = inlineArgs[0]
		l.serverNames = inlineArgs[1:]
	}

	return l, nil
}

func (f *Loader) Init(cfg *config.Map) error {
	mgr := certmagic.ACMEManager{}
	cfg.StringList("names", false, false, f.serverNames, &f.serverNames)
	cfg.String("ca", false, false, certmagic.LetsEncryptProductionCA, &mgr.CA)
	cfg.String("email", false, false, "", &mgr.Email)
	cfg.Bool("agreed", false, false, &mgr.Agreed)
	cfg.String("dns_provider", false, false, f.dnsProvider, &f.dnsProvider)
	if _, err := cfg.Process(); err != nil {
		return err
	}

	if len(f.serverNames) == 0 {
		return fmt.Errorf("tls.loader.acme: at least one domain name required")
	}

	provFunc := providers[f.dnsProvider]
	if provFunc == nil {
		return fmt.Errorf("tls.loader.acme: unknown DNS provider: %s", f.dnsProvider)
	}
	var err error
	mgr.DNSProvider, err = provFunc()
	if err != nil {
		return err
	}

	mgr.DisableHTTPChallenge = true
	mgr.DisableTLSALPNChallenge = true

	magicInit.Do(func() {
		storage = &certmagic.FileStorage{Path: filepath.Join(config.StateDirectory, "acme")}
		cache = certmagic.NewCache(certmagic.CacheOptions{
			GetConfigForCert: func(_ certmagic.Certificate) (*certmagic.Config, error) {
				return &certmagic.Config{
					Storage: storage,
				}, nil
			},
		})

		hooks.AddHook(hooks.EventShutdown, certmagic.CleanUpOwnLocks)
	})
	f.magic = certmagic.New(cache, certmagic.Config{
		Storage: storage,
	})
	f.manager = certmagic.NewACMEManager(f.magic, mgr)
	f.magic.Issuer = f.manager

	asyncCtx, cancelACME := context.WithCancel(context.Background())
	f.cancelACME = cancelACME
	if err := f.magic.ManageAsync(asyncCtx, f.serverNames); err != nil {
		return err
	}

	return nil
}

func (f *Loader) Close() error {
	f.cancelACME()
	return nil
}

func (f *Loader) Name() string {
	return "tls.loader.self_signed"
}

func (f *Loader) InstanceName() string {
	return f.instName
}

func (f *Loader) ConfigureTLS(cfg *tls.Config) error {
	cfg.GetCertificate = f.magic.GetCertificate
	return nil
}

func init() {
	module.Register("tls.loader.acme", New)
}
