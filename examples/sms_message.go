package main

import (
	"fmt"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	params := make(url.Values, 5)

	// The phone number to send a message to
	params.Add("number", "+12106100045")

	// ISO 2-letter country code, assume numbers are based in this country. If not set numbers are
	// assumed to be in international format (with or without the leading + sign)
	params.Add("country-code", "")

	// Limit the total number of SMS allowed to the supplied phone number, if the limit is reached
	// within the TTL then error code 14 will be returned
	params.Add("limit", "10")

	// The SMS message to send. Messages are truncated to a maximum of 150 characters for ASCII content
	// OR 70 characters for UTF content
	params.Add("message", "Hello, this is a test message!")

	// Set the TTL in number of days that the 'limit' option will remember a phone number (the default
	// is 1 day and the maximum is 365 days)
	params.Add("limit-ttl", "1")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.SMSMessage(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// True if this a valid phone number
		fmt.Printf("number-valid: %t\n", data["number-valid"])

		// True if the SMS has been sent
		fmt.Printf("sent: %t\n", data["sent"])

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
