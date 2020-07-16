//+build acme_dns_gcloud

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/gcloud"
)

func init() {
	providers["gcloud"] = func() (challenge.Provider, error) {
		return gcloud.NewDNSProvider()
	}
}
