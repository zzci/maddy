//+build acme_dns_dnsmadeeasy

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/dnsmadeeasy"
)

func init() {
	providers["dnsmadeeasy"] = func() (challenge.Provider, error) {
		return dnsmadeeasy.NewDNSProvider()
	}
}
