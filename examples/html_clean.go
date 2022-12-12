package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	file, err := ioutil.TempFile("", "html-clean-*.txt")
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

	params := make(url.Values, 2)

	// The level of sanitization, possible values are: plain-text: reduce the content to plain text only
	// (no HTML tags at all) simple-text: allow only very basic text formatting tags like b, em, i,
	// strong, u basic-html: allow advanced text formatting and hyper links basic-html-with-images: same
	// as basic html but also allows image tags advanced-html: same as basic html with images but also
	// allows many more common HTML tags like table, ul, dl, pre
	params.Add("output-type", "plain-text")

	// The HTML content. This can be either a URL to load from, a file upload or an HTML content string
	params.Add("content", "<div>Some HTML to clean...</div><script>alert()</script>")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.HTMLClean(params, file)
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
