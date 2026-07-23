package ip

import (
	"net"

	"github.com/datashelll/SnowLeo/internal/models"
)

func Lookup(input string) (models.IPResponse, error) {
	response := models.IPResponse{
		Input: input,
	}

	if net.ParseIP(input) != nil {
		response.InputType = "ip"

		hostnames, err := net.LookupAddr(input)
		if err != nil {
			return response, err
		}

		response.Hostnames = hostnames
	} else {
		response.InputType = "hostname"

		ips, err := net.LookupIP(input)
		if err != nil {
			return response, err
		}

		for _, ip := range ips {
			response.IPs = append(response.IPs, ip.String())
		}
	}

	return response, nil
}
