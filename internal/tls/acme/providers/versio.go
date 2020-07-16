//+build acme_dns_versio

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/versio"
)

func init() {
	providers["versio"] = func() (challenge.Provider, error) {
		return versio.NewDNSProvider()
	}
}
