package main

import (
	"fmt"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	params := make(url.Values, 3)

	// The value to convert from (e.g. 10.95)
	params.Add("from-value", "100")

	// The type of the value to convert from (e.g. USD)
	params.Add("from-type", "USD")

	// The type to convert to (e.g. EUR)
	params.Add("to-type", "EUR")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.Convert(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// The type of the value being converted from
		fmt.Printf("from-type: \"%s\"\n", data["from-type"])

		// The value being converted from
		fmt.Printf("from-value: \"%s\"\n", data["from-value"])

		// The result of the conversion in string format
		fmt.Printf("result: \"%s\"\n", data["result"])

		// The result of the conversion as a floating-point number
		fmt.Printf("result-float: %.f\n", data["result-float"])

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
