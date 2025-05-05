package main

import (
	"fmt"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	params := make(url.Values, 4)

	// The value to convert from (e.g. 10.95)
	params.Add("from-value", "100")

	// The type of the value to convert from (e.g. USD)
	params.Add("from-type", "USD")

	// The type to convert to (e.g. EUR)
	params.Add("to-type", "EUR")

	// Convert using the rate on a historical date, accepted date formats are: YYYY-MM-DD, YYYY-MM,
	// YYYY. Historical rates are stored with daily granularity so the date format YYYY-MM-DD is
	// preferred for the highest precision. If an invalid date or a date too far into the past is
	// supplied then the API will respond with 'valid' as false and an empty 'historical-date'
	params.Add("historical-date", "")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.Convert(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// The full name of the type being converted from
		fmt.Printf("from-name: \"%s\"\n", data["from-name"])

		// The standard UTF-8 symbol used to represent the type being converted from
		fmt.Printf("from-symbol: \"%s\"\n", data["from-symbol"])

		// The type of the value being converted from
		fmt.Printf("from-type: \"%s\"\n", data["from-type"])

		// The value being converted from
		fmt.Printf("from-value: \"%s\"\n", data["from-value"])

		// If a historical conversion was made using the 'historical-date' request option this will contain
		// the exact date used for the conversion in ISO format: YYYY-MM-DD
		fmt.Printf("historical-date: \"%s\"\n", data["historical-date"])

		// The result of the conversion in string format
		fmt.Printf("result: \"%s\"\n", data["result"])

		// The result of the conversion as a floating-point number
		fmt.Printf("result-float: %.f\n", data["result-float"])

		// The full name of the type being converted to
		fmt.Printf("to-name: \"%s\"\n", data["to-name"])

		// The standard UTF-8 symbol used to represent the type being converted to
		fmt.Printf("to-symbol: \"%s\"\n", data["to-symbol"])

		// The type being converted to
		fmt.Printf("to-type: \"%s\"\n", data["to-type"])

		// True if the conversion was successful and produced a valid result
		fmt.Printf("valid: %t\n", data["valid"])

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
