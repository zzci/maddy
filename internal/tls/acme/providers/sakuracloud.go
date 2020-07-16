//+build acme_dns_sakuracloud

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/sakuracloud"
)

func init() {
	providers["sakuracloud"] = func() (challenge.Provider, error) {
		return sakuracloud.NewDNSProvider()
	}
}
