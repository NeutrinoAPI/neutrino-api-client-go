package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	file, err := ioutil.TempFile("", "html-render-*.pdf")
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

	params := make(url.Values, 25)

	// Inject custom CSS into the HTML. e.g. 'body { background-color: red;}'
	params.Add("css", "")

	// The footer HTML to insert into each page. The following dynamic tags are supported: {date},
	// {title}, {url}, {pageNumber}, {totalPages}
	params.Add("footer", "")

	// The document title
	params.Add("title", "")

	// The HTML content. This can be either a URL to load from, a file upload (multipart/form-data) or
	// an HTML content string
	params.Add("content", "<h1>TEST DOCUMENT</h1><p>Hello, this is a test page...</p>")

	// Set the PDF page width explicitly (in mm)
	params.Add("page-width", "")

	// Timeout in seconds. Give up if still trying to load the HTML content after this number of seconds
	params.Add("timeout", "300")

	// Render the final document in grayscale
	params.Add("grayscale", "false")

	// The document left margin (in mm)
	params.Add("margin-left", "0")

	// Set the document page size, can be one of: A0 - A9, B0 - B10, Comm10E, DLE or Letter
	params.Add("page-size", "A4")

	// Ignore any TLS/SSL certificate errors
	params.Add("ignore-certificate-errors", "false")

	// Set the PDF page height explicitly (in mm)
	params.Add("page-height", "")

	// The document top margin (in mm)
	params.Add("margin-top", "0")

	// For image rendering set the background color in hexadecimal notation (e.g. #0000ff). For PNG
	// output the special value of 'transparent' can be used to create a transparent PNG
	params.Add("bg-color", "")

	// The document margin (in mm)
	params.Add("margin", "0")

	// If rendering to an image format (PNG or JPG) use this image width (in pixels)
	params.Add("image-width", "1024")

	// Which format to output, available options are: PDF, PNG, JPG
	params.Add("format", "PDF")

	// Set the zoom factor when rendering the page (2.0 for double size, 0.5 for half size)
	params.Add("zoom", "1")

	// The document right margin (in mm)
	params.Add("margin-right", "0")

	// Number of seconds to wait before rendering the page (can be useful for pages with animations etc)
	params.Add("delay", "0")

	// If rendering to an image format (PNG or JPG) use this image height (in pixels). The default is
	// automatic which dynamically sets the image height based on the content
	params.Add("image-height", "")

	// The header HTML to insert into each page. The following dynamic tags are supported: {date},
	// {title}, {url}, {pageNumber}, {totalPages}
	params.Add("header", "<div style='width: 100%; font-size: 8pt;'>{pageNumber} of {totalPages} - {date}</div>")

	// The document bottom margin (in mm)
	params.Add("margin-bottom", "0")

	// Set the document to landscape orientation
	params.Add("landscape", "false")

	// Execute JavaScript on the website. This parameter accepts JavaScript as either a string
	// containing JavaScript or for sending multiple separate statements a JSON array or POST array can
	// also be used. You can also use the following specially defined user interaction functions:
	// sleep(seconds); Just wait/sleep for the specified number of seconds. click('selector'); Click on
	// the first element matching the given selector. focus('selector'); Focus on the first element
	// matching the given selector. keys('characters'); Send the specified keyboard characters. Use
	// click() or focus() first to send keys to a specific element. enter(); Send the Enter key. tab();
	// Send the Tab key.
	params.Add("exec", "")

	// Override the browsers default user-agent string with this one
	params.Add("user-agent", "")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.HTMLRender(params, file)
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
