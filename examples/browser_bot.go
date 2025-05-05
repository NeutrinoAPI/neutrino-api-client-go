package main

import (
	"fmt"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	params := make(url.Values, 7)

	// Delay in seconds to wait before capturing any page data, executing selectors or JavaScript
	params.Add("delay", "3")

	// Ignore any TLS/SSL certificate errors and load the page anyway
	params.Add("ignore-certificate-errors", "false")

	// Extract content from the page DOM using this selector. Commonly known as a CSS selector, you can
	// find a good reference here
	params.Add("selector", ".button")

	// The URL to load
	params.Add("url", "https://www.neutrinoapi.com/")

	// Timeout in seconds. Give up if still trying to load the page after this number of seconds
	params.Add("timeout", "30")

	// Execute JavaScript on the website. This parameter accepts JavaScript as either a string
	// containing JavaScript or for sending multiple separate statements a JSON array or POST array can
	// also be used. If a statement returns any value it will be returned in the 'exec-results'
	// response. You can also use the following specially defined user interaction functions:
	// sleep(seconds); Just wait/sleep for the specified number of seconds. click('selector'); Click on
	// the first element matching the given selector. focus('selector'); Focus on the first element
	// matching the given selector. keys('characters'); Send the specified keyboard characters. Use
	// click() or focus() first to send keys to a specific element. enter(); Send the Enter key. tab();
	// Send the Tab key.
	params.Add("exec", "[click('#button-id'), sleep(1), click('.class'), keys('1234'), enter()]")

	// Override the browsers default user-agent string with this one
	params.Add("user-agent", "")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.BrowserBot(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// The complete raw, decompressed and decoded page content. Usually will be either HTML, JSON or XML
		fmt.Printf("content: \"%s\"\n", data["content"])

		// The size of the returned content in bytes
		fmt.Printf("content-size: %.f\n", data["content-size"])

		// Array containing all the elements matching the supplied selector
		fmt.Printf("elements:\n")
		elements := data["elements"].([]interface{})
		for _, item := range elements {
			itemMap := item.(map[string]interface{})

			// The 'class' attribute of the element
			fmt.Printf("    class: \"%s\"\n", itemMap["class"])

			// The 'href' attribute of the element
			fmt.Printf("    href: \"%s\"\n", itemMap["href"])

			// The raw HTML of the element
			fmt.Printf("    html: \"%s\"\n", itemMap["html"])

			// The 'id' attribute of the element
			fmt.Printf("    id: \"%s\"\n", itemMap["id"])

			// The plain-text content of the element with normalized whitespace
			fmt.Printf("    text: \"%s\"\n", itemMap["text"])
		fmt.Println()
		}

		// Contains the error message if an error has occurred ('is-error' will be true)
		fmt.Printf("error-message: \"%s\"\n", data["error-message"])

		// If you executed any JavaScript this array holds the results as objects
		fmt.Printf("exec-results:\n")
		execResults := data["exec-results"].([]interface{})
		for _, item := range execResults {
			itemMap := item.(map[string]interface{})

			// The result of the executed JavaScript statement. Will be empty if the statement returned nothing
			fmt.Printf("    result: \"%s\"\n", itemMap["result"])

			// The JavaScript statement that was executed
			fmt.Printf("    statement: \"%s\"\n", itemMap["statement"])
		fmt.Println()
		}

		// The redirected URL if the URL responded with an HTTP redirect
		fmt.Printf("http-redirect-url: \"%s\"\n", data["http-redirect-url"])

		// The HTTP status code the URL returned
		fmt.Printf("http-status-code: %.f\n", data["http-status-code"])

		// The HTTP status message the URL returned
		fmt.Printf("http-status-message: \"%s\"\n", data["http-status-message"])

		// True if an error has occurred loading the page. Check the 'error-message' field for details
		fmt.Printf("is-error: %t\n", data["is-error"])

		// True if the HTTP status is OK (200)
		fmt.Printf("is-http-ok: %t\n", data["is-http-ok"])

		// True if the URL responded with an HTTP redirect
		fmt.Printf("is-http-redirect: %t\n", data["is-http-redirect"])

		// True if the page is secured using TLS/SSL
		fmt.Printf("is-secure: %t\n", data["is-secure"])

		// True if a timeout occurred while loading the page. You can set the timeout with the request
		// parameter 'timeout'
		fmt.Printf("is-timeout: %t\n", data["is-timeout"])

		// The ISO 2-letter language code of the page. Extracted from either the HTML document or via HTTP
		// headers
		fmt.Printf("language-code: \"%s\"\n", data["language-code"])

		// The number of seconds taken to load the page (from initial request until DOM ready)
		fmt.Printf("load-time: %.f\n", data["load-time"])

		// The document MIME type
		fmt.Printf("mime-type: \"%s\"\n", data["mime-type"])

		// Map containing all the HTTP response headers the URL responded with
		fmt.Printf("response-headers: %s\n", data["response-headers"])

		// Map containing details of the TLS/SSL setup
		fmt.Printf("security-details: %s\n", data["security-details"])

		// The HTTP servers hostname (PTR/RDNS record)
		fmt.Printf("server-hostname: \"%s\"\n", data["server-hostname"])

		// The HTTP servers IP address
		fmt.Printf("server-ip: \"%s\"\n", data["server-ip"])

		// The document title
		fmt.Printf("title: \"%s\"\n", data["title"])

		// The requested URL. This may not be the same as the final destination URL, if the URL redirects
		// then it will be set in 'http-redirect-url' and 'is-http-redirect' will also be true
		fmt.Printf("url: \"%s\"\n", data["url"])

		// Structure of url-components
		fmt.Printf("url-components: %s\n", data["url-components"])

		// True if the URL supplied is valid
		fmt.Printf("url-valid: %t\n", data["url-valid"])

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
