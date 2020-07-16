//+build acme_dns_exec

package acme

import (
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/exec"
)

func init() {
	providers["exec"] = func() (challenge.Provider, error) {
		return exec.NewDNSProvider()
	}
}
