//+build acme_dns_namesilo

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/namesilo"
)

func init() {
	providers["namesilo"] = func() (challenge.Provider, error) {
		return namesilo.NewDNSProvider()
	}
}
