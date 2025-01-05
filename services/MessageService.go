package services

import (
	"bytes"
	"fmt"
	"html/template"
	"time"
	"whatsapp-gtm/config"
	"whatsapp-gtm/enums"
	"whatsapp-gtm/helpers"
	"whatsapp-gtm/models"
)

func HandleMessage(gtmRequest models.GtmRequest) {
	message := BuildMessageBody(gtmRequest)
	helpers.SendMessage(message)
	// StoreMessage(gtmRequest, message)
}

func BuildMessageBody(gtm models.GtmRequest) models.Message {
	data := models.MessageData{
		Acao:            gtm.Acao,
		ResponsavelAcao: gtm.ResponsavelAcao,
		Descricao:       gtm.Descricao,
		ValorOrcado:     gtm.ValorOrcado,
		ValorAjustado:   gtm.ValorAjustado,
		Cliente:         gtm.Cliente,
		IDGTM:           gtm.IDGtm,
	}

	messageTemplate := `*Detalhes da Solicitação*
📌 *Ação*: {{.Acao}}
👤 *Responsavel Ação*: {{.ResponsavelAcao}}
📝 *Descrição*: {{.Descricao}}
💸 *Valor Orçado*: R$ {{.ValorOrcado}}
💰 *Valor Ajustado*: R$ {{.ValorAjustado}}
🏢 *Cliente*: {{.Cliente}}
🔍 *ID GTM*: {{.IDGTM}}

*Escolha uma opção:*
1️⃣ Aprovar a Solicitação
2️⃣ Reprovar a Solicitação`

	tmpl, err := template.New("message").Parse(messageTemplate)
	if err != nil {
		fmt.Println("Erro ao compilar o template:", err)
	}

	var messageBody bytes.Buffer
	err = tmpl.Execute(&messageBody, data)
	if err != nil {
		fmt.Println("Erro ao aplicar os dados ao template:", err)
	}
	message := models.Message{
		APIKey:             config.ApiKey,
		PhoneNumber:        config.PhoneNumber,
		ContactPhoneNumber: gtm.PhoneNumber,
		MessageCustomID:    gtm.IdUser,
		MessageType:        "text",
		MessageBody:        messageBody.String(),
	}
	return message
}

func StoreMessage(gtm models.GtmRequest, message models.Message) {
	db := helpers.Db()

	gtmMessage := models.GtmMessage{
		Status:          enums.Pending,
		PhoneNumber:     message.ContactPhoneNumber,
		Commentary:      "",
		UserID:          gtm.IdUser,
		Acao:            gtm.Acao,
		ResponsavelAcao: gtm.Acao,
		Descricao:       gtm.Descricao,
		ValorOrcado:     gtm.ValorOrcado,
		ValorAjustado:   gtm.ValorAjustado,
		Cliente:         gtm.Cliente,
		IDGtm:           gtm.IDGtm,
		CreatedAt:       time.Now(),
	}
	db.Create(&gtmMessage)
}
