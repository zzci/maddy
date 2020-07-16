//+build acme_dns_acmedns

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/acmedns"
)

func init() {
	providers["acmedns"] = func() (challenge.Provider, error) {
		return acmedns.NewDNSProvider()
	}
}
