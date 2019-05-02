package controllers

import (
	"am.ca-server/helpers"
	"context"
	"encoding/json"
	"fmt"
	"github.com/mailgun/mailgun-go"
	"net/http"
	"os"
	"time"
)

// BaseURL Route
// Route : '/'
// Type  : 'GET'
func BaseURL(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(helpers.Response{Error: true, Code: 404, Message: "Invalid Route"})
}

// EmailService Route
// Route : '/email
// Type  : 'POST'
func EmailService(w http.ResponseWriter, r *http.Request) {
	// Read body
	var responseEmail helpers.Email
	json.NewDecoder(r.Body).Decode(&responseEmail)

	// set headers
	w.Header().Set("content-type", "application/json")

	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_API_KEY"))

	// Create and format mailgun email
	messageSubject := fmt.Sprintf("[alexmontague.ca] - %s", responseEmail.Subject)
	messageBody := fmt.Sprintf("Sent By: %s\n\nSender Email: %s\n\n%s\n", responseEmail.Sender, responseEmail.FromEmail, responseEmail.Message)
	message := mg.NewMessage("info@bookbuy.ca", messageSubject, messageBody, responseEmail.ToEmail)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message
	resp, id, err := mg.Send(ctx, message)
	if err != nil {
		json.NewEncoder(w).Encode(helpers.Response{Error: true, Code: 401, Message: "Something went wrong. Sorry!"})
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("ID: %s Resp: %s\n", id, resp)

	json.NewEncoder(w).Encode(helpers.Response{Error: false, Code: 200, Message: "Email Received"})
}
