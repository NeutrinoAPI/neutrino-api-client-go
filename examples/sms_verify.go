package main

import (
	"fmt"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	params := make(url.Values, 8)

	// The phone number to send a verification code to
	params.Add("number", "+12106100045")

	// ISO 2-letter country code, assume numbers are based in this country. If not set numbers are
	// assumed to be in international format (with or without the leading + sign)
	params.Add("country-code", "")

	// Pass in your own security code. This is useful if you have implemented TOTP or similar 2FA
	// methods. If not set then we will generate a secure random code
	params.Add("security-code", "")

	// The language to send the verification code in, available languages are:
	// • de - German
	// • en - English
	// • es - Spanish
	// • fr - French
	// • it - Italian
	// • pt - Portuguese
	// • ru - Russian
	params.Add("language-code", "en")

	// The number of digits to use in the security code (must be between 4 and 12)
	params.Add("code-length", "5")

	// Limit the total number of SMS allowed to the supplied phone number, if the limit is reached
	// within the TTL then error code 14 will be returned
	params.Add("limit", "10")

	// Set a custom brand or product name in the verification message
	params.Add("brand-name", "")

	// Set the TTL in number of days that the 'limit' option will remember a phone number (the default
	// is 1 day and the maximum is 365 days)
	params.Add("limit-ttl", "1")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.SMSVerify(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// True if this a valid phone number
		fmt.Printf("number-valid: %t\n", data["number-valid"])

		// The security code generated, you can save this code to perform your own verification or you can
		// use the Verify Security Code API
		fmt.Printf("security-code: \"%s\"\n", data["security-code"])

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
