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
	response := neutrinoAPIClient.EmailVerify(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// The email domain
		fmt.Printf("domain: \"%s\"\n", data["domain"])

		// True if this address has a domain error (e.g. no valid mail server records)
		fmt.Printf("domain-error: %t\n", data["domain-error"])

		// The email address. If you have used the fix-typos option then this will be the fixed address
		fmt.Printf("email: \"%s\"\n", data["email"])

		// True if this email domain has a catch-all policy (it will accept mail for any username)
		fmt.Printf("is-catch-all: %t\n", data["is-catch-all"])

		// True if the mail server responded with a temporary failure (either a 4xx response code or
		// unresponsive server). You can retry this address later, we recommend waiting at least 15 minutes
		// before retrying
		fmt.Printf("is-deferred: %t\n", data["is-deferred"])

		// True if this address is a disposable, temporary or darknet related email address
		fmt.Printf("is-disposable: %t\n", data["is-disposable"])

		// True if this address is a free-mail address
		fmt.Printf("is-freemail: %t\n", data["is-freemail"])

		// True if this address is for a person. False if this is a role based address, e.g. admin@, help@,
		// office@, etc.
		fmt.Printf("is-personal: %t\n", data["is-personal"])

		// The email service provider domain
		fmt.Printf("provider: \"%s\"\n", data["provider"])

		// The raw SMTP response message received during verification
		fmt.Printf("smtp-response: \"%s\"\n", data["smtp-response"])

		// The SMTP verification status for the address:
		// • ok - SMTP verification was successful, this is a real address that can receive mail
		// • invalid - this is not a valid email address (has either a domain or syntax error)
		// • absent - this address is not registered with the email service provider
		// • unresponsive - the mail server(s) for this address timed-out or refused to open an SMTP
		//   connection
		// • unknown - sorry, we could not reliably determine the real status of this address (this
		//   address may or may not exist)
		fmt.Printf("smtp-status: \"%s\"\n", data["smtp-status"])

		// True if this address has a syntax error
		fmt.Printf("syntax-error: %t\n", data["syntax-error"])

		// True if typos have been fixed
		fmt.Printf("typos-fixed: %t\n", data["typos-fixed"])

		// Is this a valid email address (syntax and domain is valid)
		fmt.Printf("valid: %t\n", data["valid"])

		// True if this address has passed SMTP verification. Check the smtp-status and smtp-response fields
		// for specific verification details
		fmt.Printf("verified: %t\n", data["verified"])

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
