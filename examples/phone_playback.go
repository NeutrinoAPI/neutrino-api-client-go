package main

import (
	"fmt"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	params := make(url.Values, 4)

	// The phone number to call. Must be in valid international format
	params.Add("number", "+12106100045")

	// Limit the total number of calls allowed to the supplied phone number, if the limit is reached
	// within the TTL then error code 14 will be returned
	params.Add("limit", "3")

	// A URL to a valid audio file. Accepted audio formats are:
	// • MP3
	// • WAV
	// • OGG You can use the following MP3 URL for testing:
	//   https://www.neutrinoapi.com/test-files/test1.mp3
	params.Add("audio-url", "https://www.neutrinoapi.com/test-files/test1.mp3")

	// Set the TTL in number of days that the 'limit' option will remember a phone number (the default
	// is 1 day and the maximum is 365 days)
	params.Add("limit-ttl", "1")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.PhonePlayback(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// True if the call is being made now
		fmt.Printf("calling: %t\n", data["calling"])

		// True if this a valid phone number
		fmt.Printf("number-valid: %t\n", data["number-valid"])

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
