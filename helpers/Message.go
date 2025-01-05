package helpers

import (
	"bytes"
	"encoding/json"
	"whatsapp-gtm/config"
	"whatsapp-gtm/models"
	"fmt"
	"net/http"
)

func SendMessage(message models.Message) error {
	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("erro ao converter mensagem para JSON: %v", err)
	}
	resp, err := http.Post(config.SendUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("erro ao fazer a requisição POST: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("recebido status não-200: %s", resp.Status)
	}

	return nil
}
