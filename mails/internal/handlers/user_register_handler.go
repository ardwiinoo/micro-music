package handlers

import (
	"encoding/json"

	"github.com/ardwiinoo/micro-music/mails/internal/email"
)

type UserRegisteredPayload struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	ActLink string `json:"activation_link"`
}

func UserRegisteredHandler(mailer *email.Mailer) func([]byte) error {
	return func(payload []byte) error {
		var data UserRegisteredPayload
		
		if err := json.Unmarshal(payload, &data); err != nil {
			return err
		}

		return mailer.SendMail(
			data.Email,
			"Welcome to MicroMusic",
			"user_registered.html", 
			map[string]interface{}{
				"Name":          data.Name,
				"ActivationURL": data.ActLink,
			},
		)
	}
}
