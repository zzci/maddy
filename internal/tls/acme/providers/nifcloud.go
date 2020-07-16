//+build acme_dns_nifcloud

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/nifcloud"
)

func init() {
	providers["nifcloud"] = func() (challenge.Provider, error) {
		return nifcloud.NewDNSProvider()
	}
}
