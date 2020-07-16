//+build acme_dns_cloudflare

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/cloudflare"
)

func init() {
	providers["cloudflare"] = func() (challenge.Provider, error) {
		return cloudflare.NewDNSProvider()
	}
}
