package repository

type Customer struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	SmsNumber string `json:"sms_number"`
	Enabled   bool   `json:"enabled"`
}
