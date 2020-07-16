//+build acme_dns_oraclecloud

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/oraclecloud"
)

func init() {
	providers["oraclecloud"] = func() (challenge.Provider, error) {
		return oraclecloud.NewDNSProvider()
	}
}
