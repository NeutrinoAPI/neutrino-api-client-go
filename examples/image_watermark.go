package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	file, err := ioutil.TempFile("", "image-watermark-*.png")
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

	params := make(url.Values, 7)

	// The output image format, can be either png or jpg
	params.Add("format", "png")

	// If set resize the resulting image to this width (in px) while preserving aspect ratio
	params.Add("width", "")

	// The URL or Base64 encoded Data URL for the source image (you can also upload an image file
	// directly in which case this field is ignored)
	params.Add("image-url", "https://www.neutrinoapi.com/img/LOGO.png")

	// The position of the watermark image, possible values are: center, top-left, top-center,
	// top-right, bottom-left, bottom-center, bottom-right
	params.Add("position", "center")

	// The URL or Base64 encoded Data URL for the watermark image (you can also upload an image file
	// directly in which case this field is ignored)
	params.Add("watermark-url", "https://www.neutrinoapi.com/img/icons/security.png")

	// The opacity of the watermark (0 to 100)
	params.Add("opacity", "50")

	// If set resize the resulting image to this height (in px) while preserving aspect ratio
	params.Add("height", "")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.ImageWatermark(params, file)
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
