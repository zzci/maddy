package module

import (
	"crypto/tls"
)

// TLSLoader interface is module interface that can be used to supply TLS
// certificates to TLS-enabled endpoints.
//
// Note that loading function may be called for each connection - it is
// highly recommended to cache parsed form of certificate.
//
// Function ConfigureTLS should do necessary changes to the config object
// to make it use provided certificates. It may change GetCertificates
// or Certificates.
//
// Modules implementing this interface should be registered with prefix
// "tls.loader." in name.
type TLSLoader interface {
	ConfigureTLS(cfg *tls.Config) error
}
