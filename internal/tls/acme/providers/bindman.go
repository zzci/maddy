//+build acme_dns_bindman

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/bindman"
)

func init() {
	providers["bindman"] = func() (challenge.Provider, error) {
		return bindman.NewDNSProvider()
	}
}
