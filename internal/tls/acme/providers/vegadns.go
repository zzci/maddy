//+build acme_dns_vegadns

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/vegadns"
)

func init() {
	providers["vegadns"] = func() (challenge.Provider, error) {
		return vegadns.NewDNSProvider()
	}
}
