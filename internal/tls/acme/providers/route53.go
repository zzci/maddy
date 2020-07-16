//+build acme_dns_route53

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/route53"
)

func init() {
	providers["route53"] = func() (challenge.Provider, error) {
		return route53.NewDNSProvider()
	}
}
