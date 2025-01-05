package models

type GtmRequest struct {
	IdUser          int    `json:"id_user"`
	Acao            string `json:"acao"`
	ResponsavelAcao string `json:"responsavel_acao"`
	Descricao       string `json:"descricao"`
	ValorOrcado     string `json:"valor_orcado"`
	ValorAjustado   string `json:"valor_ajustado"`
	Cliente         string `json:"cliente"`
	PhoneNumber     string `json:"phone_number"`
	IDGtm           int    `json:"id_gtm"`
}
