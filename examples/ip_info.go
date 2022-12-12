package main

import (
	"fmt"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	params := make(url.Values, 2)

	// IPv4 or IPv6 address
	params.Add("ip", "1.1.1.1")

	// Do a reverse DNS (PTR) lookup. This option can add extra delay to the request so only use it if
	// you need it
	params.Add("reverse-lookup", "false")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.IPInfo(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// Name of the city (if detectable)
		fmt.Printf("city: \"%s\"\n", data["city"])

		// ISO 2-letter continent code
		fmt.Printf("continent-code: \"%s\"\n", data["continent-code"])

		// Full country name
		fmt.Printf("country: \"%s\"\n", data["country"])

		// ISO 2-letter country code
		fmt.Printf("country-code: \"%s\"\n", data["country-code"])

		// ISO 3-letter country code
		fmt.Printf("country-code3: \"%s\"\n", data["country-code3"])

		// ISO 4217 currency code associated with the country
		fmt.Printf("currency-code: \"%s\"\n", data["currency-code"])

		// The IPs host domain (only set if reverse-lookup has been used)
		fmt.Printf("host-domain: \"%s\"\n", data["host-domain"])

		// The IPs full hostname (only set if reverse-lookup has been used)
		fmt.Printf("hostname: \"%s\"\n", data["hostname"])

		// The IP address
		fmt.Printf("ip: \"%s\"\n", data["ip"])

		// True if this is a bogon IP address such as a private network, local network or reserved address
		fmt.Printf("is-bogon: %t\n", data["is-bogon"])

		// True if this is a IPv4 mapped IPv6 address
		fmt.Printf("is-v4-mapped: %t\n", data["is-v4-mapped"])

		// True if this is a IPv6 address. False if IPv4
		fmt.Printf("is-v6: %t\n", data["is-v6"])

		// Location latitude
		fmt.Printf("latitude: %.f\n", data["latitude"])

		// Location longitude
		fmt.Printf("longitude: %.f\n", data["longitude"])

		// Name of the region (if detectable)
		fmt.Printf("region: \"%s\"\n", data["region"])

		// ISO 3166-2 region code (if detectable)
		fmt.Printf("region-code: \"%s\"\n", data["region-code"])

		// Map containing timezone details for the location
		fmt.Printf("timezone: %s\n", data["timezone"])

		// True if this is a valid IPv4 or IPv6 address
		fmt.Printf("valid: %t\n", data["valid"])

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
