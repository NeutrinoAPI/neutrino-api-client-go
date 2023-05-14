package main

import (
	"fmt"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
	"strings"
)

func main() {
	params := make(url.Values, 4)

	// The location latitude in decimal degrees format
	params.Add("latitude", "-41.2775847")

	// The location longitude in decimal degrees format
	params.Add("longitude", "174.7775229")

	// The language to display results in, available languages are:
	// • de, en, es, fr, it, pt, ru
	params.Add("language-code", "en")

	// The zoom level to respond with:
	// • address - the most precise address available
	// • street - the street level
	// • city - the city level
	// • state - the state level
	// • country - the country level
	params.Add("zoom", "address")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.GeocodeReverse(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// The complete address using comma-separated values
		fmt.Printf("address: \"%s\"\n", data["address"])

		// The components which make up the address such as road, city, state, etc
		fmt.Printf("address-components: %s\n", data["address-components"])

		// The city of the location
		fmt.Printf("city: \"%s\"\n", data["city"])

		// The country of the location
		fmt.Printf("country: \"%s\"\n", data["country"])

		// The ISO 2-letter country code of the location
		fmt.Printf("country-code: \"%s\"\n", data["country-code"])

		// The ISO 3-letter country code of the location
		fmt.Printf("country-code3: \"%s\"\n", data["country-code3"])

		// ISO 4217 currency code associated with the country
		fmt.Printf("currency-code: \"%s\"\n", data["currency-code"])

		// True if these coordinates map to a real location
		fmt.Printf("found: %t\n", data["found"])

		// The location latitude
		fmt.Printf("latitude: %.f\n", data["latitude"])

		// Array of strings containing any location tags associated with the address. Tags are additional
		// pieces of metadata about a specific location, there are thousands of different tags. Some
		// examples of tags: shop, office, cafe, bank, pub
		locationTags := strings.Fields(fmt.Sprint(data["location-tags"]))
		fmt.Printf("location-tags: %s\n", strings.Join(locationTags, ", "))

		// The detected location type ordered roughly from most to least precise, possible values are:
		// • address - indicates a precise street address
		// • street - accurate to the street level but may not point to the exact location of the
		//   house/building number
		// • city - accurate to the city level, this includes villages, towns, suburbs, etc
		// • postal-code - indicates a postal code area (no house or street information present)
		// • railway - location is part of a rail network such as a station or railway track
		// • natural - indicates a natural feature, for example a mountain peak or a waterway
		// • island - location is an island or archipelago
		// • administrative - indicates an administrative boundary such as a country, state or province
		fmt.Printf("location-type: \"%s\"\n", data["location-type"])

		// The location longitude
		fmt.Printf("longitude: %.f\n", data["longitude"])

		// The formatted address using local standards suitable for printing on an envelope
		fmt.Printf("postal-address: \"%s\"\n", data["postal-address"])

		// The postal code for the location
		fmt.Printf("postal-code: \"%s\"\n", data["postal-code"])

		// The ISO 3166-2 region code for the location
		fmt.Printf("region-code: \"%s\"\n", data["region-code"])

		// The state of the location
		fmt.Printf("state: \"%s\"\n", data["state"])

		// Map containing timezone details
		fmt.Printf("timezone: %s\n", data["timezone"])

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
