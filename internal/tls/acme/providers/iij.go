//+build acme_dns_iij

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/iij"
)

func init() {
	providers["iij"] = func() (challenge.Provider, error) {
		return iij.NewDNSProvider()
	}
}
