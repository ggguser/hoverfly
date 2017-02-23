package goproxy

import (
	"crypto/tls"
	"crypto/x509"
)

func init() {
	if goproxyCaErr != nil {
		panic("Error parsing builtin CA " + goproxyCaErr.Error())
	}
	var err error
	if GoproxyCa.Leaf, err = x509.ParseCertificate(GoproxyCa.Certificate[0]); err != nil {
		panic("Error parsing builtin CA " + err.Error())
	}
}

var tlsClientSkipVerify = &tls.Config{InsecureSkipVerify: true}

var defaultTLSConfig = &tls.Config{
	InsecureSkipVerify: true,
}

var CA_CERT = []byte("-----BEGIN CERTIFICATE-----\nMIIC5DCCAcygAwIBAgIQXO4QJIKF7QigtZiEtBiwJjANBgkqhkiG9w0BAQsFADAA\nMB4XDTE2MDEyMzE3MTg0NVoXDTI2MDEyMDE3MTg0NVowADCCASIwDQYJKoZIhvcN\nAQEBBQADggEPADCCAQoCggEBANPdwufJajG29dPoETUKmyRzZiiezBXBEuG1eoCR\nGASEldPsfwSAFOtNXv5m5/bgd4mm01+u4eXvlecDPuqZWyc52U/BkiBnLuQoBo5a\nVxjGv1A6QFgElJ8SRZXd+9yxZqlyHVYH4efOoU0u62LdgQiJ99Rkry2aMny2PLKP\noSwTcHOYxSaJhFBNEqSZlOfhOsBR22y4w6VCNb3+jFWhNzaYee3oegQLhArjm3+m\nHtXJHIkkknC8lTi74jupv6RmEGadnmDZEhsKPA+kqjYzfIWjES7M04AHg6G8LpkZ\nyQxfOtALpkdF+GHblzBCa1zN5S8JRADVMmRRbGH/wTEz6WUCAwEAAaNaMFgwDgYD\nVR0PAQH/BAQDAgKkMBMGA1UdJQQMMAoGCCsGAQUFBwMBMBIGA1UdEwEB/wQIMAYB\nAf8CAQEwHQYDVR0OBBYEFPq1RmsNXdqrPcJ9uI4rPxHJWjv1MA0GCSqGSIb3DQEB\nCwUAA4IBAQBv8pqjyBLdwdow0aG67rg6Rmz6pIN0QYzMqPc19K8WTXme+H7o4BNT\nI13K8S4qT1t+uTcBwh2dUACUnq6oFbBNDH6k68ziEy5MtcbJM6CqkWlyqdLIbjwa\nmrhP0kUUk5H23xSKEOwSrDJRukBngSzmadCYRYTz7ni6/AriiP0U9CLNOQfV8SbF\nvrY07rj6MTQc6GsYKrM9/hPQaYc++xp9A33OKTzDjr4cIxIzo9nlLUEz+bRceQ+V\nwLj1tZJxGgA1Ewyc+Z5D8j+YOKSoNSAeFEjTD4zhRCJyjEtQoyPbk4ADE802o1bd\nHnH/LoLCHEWcDxthmefQCPgUII4rif/L\n-----END CERTIFICATE-----\n")

var CA_KEY = []byte("-----BEGIN RSA PRIVATE KEY-----\nMIIEpQIBAAKCAQEA093C58lqMbb10+gRNQqbJHNmKJ7MFcES4bV6gJEYBISV0+x/\nBIAU601e/mbn9uB3iabTX67h5e+V5wM+6plbJznZT8GSIGcu5CgGjlpXGMa/UDpA\nWASUnxJFld373LFmqXIdVgfh586hTS7rYt2BCIn31GSvLZoyfLY8so+hLBNwc5jF\nJomEUE0SpJmU5+E6wFHbbLjDpUI1vf6MVaE3Nph57eh6BAuECuObf6Ye1ckciSSS\ncLyVOLviO6m/pGYQZp2eYNkSGwo8D6SqNjN8haMRLszTgAeDobwumRnJDF860Aum\nR0X4YduXMEJrXM3lLwlEANUyZFFsYf/BMTPpZQIDAQABAoIBAQC6cz/SkhvFspj2\nqxVxk2rjEjeGafF695YxUm+Dc60qVLAyd790a299gHKn+lILnpE0b783RoWAwG8w\nhVe6R8nDZJKNMPHzWDsZCOx0HKbnpAi7hvgXPbi5oO/iKyA6oViSqF2O15MEWID1\nluQJ9ptWs2yJ2y2bOUdTH2GdVu9lA/9e0wV55fx9uFj4pHODaJm1qKOl1Jk0i7zy\nwUGzmJgXvzYnq9h7XErfPcOsh6Rpkc9a0VpKVgmZtBzzikqcM62OgO9BxqCqFN7U\niBiR9Xq5FyfaeXnPYW/qV1i3dgdd4wRmO3ksQNSPf7cLE122WbVuYFw5dIiyw7lK\ndsyXDW7FAoGBAPurwwBk0n3p0BlZDXx56JO/IOjfP+p4skoG9opxcOVznA0zA9p0\npSezNecimyh6AiPFApruLvIRk+crMZQRVADEQ/MhzvXedEvevtI8SlEKrV9UgRcn\nbp62EZqagf1AJEDWC6G+DxEv1s3ZF4Y7bdHWLDFAv4Ciw7bjQZVDdDq/AoGBANeC\nuAYUtQDxLMuOAWidW6SX4wMUg05AiVPAKka7NIq8H1ylFNpnEjYi9Af4BZeA1/IO\neQWap101qbqW/tIlgnwR6NSLfLSvFCSon0ysQ9lHNJhmC7YOoDQRQHQssIrbJaN3\nQ6JxVZ10Un8IowqxDZIVfXCFuw5fNorOv9hAbFjbAoGAJzBGzB/m+v5WjivkwrZE\n9gSz/i8NR9iFgqt05nflqYUIDrIb7n9tXDI2uYgU+weMn79EuZVPMBh2nG+IZ9MO\n7pOhNRHVpUl/eHT158zFkbsE5ixFcbKNMh+NvDJE/YdoXcQ2yXfL5tQ5MZKVbCyC\n3ELqXL0FVOWDbk4S30hCqAcCgYEAo2+R6aKohjdgllpyPQkhJ9i8I2jaD20n+CjC\npvNv7EqwqgzTnLIQAJhPYv+4FeZzXjGVnCdmB20b89JxG6OwqjDW1uGVyF0CNK7g\naEA4ED5M58pz1TSQUAxJShFeLV/20lovI7E5kXhW29oL857EQOYlREFW05Zngas7\nmF97C4MCgYEAj3bNc1qsHiHXlpzhEA3zWSOVMI4gBin9NZihrr8FO5W621BD8uu7\noWAkjT0AziDvmjdnODdUcAGOd3KX5bYpKuiYCNN3E6MvOhreaurJH23xJM+nsJcR\nzoVMfHWf6lDLzBU96kHtQzaokNoq3TNr8N8prqu2h9poyu5ZNhcrN0Q=\n-----END RSA PRIVATE KEY-----\n")
    
var GoproxyCa, goproxyCaErr = tls.X509KeyPair(CA_CERT, CA_KEY)