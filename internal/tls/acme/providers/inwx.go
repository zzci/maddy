//+build acme_dns_inwx

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/inwx"
)

func init() {
	providers["inwx"] = func() (challenge.Provider, error) {
		return inwx.NewDNSProvider()
	}
}
