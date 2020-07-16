//+build acme_dns_glesys

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/glesys"
)

func init() {
	providers["glesys"] = func() (challenge.Provider, error) {
		return glesys.NewDNSProvider()
	}
}
