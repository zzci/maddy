//+build acme_dns_zoneee

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/zoneee"
)

func init() {
	providers["zoneee"] = func() (challenge.Provider, error) {
		return zoneee.NewDNSProvider()
	}
}
