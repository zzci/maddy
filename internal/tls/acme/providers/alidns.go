//+build acme_dns_alidns

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/alidns"
)

func init() {
	providers["alidns"] = func() (challenge.Provider, error) {
		return alidns.NewDNSProvider()
	}
}
