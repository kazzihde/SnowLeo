package whois

import (
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"

	"github.com/datashelll/SnowLeo/internal/models"
)

func Lookup(input string) (models.WhoisResponse, error) {
	raw, err := whois.Whois(input)
	if err != nil {
		return models.WhoisResponse{}, err
	}

	parsed, err := whoisparser.Parse(raw)
	if err != nil {
		return models.WhoisResponse{}, err
	}

	response := models.WhoisResponse{
		Input:       input,
		Registrar:   parsed.Registrar.Name,
		Created:     parsed.Domain.CreatedDate,
		Updated:     parsed.Domain.UpdatedDate,
		Expires:     parsed.Domain.ExpirationDate,
		NameServers: parsed.Domain.NameServers,
		Status:      parsed.Domain.Status,
	}

	return response, nil
}
