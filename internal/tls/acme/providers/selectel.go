//+build acme_dns_selectel

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/selectel"
)

func init() {
	providers["selectel"] = func() (challenge.Provider, error) {
		return selectel.NewDNSProvider()
	}
}
