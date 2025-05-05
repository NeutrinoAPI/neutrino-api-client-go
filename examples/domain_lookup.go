package main

import (
	"fmt"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
	"strings"
)

func main() {
	params := make(url.Values, 2)

	// A domain name, hostname, FQDN, URL, HTML link or email address to lookup
	params.Add("host", "neutrinoapi.com")

	// For domains that we have never seen before then perform various live checks and realtime
	// reconnaissance. NOTE: this option may add additional non-deterministic delay to the request, if
	// you require consistently fast API response times or just want to check our domain blocklists then
	// you can disable this option
	params.Add("live", "true")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.DomainLookup(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// The number of days since the domain was registered. A domain age of under 90 days is generally
		// considered to be potentially risky. A value of 0 indicates no registration date was found for
		// this domain
		fmt.Printf("age: %.f\n", data["age"])

		// An array of strings indicating which blocklist categories this domain is listed on. Current
		// possible values are:
		// • phishing - Domain has recently been hosting phishing links or involved in the sending of
		//   phishing messages
		// • malware - Domain has recently been hosting malware or involved in the distribution of malware
		// • spam - Domain has recently been sending spam either directly or indirectly
		// • anonymizer - Domain is involved in anonymizer activity such as disposable email, hosting
		//   proxies or tor services
		// • nefarious - Domain is involved in nefarious or malicious activity such as hacking, fraud or
		//   other abusive behavior
		blocklists := strings.Fields(fmt.Sprint(data["blocklists"]))
		fmt.Printf("blocklists: %s\n", strings.Join(blocklists, ", "))

		// The primary domain of the DNS provider for this domain
		fmt.Printf("dns-provider: \"%s\"\n", data["dns-provider"])

		// The primary domain name excluding any subdomains. This is also referred to as the second-level
		// domain (SLD)
		fmt.Printf("domain: \"%s\"\n", data["domain"])

		// The fully qualified domain name (FQDN)
		fmt.Printf("fqdn: \"%s\"\n", data["fqdn"])

		// This domain is hosting adult content such as porn, webcams, escorts, etc
		fmt.Printf("is-adult: %t\n", data["is-adult"])

		// Is this domain under a government or military TLD
		fmt.Printf("is-gov: %t\n", data["is-gov"])

		// Consider this domain malicious as it is currently listed on at least 1 blocklist
		fmt.Printf("is-malicious: %t\n", data["is-malicious"])

		// Is this domain under an OpenNIC TLD
		fmt.Printf("is-opennic: %t\n", data["is-opennic"])

		// True if this domain is unseen and is currently being processed in the background. This field only
		// matters when the 'live' lookup setting has been explicitly disabled and indicates that not all
		// domain data my be present yet
		fmt.Printf("is-pending: %t\n", data["is-pending"])

		// Is the FQDN a subdomain of the primary domain
		fmt.Printf("is-subdomain: %t\n", data["is-subdomain"])

		// The primary domain of the email provider for this domain. An empty value indicates the domain has
		// no valid MX records
		fmt.Printf("mail-provider: \"%s\"\n", data["mail-provider"])

		// The domains estimated global traffic rank with the highest rank being 1. A value of 0 indicates
		// the domain is currently ranked outside of the top 1M of domains
		fmt.Printf("rank: %.f\n", data["rank"])

		// The ISO date this domain was registered or first seen on the internet. An empty value indicates
		// we could not reliably determine the date
		fmt.Printf("registered-date: \"%s\"\n", data["registered-date"])

		// The IANA registrar ID (0 if no registrar ID was found)
		fmt.Printf("registrar-id: %.f\n", data["registrar-id"])

		// The name of the domain registrar owning this domain
		fmt.Printf("registrar-name: \"%s\"\n", data["registrar-name"])

		// An array of objects containing details on which specific blocklist sensors have detected this
		// domain
		fmt.Printf("sensors:\n")
		sensors := data["sensors"].([]interface{})
		for _, item := range sensors {
			itemMap := item.(map[string]interface{})

			// The primary blocklist category this sensor belongs to
			fmt.Printf("    blocklist: \"%s\"\n", itemMap["blocklist"])

			// Contains details about the sensor source and what type of malicious activity was detected
			fmt.Printf("    description: \"%s\"\n", itemMap["description"])

			// The sensor ID. This is a permanent and unique ID for each sensor
			fmt.Printf("    id: %.f\n", itemMap["id"])
		fmt.Println()
		}

		// The top-level domain (TLD)
		fmt.Printf("tld: \"%s\"\n", data["tld"])

		// For a country code top-level domain (ccTLD) this will contain the associated ISO 2-letter country
		// code
		fmt.Printf("tld-cc: \"%s\"\n", data["tld-cc"])

		// True if a valid domain was found. For a domain to be considered valid it must be registered and
		// have valid DNS NS records
		fmt.Printf("valid: %t\n", data["valid"])

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
