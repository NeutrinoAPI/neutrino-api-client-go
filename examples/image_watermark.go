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

	params := make(url.Values, 9)

	// The resize mode to use, we support 3 main resizing modes:
	// • scale Resize to within the width and height specified while preserving aspect ratio. In this
	//   mode the width or height will be automatically adjusted to fit the aspect ratio
	// • pad Resize to exactly the width and height specified while preserving aspect ratio and pad
	//   any space left over. Any padded space will be filled in with the 'bg-color' value
	// • crop Resize to exactly the width and height specified while preserving aspect ratio and crop
	//   any space which fall outside the area. The cropping window is centered on the original image
	params.Add("resize-mode", "scale")

	// The output image format, can be either png or jpg
	params.Add("format", "png")

	// If set resize the resulting image to this width (in px)
	params.Add("width", "")

	// The URL or Base64 encoded Data URL for the source image. You can also upload an image file
	// directly using multipart/form-data
	params.Add("image-url", "https://www.neutrinoapi.com/img/LOGO.png")

	// The position of the watermark image, possible values are: center, top-left, top-center,
	// top-right, bottom-left, bottom-center, bottom-right
	params.Add("position", "center")

	// The URL or Base64 encoded Data URL for the watermark image. You can also upload an image file
	// directly using multipart/form-data
	params.Add("watermark-url", "https://www.neutrinoapi.com/img/icons/security.png")

	// The opacity of the watermark (0 to 100)
	params.Add("opacity", "50")

	// The image background color in hexadecimal notation (e.g. #0000ff). For PNG output the special
	// value of 'transparent' can also be used. For JPG output the default is black (#000000)
	params.Add("bg-color", "transparent")

	// If set resize the resulting image to this height (in px)
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
