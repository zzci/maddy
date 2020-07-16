//+build acme_dns_designate

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/designate"
)

func init() {
	providers["designate"] = func() (challenge.Provider, error) {
		return designate.NewDNSProvider()
	}
}
