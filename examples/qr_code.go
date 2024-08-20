package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	file, err := ioutil.TempFile("", "qr-code-*.png")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
        return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}(file)

	params := make(url.Values, 6)

	// The barcode format to output. Accepted formats are: qr, c128
	params.Add("code-format", "qr")

	// The width of the QR code (in px)
	params.Add("width", "256")

	// The QR code foreground color
	params.Add("fg-color", "#000000")

	// The QR code background color
	params.Add("bg-color", "#ffffff")

	// The content to encode into the QR code (e.g. a URL or a phone number)
	params.Add("content", "https://www.neutrinoapi.com/signup/")

	// The height of the QR code (in px)
	params.Add("height", "256")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.QRCode(params, file)
	if response.IsOK() {
		fmt.Printf("API Response OK, output saved to: %s\n", *response.File)
	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
