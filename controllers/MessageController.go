package controllers

import (
	"encoding/json"
	"whatsapp-gtm/models"
	"whatsapp-gtm/services"
	"io"
	"net/http"
)

func Send(w http.ResponseWriter, r *http.Request) {
	res, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
	}

	var gtmRequest models.GtmRequest

	err = json.Unmarshal(res, &gtmRequest)
	if err != nil {
		http.Error(w, "Failed to deserealized json", http.StatusInternalServerError)
	}

	services.HandleMessage(gtmRequest)
}
