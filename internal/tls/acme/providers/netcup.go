//+build acme_dns_netcup

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/netcup"
)

func init() {
	providers["netcup"] = func() (challenge.Provider, error) {
		return netcup.NewDNSProvider()
	}
}
