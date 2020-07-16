//+build acme_dns_rackspace

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/rackspace"
)

func init() {
	providers["rackspace"] = func() (challenge.Provider, error) {
		return rackspace.NewDNSProvider()
	}
}
