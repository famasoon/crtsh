package parser

import (
	"crypto/x509"
	"encoding/pem"
	"testing"
)

const examplePEM = `
-----BEGIN CERTIFICATE-----
MIIF6DCCBNCgAwIBAgIQBBHej1O0YvalqGG3EuxrWTANBgkqhkiG9w0BAQsFADBw
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMS8wLQYDVQQDEyZEaWdpQ2VydCBTSEEyIEhpZ2ggQXNz
dXJhbmNlIFNlcnZlciBDQTAeFw0xNDExMDYwMDAwMDBaFw0xNTExMTMxMjAwMDBa
MIGlMQswCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEUMBIGA1UEBxML
TG9zIEFuZ2VsZXMxPDA6BgNVBAoTM0ludGVybmV0IENvcnBvcmF0aW9uIGZvciBB
c3NpZ25lZCBOYW1lcyBhbmQgTnVtYmVyczETMBEGA1UECxMKVGVjaG5vbG9neTEY
MBYGA1UEAxMPd3d3LmV4YW1wbGUub3JnMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A
MIIBCgKCAQEAnmY/UqPRjLZ83+1UdAik5H5ANlOJiNonmNo7ZlX3JA1pPtHLP+bW
rTqeZX/276hrg7DK0k5dMf8r9w7Dt4shPxtL9hvcZpy7wH1nFUEoypKps8u0ITqD
b7gj3dTXzASRgxTSXwYIb6mXC6F+NXzKm0WMJ+txdgq5Xj+byJiuiQUK5NCbovfk
JZ2f8eByppcbGDVai55TZww9Xb29KD+Tp2TnGzpBQMoHRgkMCFEOLiEHjX0HhEv5
wDhltTGgvy7nZrxAH2RRxaHm9vtdXB1ql6Cr6RrosC6JJB4HNTkJzNW0HEbeIHwG
gB4I8gcTYDgn8q4+aM8V74gdfgYI9wdC4wIDAQABo4ICRjCCAkIwHwYDVR0jBBgw
FoAUUWj/kK8CB3U8zNllZGKiErhZcjswHQYDVR0OBBYEFLAAp/Qi6bHOIWEXxMRu
cWTI5gxVMIGBBgNVHREEejB4gg93d3cuZXhhbXBsZS5vcmeCC2V4YW1wbGUuY29t
ggtleGFtcGxlLmVkdYILZXhhbXBsZS5uZXSCC2V4YW1wbGUub3Jngg93d3cuZXhh
bXBsZS5jb22CD3d3dy5leGFtcGxlLmVkdYIPd3d3LmV4YW1wbGUubmV0MA4GA1Ud
DwEB/wQEAwIFoDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwdQYDVR0f
BG4wbDA0oDKgMIYuaHR0cDovL2NybDMuZGlnaWNlcnQuY29tL3NoYTItaGEtc2Vy
dmVyLWczLmNybDA0oDKgMIYuaHR0cDovL2NybDQuZGlnaWNlcnQuY29tL3NoYTIt
aGEtc2VydmVyLWczLmNybDBCBgNVHSAEOzA5MDcGCWCGSAGG/WwBATAqMCgGCCsG
AQUFBwIBFhxodHRwczovL3d3dy5kaWdpY2VydC5jb20vQ1BTMIGDBggrBgEFBQcB
AQR3MHUwJAYIKwYBBQUHMAGGGGh0dHA6Ly9vY3NwLmRpZ2ljZXJ0LmNvbTBNBggr
BgEFBQcwAoZBaHR0cDovL2NhY2VydHMuZGlnaWNlcnQuY29tL0RpZ2lDZXJ0U0hB
MkhpZ2hBc3N1cmFuY2VTZXJ2ZXJDQS5jcnQwDAYDVR0TAQH/BAIwADANBgkqhkiG
9w0BAQsFAAOCAQEAXqwhJN7bOXiob/NghAastULTy1TLg/rNY67IgUTWob8V2/Hy
FcSnPiQeWCNly6nqUN0wZUFlOzUTrxoHVsGycg6NESs0+2cYHvrZxGCb3GcPsCX6
bm1CGIFhsCbPMImgg2nC82CfyEvMNHkUDBki7eQwyo26wrKjzayzBboV3HNhxMOl
5tqpnLRGyyIbKAeKepRO+6cNlvMawUPZWbzNL9UOMMMl6iYk+2ttvpNE288TO/vV
tOiS1jXb8xWWRRZyxrZbpaybPN3qkrNdqxBlyuPIy2u0UKYuovcup8a9x7ZfoJsB
I5JUNzQIPHaH0kP40DdTBNmczS4UiWaoY3pnlw==
-----END CERTIFICATE-----`

func TestEnumDNS(t *testing.T) {
	var tests = []string{
		"www.example.org",
		"example.com",
		"example.edu",
		"example.net",
		"example.org",
		"www.example.com",
		"www.example.edu",
		"www.example.net",
	}

	block, _ := pem.Decode([]byte(examplePEM))
	if block == nil {
		t.Fatal("Can not read pem file")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		t.Fatal("Can not parse certficate")
	}

	for key, record := range EnumDNS(cert) {
		if record != tests[key] {
			t.Errorf("#%d: got: %s want %s", key, record, tests[key])
		}
	}
}
