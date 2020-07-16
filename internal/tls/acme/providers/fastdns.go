//+build acme_dns_fastdns

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/fastdns"
)

func init() {
	providers["fastdns"] = func() (challenge.Provider, error) {
		return fastdns.NewDNSProvider()
	}
}
