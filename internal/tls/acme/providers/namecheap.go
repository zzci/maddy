//+build acme_dns_namecheap

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/namecheap"
)

func init() {
	providers["namecheap"] = func() (challenge.Provider, error) {
		return namecheap.NewDNSProvider()
	}
}
