//+build acme_dns_azure

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/azure"
)

func init() {
	providers["azure"] = func() (challenge.Provider, error) {
		return azure.NewDNSProvider()
	}
}
