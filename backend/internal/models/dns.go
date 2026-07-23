package models

type DNSResponse struct {
	Domain      string   `json:"domain"`
	ARecords    []string `json:"a_records"`
	AAAARecords []string `json:"aaaa_records"`
	MXRecords   []string `json:"mx_records"`
	TXTRecords  []string `json:"txt_records"`
}
