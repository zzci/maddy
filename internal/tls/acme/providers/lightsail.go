//+build acme_dns_lightsail

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/lightsail"
)

func init() {
	providers["lightsail"] = func() (challenge.Provider, error) {
		return lightsail.NewDNSProvider()
	}
}
