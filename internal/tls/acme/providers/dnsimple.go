//+build acme_dns_dnsimple

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/dnsimple"
)

func init() {
	providers["dnsimple"] = func() (challenge.Provider, error) {
		return dnsimple.NewDNSProvider()
	}
}
