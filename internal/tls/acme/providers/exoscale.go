//+build acme_dns_exoscale

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/exoscale"
)

func init() {
	providers["exoscale"] = func() (challenge.Provider, error) {
		return exoscale.NewDNSProvider()
	}
}
