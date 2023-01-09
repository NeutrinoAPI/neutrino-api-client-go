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

	params := make(url.Values, 6)

	// The resize mode to use, we support 3 main resizing modes:
	// • scale Resize to within the width and height specified while preserving aspect ratio. In this
	//   mode the width or height will be automatically adjusted to fit the aspect ratio
	// • pad Resize to exactly the width and height specified while preserving aspect ratio and pad
	//   any space left over. Any padded space will be filled in with the 'bg-color' value
	// • crop Resize to exactly the width and height specified while preserving aspect ratio and crop
	//   any space which fall outside the area. The cropping window is centered on the original image
	params.Add("resize-mode", "scale")

	// The width to resize to (in px)
	params.Add("width", "32")

	// The output image format, can be either png or jpg
	params.Add("format", "png")

	// The URL or Base64 encoded Data URL for the source image. You can also upload an image file
	// directly using multipart/form-data
	params.Add("image-url", "https://www.neutrinoapi.com/img/LOGO.png")

	// The image background color in hexadecimal notation (e.g. #0000ff). For PNG output the special
	// value of 'transparent' can also be used. For JPG output the default is black (#000000)
	params.Add("bg-color", "transparent")

	// The height to resize to (in px). If you don't set this field then the height will be automatic
	// based on the requested width and image aspect ratio
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
