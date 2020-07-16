//+build acme_dns_namedotcom

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/namedotcom"
)

func init() {
	providers["namedotcom"] = func() (challenge.Provider, error) {
		return namedotcom.NewDNSProvider()
	}
}
