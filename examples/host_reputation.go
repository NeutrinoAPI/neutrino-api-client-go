package main

import (
	"fmt"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	params := make(url.Values, 3)

	// An IP address, domain name, FQDN or URL. If you supply a domain/URL it will be checked against
	// the URI DNSBL lists
	params.Add("host", "neutrinoapi.com")

	// Only check lists with this rating or better
	params.Add("list-rating", "3")

	// Only check these DNSBL zones/hosts. Multiple zones can be supplied as comma-separated values
	params.Add("zones", "")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.HostReputation(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// The IP address or host name
		fmt.Printf("host: \"%s\"\n", data["host"])

		// Is this host blacklisted
		fmt.Printf("is-listed: %t\n", data["is-listed"])

		// The number of DNSBLs the host is listed on
		fmt.Printf("list-count: %.f\n", data["list-count"])

		// Array of objects for each DNSBL checked
		fmt.Printf("lists:\n")
		lists := data["lists"].([]interface{})
		for _, item := range lists {
			itemMap := item.(map[string]interface{})

			// True if the host is currently black-listed
			fmt.Printf("    is-listed: %t\n", itemMap["is-listed"])

			// The hostname of the DNSBL
			fmt.Printf("    list-host: \"%s\"\n", itemMap["list-host"])

			// The name of the DNSBL
			fmt.Printf("    list-name: \"%s\"\n", itemMap["list-name"])

			// The list rating [1-3] with 1 being the best rating and 3 the lowest rating
			fmt.Printf("    list-rating: %.f\n", itemMap["list-rating"])

			// The DNSBL server response time in milliseconds
			fmt.Printf("    response-time: %.f\n", itemMap["response-time"])

			// The specific return code for this listing (only set if listed)
			fmt.Printf("    return-code: \"%s\"\n", itemMap["return-code"])

			// The TXT record returned for this listing (only set if listed)
			fmt.Printf("    txt-record: \"%s\"\n", itemMap["txt-record"])
		fmt.Println()
		}

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
