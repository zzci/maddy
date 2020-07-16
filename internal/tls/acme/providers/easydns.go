//+build acme_dns_easydns

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/easydns"
)

func init() {
	providers["easydns"] = func() (challenge.Provider, error) {
		return easydns.NewDNSProvider()
	}
}
