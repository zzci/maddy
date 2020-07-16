//+build acme_dns_ovh

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/ovh"
)

func init() {
	providers["ovh"] = func() (challenge.Provider, error) {
		return ovh.NewDNSProvider()
	}
}
