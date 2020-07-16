//+build acme_dns_mydnsjp

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/mydnsjp"
)

func init() {
	providers["mydnsjp"] = func() (challenge.Provider, error) {
		return mydnsjp.NewDNSProvider()
	}
}
