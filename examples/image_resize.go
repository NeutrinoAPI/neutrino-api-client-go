package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	file, err := ioutil.TempFile("", "image-resize-*.png")
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

	params := make(url.Values, 4)

	// The width to resize to (in px) while preserving aspect ratio
	params.Add("width", "32")

	// The output image format, can be either png or jpg
	params.Add("format", "png")

	// The URL or Base64 encoded Data URL for the source image (you can also upload an image file
	// directly in which case this field is ignored)
	params.Add("image-url", "https://www.neutrinoapi.com/img/LOGO.png")

	// The height to resize to (in px) while preserving aspect ratio. If you don't set this field then
	// the height will be automatic based on the requested width and images aspect ratio
	params.Add("height", "32")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.ImageResize(params, file)
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
