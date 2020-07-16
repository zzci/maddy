//+build acme_dns_ns1

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/ns1"
)

func init() {
	providers["ns1"] = func() (challenge.Provider, error) {
		return ns1.NewDNSProvider()
	}
}
