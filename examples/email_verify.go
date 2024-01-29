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

		// The domain name of this email address
		fmt.Printf("domain: \"%s\"\n", data["domain"])

		// True if this address has any domain name or DNS related errors. Check the 'domain-status' field
		// for the detailed error reason
		fmt.Printf("domain-error: %t\n", data["domain-error"])

		// The email domain status, possible values are:
		// • ok - the domain is in working order and can receive email
		// • invalid - the domain is not a conformant hostname. May contain invalid syntax or characters
		// • no-service - the domain owner has indicated there is no mail service on the domain (also
		//   known as the 'Null MX')
		// • no-mail - the domain has no valid MX records so cannot receive email
		// • mx-invalid - MX records contain invalid or non-conformant hostname values
		// • mx-bogon - MX records point to bogon IP addresses
		// • resolv-error - MX records do not resolve to any valid IP addresses
		fmt.Printf("domain-status: \"%s\"\n", data["domain-status"])

		// The complete email address. If you enabled the 'fix-typos' option then this will be the corrected
		// address
		fmt.Printf("email: \"%s\"\n", data["email"])

		// True if this email domain has a catch-all policy. A catch-all domain will accept mail for any
		// username so therefor the 'smtp-status' will always be 'ok'
		fmt.Printf("is-catch-all: %t\n", data["is-catch-all"])

		// True if the mail server responded with a temporary failure (either a 4xx response code or
		// unresponsive server). You can retry this address later, we recommend waiting at least 15 minutes
		// before retrying
		fmt.Printf("is-deferred: %t\n", data["is-deferred"])

		// True if this address is a disposable, temporary or darknet related email address
		fmt.Printf("is-disposable: %t\n", data["is-disposable"])

		// True if this address is from a free email provider
		fmt.Printf("is-freemail: %t\n", data["is-freemail"])

		// True if this address likely belongs to a person. False if this is a role based address, e.g.
		// admin@, help@, office@, etc.
		fmt.Printf("is-personal: %t\n", data["is-personal"])

		// The first resolved IP address of the primary MX server, may be empty if there are domain errors
		// present
		fmt.Printf("mx-ip: \"%s\"\n", data["mx-ip"])

		// The domain name of the email hosting provider
		fmt.Printf("provider: \"%s\"\n", data["provider"])

		// The raw SMTP response message received during verification
		fmt.Printf("smtp-response: \"%s\"\n", data["smtp-response"])

		// The SMTP username verification status for this address:
		// • ok - verification was successful, this is a real username that can receive mail
		// • absent - this username or domain is not registered with the email service provider
		// • invalid - not a valid email address, check the 'domain-status' field for specific details
		// • unresponsive - the mail servers for this domain have repeatedly timed-out or refused multiple
		//   connection attempts
		// • unknown - sorry, we could not reliably determine the status of this username
		fmt.Printf("smtp-status: \"%s\"\n", data["smtp-status"])

		// True if this address has any syntax errors or is not in RFC compliant formatting
		fmt.Printf("syntax-error: %t\n", data["syntax-error"])

		// True if any typos have been fixed. The 'fix-typos' option must be enabled for this to work
		fmt.Printf("typos-fixed: %t\n", data["typos-fixed"])

		// Is this a valid email address. To be valid an email must have: correct syntax, a registered and
		// active domain name, correct DNS records and operational MX servers
		fmt.Printf("valid: %t\n", data["valid"])

		// True if this email address has passed SMTP username verification. Check the 'smtp-status' and
		// 'domain-status' fields for specific verification details
		fmt.Printf("verified: %t\n", data["verified"])

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
