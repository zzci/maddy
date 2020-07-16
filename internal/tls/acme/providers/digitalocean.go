//+build acme_dns_digitalocean

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/digitalocean"
)

func init() {
	providers["digitalocean"] = func() (challenge.Provider, error) {
		return digitalocean.NewDNSProvider()
	}
}
