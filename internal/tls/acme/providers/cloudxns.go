//+build acme_dns_cloudxns

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/cloudxns"
)

func init() {
	providers["cloudxns"] = func() (challenge.Provider, error) {
		return cloudxns.NewDNSProvider()
	}
}
