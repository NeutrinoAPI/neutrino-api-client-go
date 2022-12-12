package main

import (
	"fmt"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	params := make(url.Values, 2)

	// An email address
	params.Add("email", "tech@neutrinoapi.com")

	// Automatically attempt to fix typos in the address
	params.Add("fix-typos", "false")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.EmailValidate(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// The email domain
		fmt.Printf("domain: \"%s\"\n", data["domain"])

		// True if this address has a domain error (e.g. no valid mail server records)
		fmt.Printf("domain-error: %t\n", data["domain-error"])

		// The email address. If you have used the fix-typos option then this will be the fixed address
		fmt.Printf("email: \"%s\"\n", data["email"])

		// True if this address is a disposable, temporary or darknet related email address
		fmt.Printf("is-disposable: %t\n", data["is-disposable"])

		// True if this address is a free-mail address
		fmt.Printf("is-freemail: %t\n", data["is-freemail"])

		// True if this address belongs to a person. False if this is a role based address, e.g. admin@,
		// help@, office@, etc.
		fmt.Printf("is-personal: %t\n", data["is-personal"])

		// The email service provider domain
		fmt.Printf("provider: \"%s\"\n", data["provider"])

		// True if this address has a syntax error
		fmt.Printf("syntax-error: %t\n", data["syntax-error"])

		// True if typos have been fixed
		fmt.Printf("typos-fixed: %t\n", data["typos-fixed"])

		// Is this a valid email
		fmt.Printf("valid: %t\n", data["valid"])

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
