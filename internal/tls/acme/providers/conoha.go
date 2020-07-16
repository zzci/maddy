//+build acme_dns_conoha

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/conoha"
)

func init() {
	providers["conoha"] = func() (challenge.Provider, error) {
		return conoha.NewDNSProvider()
	}
}
