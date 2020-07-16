//+build acme_dns_linode

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/linode"
)

func init() {
	providers["linode"] = func() (challenge.Provider, error) {
		return linode.NewDNSProvider()
	}
}
