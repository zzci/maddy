//+build acme_dns_fixtures

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/fixtures"
)

func init() {
	providers["fixtures"] = func() (challenge.Provider, error) {
		return fixtures.NewDNSProvider()
	}
}
