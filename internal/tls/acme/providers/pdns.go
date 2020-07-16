//+build acme_dns_pdns

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/pdns"
)

func init() {
	providers["pdns"] = func() (challenge.Provider, error) {
		return pdns.NewDNSProvider()
	}
}
