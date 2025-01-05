package services

import (
	"bytes"
	"fmt"
	"html/template"
	"whatsapp-gtm/config"
	"whatsapp-gtm/enums"
	"whatsapp-gtm/helpers"
	"whatsapp-gtm/models"
)

var approveMessageTemplate = `*Detalhes da Solicitação Aprovada*
📌 *Ação*: {{.Acao}}
👤 *Responsavel Ação*: {{.ResponsavelAcao}}
📝 *Descrição*: {{.Descricao}}
💸 *Valor Orçado*: R$ {{.ValorOrcado}}
💰 *Valor Ajustado*: R$ {{.ValorAjustado}}
🏢 *Cliente*: {{.Cliente}}
🔍 *ID GTM*: {{.IDGTM}}

✅ *Solicitação*: Aprovada com sucesso!`

var reproveMessageTemplate = `*Detalhes da Solicitação Reprovada*
📌 *Ação*: {{.Acao}}
👤 *Responsavel Ação*: {{.ResponsavelAcao}}
📝 *Descrição*: {{.Descricao}}
💸 *Valor Orçado*: R$ {{.ValorOrcado}}
💰 *Valor Ajustado*: R$ {{.ValorAjustado}}
🏢 *Cliente*: {{.Cliente}}
🔍 *ID GTM*: {{.IDGTM}}

❌ *Solicitação*: Reprovada com sucesso!

🔴 *Motivo da Reprovação*: *Por favor, escreva abaixo o motivo da reprovação.*`

func HandleResponseMessage(phoneNumber string, messageBody string) error {
	db := helpers.Db()

	var message models.GtmMessage
	if err := db.Where("phone_number = ? AND (status = ? OR (status = ? AND commentary = ''))",
		phoneNumber, enums.Pending, enums.Reproved).Order("created_at ASC").First(&message).Error; err != nil {
		return fmt.Errorf("Pending or canceled message without commentary not found")
	}

	data := models.MessageData{
		Acao:            message.Acao,
		ResponsavelAcao: message.ResponsavelAcao,
		Descricao:       message.Descricao,
		ValorOrcado:     message.ValorOrcado,
		ValorAjustado:   message.ValorAjustado,
		Cliente:         message.Cliente,
		IDGTM:           message.IDGtm,
	}

	if message.Status == enums.Reproved && message.Commentary == "" {
		message.Commentary = messageBody
		CommentaryMessage(message)
		if err := db.Save(&message).Error; err != nil {
			return fmt.Errorf("Failed to update message status")
		}
		HandleReproveDescription(messageBody, message.UserID)
		return nil
	}
	switch messageBody {
	case "1":
		fmt.Println("Aprovado")
		message.Status = enums.Approved
		ApproveMessage(message, data)
		HandlePsaStatus(message, 1)
		// HandleConsultaStatus(message, "CONFIRMED")
	case "2":
		fmt.Println("Reprovado")
		message.Status = enums.Reproved
		ReproveMessage(message, data)
		HandlePsaStatus(message, 0)
		// services.HandleConsultaStatus(message, "CANCELLED")
	default:
		fmt.Println("Default")
		RetryMessage(message)
		return fmt.Errorf("Invalid response")
	}

	if err := db.Save(&message).Error; err != nil {
		return fmt.Errorf("Failed to update message status")
	}
	return nil
}

func ApproveMessage(message models.GtmMessage, data models.MessageData) error {

	tmpl, err := template.New("confirmMessage").Parse(approveMessageTemplate)
	if err != nil {
		return fmt.Errorf("erro ao compilar o template de confirmação: %v", err)
	}

	var messageBody bytes.Buffer
	err = tmpl.Execute(&messageBody, data)
	if err != nil {
		return fmt.Errorf("erro ao aplicar os dados ao template de confirmação: %v", err)
	}

	confirmedMessage := models.Message{
		APIKey:             config.ApiKey,
		PhoneNumber:        config.PhoneNumber,
		ContactPhoneNumber: message.PhoneNumber,
		MessageCustomID:    message.ID,
		MessageType:        "text",
		MessageBody:        messageBody.String(),
	}
	fmt.Print(confirmedMessage)
	helpers.SendMessage(confirmedMessage)
	return nil
}

func ReproveMessage(message models.GtmMessage, data models.MessageData) error {

	tmpl, err := template.New("confirmMessage").Parse(reproveMessageTemplate)
	if err != nil {
		return fmt.Errorf("erro ao compilar o template de reprovar: %v", err)
	}

	var messageBody bytes.Buffer
	err = tmpl.Execute(&messageBody, data)
	if err != nil {
		return fmt.Errorf("erro ao aplicar os dados ao template de confirmar: %v", err)
	}
	cancelledMessage := models.Message{
		APIKey:             config.ApiKey,
		PhoneNumber:        config.PhoneNumber,
		ContactPhoneNumber: message.PhoneNumber,
		MessageCustomID:    message.ID,
		MessageType:        "text",
		MessageBody:        messageBody.String(),
	}
	helpers.SendMessage(cancelledMessage)
	return nil
}

func CommentaryMessage(message models.GtmMessage) {
	commnentaryMessage := models.Message{
		APIKey:             config.ApiKey,
		PhoneNumber:        config.PhoneNumber,
		ContactPhoneNumber: message.PhoneNumber,
		MessageCustomID:    message.ID,
		MessageType:        "text",
		MessageBody:        fmt.Sprintf("🔴 *Motivo da Reprovação Registrado com Sucesso!*\n\n📄 *Motivo*: %s", message.Commentary),
	}
	helpers.SendMessage(commnentaryMessage)
}

func RetryMessage(message models.GtmMessage) {
	retryMessage := models.Message{
		APIKey:             config.ApiKey,
		PhoneNumber:        config.PhoneNumber,
		ContactPhoneNumber: message.PhoneNumber,
		MessageCustomID:    message.ID,
		MessageType:        "text",
		MessageBody:        "Resposta inválida. Por favor, responda com '1' para aprovar ou '2' para reprovar a solicita.",
	}
	helpers.SendMessage(retryMessage)
}

func HandlePsaStatus(message models.GtmMessage, status int) {
	db := helpers.Db()

	query := "UPDATE icl_products_services_acquisition SET is_approve = ? WHERE id = ?"

	result := db.Exec(query, status, message.UserID)

	if result.Error != nil {
		fmt.Printf("Erro ao atualizar is_approve: %v\n", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		fmt.Println("Nenhum registro encontrado para atualizar.")
		return
	}

	fmt.Println("is_approve atualizado com sucesso.")
}

func HandleReproveDescription(description string, PsaId int) {
	db := helpers.Db()

	query := "UPDATE icl_products_services_acquisition SET reprove_description = ? WHERE id = ?"

	result := db.Exec(query, description, PsaId)

	if result.Error != nil {
		fmt.Printf("Erro ao atualizar description: %v\n", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		fmt.Println("Nenhum registro encontrado para atualizar.")
		return
	}

	fmt.Println("Reprove description atualizada com sucesso.")
}
