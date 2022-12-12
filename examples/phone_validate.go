package main

import (
	"fmt"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	params := make(url.Values, 3)

	// A phone number. This can be in international format (E.164) or local format. If passing local
	// format you must also set either the 'country-code' OR 'ip' options as well
	params.Add("number", "+6495552000")

	// ISO 2-letter country code, assume numbers are based in this country. If not set numbers are
	// assumed to be in international format (with or without the leading + sign)
	params.Add("country-code", "")

	// Pass in a users IP address and we will assume numbers are based in the country of the IP address
	params.Add("ip", "")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.PhoneValidate(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// The phone number country
		fmt.Printf("country: \"%s\"\n", data["country"])

		// The phone number country as an ISO 2-letter country code
		fmt.Printf("country-code: \"%s\"\n", data["country-code"])

		// The phone number country as an ISO 3-letter country code
		fmt.Printf("country-code3: \"%s\"\n", data["country-code3"])

		// ISO 4217 currency code associated with the country
		fmt.Printf("currency-code: \"%s\"\n", data["currency-code"])

		// The international calling code
		fmt.Printf("international-calling-code: \"%s\"\n", data["international-calling-code"])

		// The number represented in full international format (E.164)
		fmt.Printf("international-number: \"%s\"\n", data["international-number"])

		// True if this is a mobile number. If the number type is unknown this value will be false
		fmt.Printf("is-mobile: %t\n", data["is-mobile"])

		// The number represented in local dialing format
		fmt.Printf("local-number: \"%s\"\n", data["local-number"])

		// The phone number location. Could be the city, region or country depending on the type of number
		fmt.Printf("location: \"%s\"\n", data["location"])

		// The network/carrier who owns the prefix (this only works for some countries, use HLR lookup for
		// global network detection)
		fmt.Printf("prefix-network: \"%s\"\n", data["prefix-network"])

		// The number type based on the number prefix. Possible values are:
		// • mobile
		// • fixed-line
		// • premium-rate
		// • toll-free
		// • voip
		// • unknown (use HLR lookup)
		fmt.Printf("type: \"%s\"\n", data["type"])

		// Is this a valid phone number
		fmt.Printf("valid: %t\n", data["valid"])

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
