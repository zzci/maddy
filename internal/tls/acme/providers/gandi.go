//+build acme_dns_gandi

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/gandi"
)

func init() {
	providers["gandi"] = func() (challenge.Provider, error) {
		return gandi.NewDNSProvider()
	}
}
