//+build acme_dns_duckdns

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/duckdns"
)

func init() {
	providers["duckdns"] = func() (challenge.Provider, error) {
		return duckdns.NewDNSProvider()
	}
}
