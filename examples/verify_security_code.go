package main

import (
	"fmt"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	params := make(url.Values, 2)

	// The security code to verify
	params.Add("security-code", "123456")

	// If set then enable additional brute-force protection by limiting the number of attempts by the
	// supplied value. This can be set to any unique identifier you would like to limit by, for example
	// a hash of the users email, phone number or IP address. Requests to this API will be ignored after
	// approximately 10 failed verification attempts
	params.Add("limit-by", "")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.VerifySecurityCode(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// True if the code is valid
		fmt.Printf("verified: %t\n", data["verified"])

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
