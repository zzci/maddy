//+build acme_dns_joker

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/joker"
)

func init() {
	providers["joker"] = func() (challenge.Provider, error) {
		return joker.NewDNSProvider()
	}
}
