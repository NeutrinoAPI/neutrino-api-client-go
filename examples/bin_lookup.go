package main

import (
	"fmt"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
	"strings"
)

func main() {
	params := make(url.Values, 2)

	// The BIN or IIN number. This is the first 6, 8 or 10 digits of a card number, use 8 (or more)
	// digits for the highest level of accuracy
	params.Add("bin-number", "48334884")

	// Pass in the customers IP address and we will return some extra information about them
	params.Add("customer-ip", "")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.BINLookup(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// The BIN number returned. You may count the number of digits in this field to determine if the BIN
		// is likely to be based on an 8-digit or 6-digit card
		fmt.Printf("bin-number: \"%s\"\n", data["bin-number"])

		// The card brand (e.g. Visa or Mastercard)
		fmt.Printf("card-brand: \"%s\"\n", data["card-brand"])

		// The card category. There are many different card categories the most common card categories are:
		// CLASSIC, BUSINESS, CORPORATE, PLATINUM, PREPAID
		fmt.Printf("card-category: \"%s\"\n", data["card-category"])

		// The card type, will always be one of: DEBIT, CREDIT, CHARGE CARD
		fmt.Printf("card-type: \"%s\"\n", data["card-type"])

		// The full country name of the issuer
		fmt.Printf("country: \"%s\"\n", data["country"])

		// The ISO 2-letter country code of the issuer
		fmt.Printf("country-code: \"%s\"\n", data["country-code"])

		// The ISO 3-letter country code of the issuer
		fmt.Printf("country-code3: \"%s\"\n", data["country-code3"])

		// ISO 4217 currency code associated with the country of the issuer
		fmt.Printf("currency-code: \"%s\"\n", data["currency-code"])

		// True if the customers IP is listed on one of our blocklists, see the IP Blocklist API
		fmt.Printf("ip-blocklisted: %t\n", data["ip-blocklisted"])

		// An array of strings indicating which blocklists this IP is listed on
		ipBlocklists := strings.Fields(fmt.Sprint(data["ip-blocklists"]))
		fmt.Printf("ip-blocklists: %s\n", strings.Join(ipBlocklists, ", "))

		// The city of the customers IP (if detectable)
		fmt.Printf("ip-city: \"%s\"\n", data["ip-city"])

		// The country of the customers IP
		fmt.Printf("ip-country: \"%s\"\n", data["ip-country"])

		// The ISO 2-letter country code of the customers IP
		fmt.Printf("ip-country-code: \"%s\"\n", data["ip-country-code"])

		// The ISO 3-letter country code of the customers IP
		fmt.Printf("ip-country-code3: \"%s\"\n", data["ip-country-code3"])

		// True if the customers IP country matches the BIN country
		fmt.Printf("ip-matches-bin: %t\n", data["ip-matches-bin"])

		// The region of the customers IP (if detectable)
		fmt.Printf("ip-region: \"%s\"\n", data["ip-region"])

		// Is this a commercial/business use card
		fmt.Printf("is-commercial: %t\n", data["is-commercial"])

		// Is this a prepaid or prepaid reloadable card
		fmt.Printf("is-prepaid: %t\n", data["is-prepaid"])

		// The card issuer
		fmt.Printf("issuer: \"%s\"\n", data["issuer"])

		// The card issuers phone number
		fmt.Printf("issuer-phone: \"%s\"\n", data["issuer-phone"])

		// The card issuers website
		fmt.Printf("issuer-website: \"%s\"\n", data["issuer-website"])

		// Is this a valid BIN or IIN number
		fmt.Printf("valid: %t\n", data["valid"])

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
