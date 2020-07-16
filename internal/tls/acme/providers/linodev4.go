//+build acme_dns_linodev4

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/linodev4"
)

func init() {
	providers["linodev4"] = func() (challenge.Provider, error) {
		return linodev4.NewDNSProvider()
	}
}
