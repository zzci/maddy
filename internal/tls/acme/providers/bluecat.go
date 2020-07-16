//+build acme_dns_bluecat

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/bluecat"
)

func init() {
	providers["bluecat"] = func() (challenge.Provider, error) {
		return bluecat.NewDNSProvider()
	}
}
