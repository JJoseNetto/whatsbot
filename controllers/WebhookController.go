package controllers

import (
	"whatsapp-gtm/services"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading body: %v", err)
		return
	}
	fmt.Println("webhook" + string(body))
	values, err := url.ParseQuery(string(body))
	if err != nil {
		fmt.Fprintf(w, "Error parse body: %v", err)
		return
	}
	event := values.Get("event")
	messageBody := values.Get("message_body")
	phoneNumber := values.Get("contact_phone_number")

	if event == "message" {
		if err := services.HandleResponseMessage(phoneNumber, messageBody); err != nil {
			fmt.Fprintf(w, "Error handling message: %v", err)
			return
		}
		fmt.Fprintf(w, "Message status updated successfully")
	}
}
