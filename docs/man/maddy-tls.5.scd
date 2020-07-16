maddy-tls(5) "maddy mail server" "maddy reference documentation"

; TITLE Advanced TLS configuration

# TLS server configuration

TLS certificates are obtained by modules called "certificate loaders". 'tls' directive
arguments specify name of loader to use and arguments. Due to syntax limitations
advanced configuration for loader should be specified using 'loader' directive, see
below.

```
tls file cert.pem key.pem {
	protocols tls1.2 tls1.3
	curve X25519
	ciphers ...
}

tls {
	loader file cert.pem key.pem {
		# Options for loader go here.
	}
	protocols tls1.2 tls1.3
	curve X25519
	ciphers ...
}
```

## Available certificate loaders

- file

	Accepts argument pairs specifying certificate and then key.
	E.g. 'tls file certA.pem keyA.pem certB.pem keyB.pem'

	If multiple certificates are listed, SNI will be used.

- acme (EXPERIMENTAL)

	Obtains TLS certificates using ACME protocol with dns-01 challenge.
	Requires additional configuration (see below).

- off

	Not really a loader but a special value for tls directive, explicitly disables TLS for
	endpoint(s).

## Advanced TLS configuration

*Note: maddy uses secure defaults and TLS handshake is resistant to active downgrade attacks.*
*There is no need to change anything in most cases.*

*Syntax*: ++
    protocols _min_version_ _max_version_ ++
    protocols _version_ ++
*Default*: tls1.0 tls1.3

Minimum/maximum accepted TLS version. If only one value is specified, it will
be the only one usable version.

Valid values are: tls1.0, tls1.1, tls1.2, tls1.3

*Syntax*: ciphers _ciphers..._ ++
*Default*: Go version-defined set of 'secure ciphers', ordered by hardware
performance

List of supported cipher suites, in preference order. Not used with TLS 1.3.

Valid values:

- RSA-WITH-RC4128-SHA
- RSA-WITH-3DES-EDE-CBC-SHA
- RSA-WITH-AES128-CBC-SHA
- RSA-WITH-AES256-CBC-SHA
- RSA-WITH-AES128-CBC-SHA256
- RSA-WITH-AES128-GCM-SHA256
- RSA-WITH-AES256-GCM-SHA384
- ECDHE-ECDSA-WITH-RC4128-SHA
- ECDHE-ECDSA-WITH-AES128-CBC-SHA
- ECDHE-ECDSA-WITH-AES256-CBC-SHA
- ECDHE-RSA-WITH-RC4128-SHA
- ECDHE-RSA-WITH-3DES-EDE-CBC-SHA
- ECDHE-RSA-WITH-AES128-CBC-SHA
- ECDHE-RSA-WITH-AES256-CBC-SHA
- ECDHE-ECDSA-WITH-AES128-CBC-SHA256
- ECDHE-RSA-WITH-AES128-CBC-SHA256
- ECDHE-RSA-WITH-AES128-GCM-SHA256
- ECDHE-ECDSA-WITH-AES128-GCM-SHA256
- ECDHE-RSA-WITH-AES256-GCM-SHA384
- ECDHE-ECDSA-WITH-AES256-GCM-SHA384
- ECDHE-RSA-WITH-CHACHA20-POLY1305
- ECDHE-ECDSA-WITH-CHACHA20-POLY1305

*Syntax*: curve _curves..._ ++
*Default*: defined by Go version

The elliptic curves that will be used in an ECDHE handshake, in preference
order.

Valid values: p256, p384, p521, X25519.

# TLS client configuration

tls_client directive allows to customize behavior of TLS client implementation,
notably adjusting minimal and maximal TLS versions and allowed cipher suites,
enabling TLS client authentication.

```
tls_client {
    protocols tls1.2 tls1.3
    ciphers ...
    curve X25519
    root_ca /etc/ssl/cert.pem

    cert /etc/ssl/private/maddy-client.pem
    key /etc/ssl/private/maddy-client.pem
}
```

*Syntax*: ++
    protocols _min_version_ _max_version_ ++
    protocols _version_ ++
*Default*: tls1.0 tls1.3

Minimum/maximum accepted TLS version. If only one value is specified, it will
be the only one usable version.

Valid values are: tls1.0, tls1.1, tls1.2, tls1.3

*Syntax*: ciphers _ciphers..._ ++
*Default*: Go version-defined set of 'secure ciphers', ordered by hardware
performance

List of supported cipher suites, in preference order. Not used with TLS 1.3.

See TLS server configuration for list of supported values.

*Syntax*: curve _curves..._ ++
*Default*: defined by Go version

The elliptic curves that will be used in an ECDHE handshake, in preference
order.

Valid values: p256, p384, p521, X25519.

*Syntax*: root_ca _paths..._ ++
*Default*: system CA pool

List of files with PEM-encoded CA certificates to use when verifying
server certificates.

*Syntax*: ++
    cert _cert_path_ ++
    key _key_path_ ++
*Default*: not specified

Present the specified certificate when server requests a client certificate.
Files should use PEM format. Both directives should be specified.

## ACME configuration

ACME dns-01 challenge is required since maddy usually has no access to port 80
or 443 to use HTTP-based challenge. This requires configuration to permit maddy
to non-interactively edit DNS records for your MX domain.

In maddy configuration you should specify DNS provider ID and domains to
acquire certificates for. Additional per-provider configuration needs to be
provided via environment variables. See https://go-acme.github.io/lego/dns/
for needed configuration and available providers.

For all providers except for rfc2136 maddy needs to be built with
--tags 'acme_dns_PROVIDER' passed to build.sh script or go build command.
e.g. ./build.sh --tags 'acme_dns_cloudflare'

Also you will probably need to add "agreed" directive to indicate you agree
to CA ToS.

Full example:
```
tls {
	loader acme rfc2136 mx.example.org {
		agreed
	}
}
```

## ACME directives

*Syntax:* ca _url_ ++
*Default:* Let's Encrypt production endpoint

HTTPS URL of ACME CA directory endpoint.

*Syntax:* agreed _boolean_ ++
*Default:* false

Indicates agreement to CA ToS.

*Syntax:* dns_provider _name_ ++
*Default:* not set

DNS provider name if one was not provided in arguments.

*Syntax:* names _list..._ ++
*Default:* none

List of domains to obtain certificates for, if none was provided in
arguments.
