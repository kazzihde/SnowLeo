package dns

import (
	"net"

	"github.com/datashelll/SnowLeo/internal/models"
)

func Lookup(domain string) models.DNSResponse {
	var result models.DNSResponse
	result.Domain = domain

	// A records
	if aRecords, err := net.LookupHost(domain); err == nil {
		result.ARecords = aRecords
	}

	// MX records
	if mxRecords, err := net.LookupMX(domain); err == nil {
		for _, mx := range mxRecords {
			result.MXRecords = append(result.MXRecords, mx.Host)
		}
	}

	// TXT records
	if txtRecords, err := net.LookupTXT(domain); err == nil {
		result.TXTRecords = txtRecords
	}

	// AAAA records
	if ipAddrs, err := net.LookupIP(domain); err == nil {
		for _, ip := range ipAddrs {
			if ip.To4() == nil {
				result.AAAARecords = append(result.AAAARecords, ip.String())
			}
		}
	}

	return result
}
