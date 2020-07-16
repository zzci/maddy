//+build acme_dns_dode

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/dode"
)

func init() {
	providers["dode"] = func() (challenge.Provider, error) {
		return dode.NewDNSProvider()
	}
}
