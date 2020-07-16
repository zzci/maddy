//+build acme_dns_vscale

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/vscale"
)

func init() {
	providers["vscale"] = func() (challenge.Provider, error) {
		return vscale.NewDNSProvider()
	}
}
