//+build acme_dns_internal

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/internal"
)

func init() {
	providers["internal"] = func() (challenge.Provider, error) {
		return internal.NewDNSProvider()
	}
}
