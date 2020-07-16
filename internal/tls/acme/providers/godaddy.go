//+build acme_dns_godaddy

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/godaddy"
)

func init() {
	providers["godaddy"] = func() (challenge.Provider, error) {
		return godaddy.NewDNSProvider()
	}
}
