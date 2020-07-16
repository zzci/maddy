//+build acme_dns_vultr

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/vultr"
)

func init() {
	providers["vultr"] = func() (challenge.Provider, error) {
		return vultr.NewDNSProvider()
	}
}
