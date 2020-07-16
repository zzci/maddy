//+build acme_dns_dreamhost

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/dreamhost"
)

func init() {
	providers["dreamhost"] = func() (challenge.Provider, error) {
		return dreamhost.NewDNSProvider()
	}
}
