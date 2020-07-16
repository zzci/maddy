//+build acme_dns_cloudns

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/cloudns"
)

func init() {
	providers["cloudns"] = func() (challenge.Provider, error) {
		return cloudns.NewDNSProvider()
	}
}
