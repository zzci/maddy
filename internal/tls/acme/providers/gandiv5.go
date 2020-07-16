//+build acme_dns_gandiv5

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/gandiv5"
)

func init() {
	providers["gandiv5"] = func() (challenge.Provider, error) {
		return gandiv5.NewDNSProvider()
	}
}
