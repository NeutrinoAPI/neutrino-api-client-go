package main

import (
	"fmt"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	params := make(url.Values, 5)

	// The URL to probe
	params.Add("url", "https://www.neutrinoapi.com/")

	// If this URL responds with html, text, json or xml then return the response. This option is useful
	// if you want to perform further processing on the URL content (e.g. with the HTML Extract or HTML
	// Clean APIs)
	params.Add("fetch-content", "false")

	// Ignore any TLS/SSL certificate errors and load the URL anyway
	params.Add("ignore-certificate-errors", "false")

	// Timeout in seconds. Give up if still trying to load the URL after this number of seconds
	params.Add("timeout", "60")

	// If the request fails for any reason try again this many times
	params.Add("retry", "0")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.URLInfo(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// The actual content this URL responded with. Only set if the 'fetch-content' option was used
		fmt.Printf("content: \"%s\"\n", data["content"])

		// The encoding format the URL uses
		fmt.Printf("content-encoding: \"%s\"\n", data["content-encoding"])

		// The size of the URL content in bytes
		fmt.Printf("content-size: %.f\n", data["content-size"])

		// The content-type this URL serves
		fmt.Printf("content-type: \"%s\"\n", data["content-type"])

		// True if this URL responded with an HTTP OK (200) status
		fmt.Printf("http-ok: %t\n", data["http-ok"])

		// True if this URL responded with an HTTP redirect
		fmt.Printf("http-redirect: %t\n", data["http-redirect"])

		// The HTTP status code this URL responded with. An HTTP status of 0 indicates a network level issue
		fmt.Printf("http-status: %.f\n", data["http-status"])

		// The HTTP status message assoicated with the status code
		fmt.Printf("http-status-message: \"%s\"\n", data["http-status-message"])

		// True if an error occurred while loading the URL. This includes network errors, TLS errors and
		// timeouts
		fmt.Printf("is-error: %t\n", data["is-error"])

		// True if a timeout occurred while loading the URL. You can set the timeout with the request
		// parameter 'timeout'
		fmt.Printf("is-timeout: %t\n", data["is-timeout"])

		// The ISO 2-letter language code of the page. Extracted from either the HTML document or via HTTP
		// headers
		fmt.Printf("language-code: \"%s\"\n", data["language-code"])

		// The time taken to load the URL content in seconds
		fmt.Printf("load-time: %.f\n", data["load-time"])

		// A key-value map of the URL query paramaters
		fmt.Printf("query: %s\n", data["query"])

		// Is this URL actually serving real content
		fmt.Printf("real: %t\n", data["real"])

		// The servers IP geo-location: full city name (if detectable)
		fmt.Printf("server-city: \"%s\"\n", data["server-city"])

		// The servers IP geo-location: full country name
		fmt.Printf("server-country: \"%s\"\n", data["server-country"])

		// The servers IP geo-location: ISO 2-letter country code
		fmt.Printf("server-country-code: \"%s\"\n", data["server-country-code"])

		// The servers hostname (PTR record)
		fmt.Printf("server-hostname: \"%s\"\n", data["server-hostname"])

		// The IP address of the server hosting this URL
		fmt.Printf("server-ip: \"%s\"\n", data["server-ip"])

		// The name of the server software hosting this URL
		fmt.Printf("server-name: \"%s\"\n", data["server-name"])

		// The servers IP geo-location: full region name (if detectable)
		fmt.Printf("server-region: \"%s\"\n", data["server-region"])

		// The document title
		fmt.Printf("title: \"%s\"\n", data["title"])

		// The fully qualified URL. This may be different to the URL requested if http-redirect is true
		fmt.Printf("url: \"%s\"\n", data["url"])

		// The URL path
		fmt.Printf("url-path: \"%s\"\n", data["url-path"])

		// The URL port
		fmt.Printf("url-port: %.f\n", data["url-port"])

		// The URL protocol, usually http or https
		fmt.Printf("url-protocol: \"%s\"\n", data["url-protocol"])

		// Is this a valid well-formed URL
		fmt.Printf("valid: %t\n", data["valid"])

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
