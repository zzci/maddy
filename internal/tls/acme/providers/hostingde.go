//+build acme_dns_hostingde

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/hostingde"
)

func init() {
	providers["hostingde"] = func() (challenge.Provider, error) {
		return hostingde.NewDNSProvider()
	}
}
