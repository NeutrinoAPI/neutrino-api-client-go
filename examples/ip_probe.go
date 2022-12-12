package main

import (
	"fmt"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
	"strings"
)

func main() {
	params := make(url.Values, 1)

	// IPv4 or IPv6 address
	params.Add("ip", "194.233.98.38")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.IPProbe(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// The age of the autonomous system (AS) in number of years since registration
		fmt.Printf("as-age: %.f\n", data["as-age"])

		// The autonomous system (AS) CIDR range
		fmt.Printf("as-cidr: \"%s\"\n", data["as-cidr"])

		// The autonomous system (AS) ISO 2-letter country code
		fmt.Printf("as-country-code: \"%s\"\n", data["as-country-code"])

		// The autonomous system (AS) ISO 3-letter country code
		fmt.Printf("as-country-code3: \"%s\"\n", data["as-country-code3"])

		// The autonomous system (AS) description / company name
		fmt.Printf("as-description: \"%s\"\n", data["as-description"])

		// Array of all the domains associated with the autonomous system (AS)
		asDomains := strings.Fields(fmt.Sprint(data["as-domains"]))
		fmt.Printf("as-domains: %s\n", strings.Join(asDomains, ", "))

		// The autonomous system (AS) number
		fmt.Printf("asn: \"%s\"\n", data["asn"])

		// Full city name (if detectable)
		fmt.Printf("city: \"%s\"\n", data["city"])

		// ISO 2-letter continent code
		fmt.Printf("continent-code: \"%s\"\n", data["continent-code"])

		// Full country name
		fmt.Printf("country: \"%s\"\n", data["country"])

		// ISO 2-letter country code
		fmt.Printf("country-code: \"%s\"\n", data["country-code"])

		// ISO 3-letter country code
		fmt.Printf("country-code3: \"%s\"\n", data["country-code3"])

		// ISO 4217 currency code associated with the country
		fmt.Printf("currency-code: \"%s\"\n", data["currency-code"])

		// The IPs host domain
		fmt.Printf("host-domain: \"%s\"\n", data["host-domain"])

		// The IPs full hostname (PTR)
		fmt.Printf("hostname: \"%s\"\n", data["hostname"])

		// The IP address
		fmt.Printf("ip: \"%s\"\n", data["ip"])

		// True if this is a bogon IP address such as a private network, local network or reserved address
		fmt.Printf("is-bogon: %t\n", data["is-bogon"])

		// True if this IP belongs to a hosting company. Note that this can still be true even if the
		// provider type is VPN/proxy, this occurs in the case that the IP is detected as both types
		fmt.Printf("is-hosting: %t\n", data["is-hosting"])

		// True if this IP belongs to an internet service provider. Note that this can still be true even if
		// the provider type is VPN/proxy, this occurs in the case that the IP is detected as both types
		fmt.Printf("is-isp: %t\n", data["is-isp"])

		// True if this IP ia a proxy
		fmt.Printf("is-proxy: %t\n", data["is-proxy"])

		// True if this is a IPv4 mapped IPv6 address
		fmt.Printf("is-v4-mapped: %t\n", data["is-v4-mapped"])

		// True if this is a IPv6 address. False if IPv4
		fmt.Printf("is-v6: %t\n", data["is-v6"])

		// True if this IP ia a VPN
		fmt.Printf("is-vpn: %t\n", data["is-vpn"])

		// A description of the provider (usually extracted from the providers website)
		fmt.Printf("provider-description: \"%s\"\n", data["provider-description"])

		// The domain name of the provider
		fmt.Printf("provider-domain: \"%s\"\n", data["provider-domain"])

		// The detected provider type, possible values are:
		// • isp - IP belongs to an internet service provider. This includes both mobile, home and
		//   business internet providers
		// • hosting - IP belongs to a hosting company. This includes website hosting, cloud computing
		//   platforms and colocation facilities
		// • vpn - IP belongs to a VPN provider
		// • proxy - IP belongs to a proxy service. This includes HTTP/SOCKS proxies and browser based
		//   proxies
		// • university - IP belongs to a university/college/campus
		// • government - IP belongs to a government department. This includes military facilities
		// • commercial - IP belongs to a commercial entity such as a corporate headquarters or company
		//   office
		// • unknown - could not identify the provider type
		fmt.Printf("provider-type: \"%s\"\n", data["provider-type"])

		// The website URL for the provider
		fmt.Printf("provider-website: \"%s\"\n", data["provider-website"])

		// Full region name (if detectable)
		fmt.Printf("region: \"%s\"\n", data["region"])

		// ISO 3166-2 region code (if detectable)
		fmt.Printf("region-code: \"%s\"\n", data["region-code"])

		// True if this is a valid IPv4 or IPv6 address
		fmt.Printf("valid: %t\n", data["valid"])

		// The domain of the VPN provider (may be empty if the VPN domain is not detectable)
		fmt.Printf("vpn-domain: \"%s\"\n", data["vpn-domain"])

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
