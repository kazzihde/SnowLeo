package models

type WhoisResponse struct {
	Input       string   `json:"input"`
	Registrar   string   `json:"registrar"`
	Created     string   `json:"created"`
	Updated     string   `json:"updated"`
	Expires     string   `json:"expires"`
	NameServers []string `json:"name_servers"`
	Status      []string `json:"status"`
}
