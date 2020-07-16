//+build acme_dns_httpreq

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/httpreq"
)

func init() {
	providers["httpreq"] = func() (challenge.Provider, error) {
		return httpreq.NewDNSProvider()
	}
}
