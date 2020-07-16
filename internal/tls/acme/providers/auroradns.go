//+build acme_dns_auroradns

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/auroradns"
)

func init() {
	providers["auroradns"] = func() (challenge.Provider, error) {
		return auroradns.NewDNSProvider()
	}
}
