//+build acme_dns_stackpath

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/stackpath"
)

func init() {
	providers["stackpath"] = func() (challenge.Provider, error) {
		return stackpath.NewDNSProvider()
	}
}
