//+build acme_dns_dyn

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/dyn"
)

func init() {
	providers["dyn"] = func() (challenge.Provider, error) {
		return dyn.NewDNSProvider()
	}
}
