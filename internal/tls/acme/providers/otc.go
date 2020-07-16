//+build acme_dns_otc

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/otc"
)

func init() {
	providers["otc"] = func() (challenge.Provider, error) {
		return otc.NewDNSProvider()
	}
}
