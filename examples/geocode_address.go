package main

import (
	"fmt"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
	"strings"
)

func main() {
	params := make(url.Values, 10)

	// The full address, partial address or name of a place to try and locate. Comma separated address
	// components are preferred.
	params.Add("address", "1 Molesworth Street, Thorndon, Wellington 6011")

	// The house/building number to locate
	params.Add("house-number", "")

	// The street/road name to locate
	params.Add("street", "")

	// The city/town name to locate
	params.Add("city", "")

	// The county/region name to locate
	params.Add("county", "")

	// The state name to locate
	params.Add("state", "")

	// The postal code to locate
	params.Add("postal-code", "")

	// Limit result to this country (the default is no country bias)
	params.Add("country-code", "")

	// The language to display results in, available languages are:
	// • de, en, es, fr, it, pt, ru, zh
	params.Add("language-code", "en")

	// If no matches are found for the given address, start performing a recursive fuzzy search until a
	// geolocation is found. This option is recommended for processing user input or implementing
	// auto-complete. We use a combination of approximate string matching and data cleansing to find
	// possible location matches
	params.Add("fuzzy-search", "false")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.GeocodeAddress(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// The number of possible matching locations found
		fmt.Printf("found: %.f\n", data["found"])

		// Array of matching location objects
		fmt.Printf("locations:\n")
		locations := data["locations"].([]interface{})
		for _, item := range locations {
			itemMap := item.(map[string]interface{})

			// The fully formatted address
			fmt.Printf("    address: \"%s\"\n", itemMap["address"])

			// The components which make up the address such as road, city, state, etc
			fmt.Printf("    address-components: %s\n", itemMap["address-components"])

			// The city of the location
			fmt.Printf("    city: \"%s\"\n", itemMap["city"])

			// The country of the location
			fmt.Printf("    country: \"%s\"\n", itemMap["country"])

			// The ISO 2-letter country code of the location
			fmt.Printf("    country-code: \"%s\"\n", itemMap["country-code"])

			// The ISO 3-letter country code of the location
			fmt.Printf("    country-code3: \"%s\"\n", itemMap["country-code3"])

			// ISO 4217 currency code associated with the country
			fmt.Printf("    currency-code: \"%s\"\n", itemMap["currency-code"])

			// The location latitude
			fmt.Printf("    latitude: %.f\n", itemMap["latitude"])

			// Array of strings containing any location tags associated with the address. Tags are additional
			// pieces of metadata about a specific location, there are thousands of different tags. Some
			// examples of tags: shop, office, cafe, bank, pub
			locationTags := strings.Fields(fmt.Sprint(itemMap["location-tags"]))
			fmt.Printf("    location-tags: %s\n", strings.Join(locationTags, ", "))

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
			fmt.Printf("    location-type: \"%s\"\n", itemMap["location-type"])

			// The location longitude
			fmt.Printf("    longitude: %.f\n", itemMap["longitude"])

			// The postal code for the location
			fmt.Printf("    postal-code: \"%s\"\n", itemMap["postal-code"])

			// The state of the location
			fmt.Printf("    state: \"%s\"\n", itemMap["state"])

			// Map containing timezone details for the location
			fmt.Printf("    timezone: %s\n", itemMap["timezone"])
		fmt.Println()
		}

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
