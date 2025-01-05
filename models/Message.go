package models

type Message struct {
	APIKey             string `json:"apikey"`
	PhoneNumber        string `json:"phone_number"`
	ContactPhoneNumber string `json:"contact_phone_number"`
	MessageCustomID    int    `json:"message_custom_id"`
	MessageType        string `json:"message_type"`
	MessageBody        string `json:"message_body"`
}
