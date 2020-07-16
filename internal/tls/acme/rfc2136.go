package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/rfc2136"
)

func init() {
	providers["rfc2136"] = func() (challenge.Provider, error) {
		return rfc2136.NewDNSProvider()
	}
}
