//+build acme_dns_dnspod

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/dnspod"
)

func init() {
	providers["dnspod"] = func() (challenge.Provider, error) {
		return dnspod.NewDNSProvider()
	}
}
