package models

type IPResponse struct {
	Input     string   `json:"input"`
	InputType string   `json:"input_type"`
	Hostnames []string `json:"hostnames"`
	IPs       []string `json:"ips"`
}
