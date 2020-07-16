//+build acme_dns_transip

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/transip"
)

func init() {
	providers["transip"] = func() (challenge.Provider, error) {
		return transip.NewDNSProvider()
	}
}
