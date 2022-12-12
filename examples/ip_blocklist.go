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

	// An IPv4 or IPv6 address. Accepts standard IP notation (with or without port number), CIDR
	// notation and IPv6 compressed notation. If multiple IPs are passed using comma-separated values
	// the first non-bogon address on the list will be checked
	params.Add("ip", "104.244.72.115")

	// Include public VPN provider IP addresses. NOTE: For more advanced VPN detection including the
	// ability to identify private and stealth VPNs use the IP Probe API
	params.Add("vpn-lookup", "false")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.IPBlocklist(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// An array of strings indicating which blocklist categories this IP is listed on
		blocklists := strings.Fields(fmt.Sprint(data["blocklists"]))
		fmt.Printf("blocklists: %s\n", strings.Join(blocklists, ", "))

		// The CIDR address for this listing (only set if the IP is listed)
		fmt.Printf("cidr: \"%s\"\n", data["cidr"])

		// The IP address
		fmt.Printf("ip: \"%s\"\n", data["ip"])

		// IP is hosting a malicious bot or is part of a botnet. This is a broad category which includes
		// brute-force crackers
		fmt.Printf("is-bot: %t\n", data["is-bot"])

		// IP has been flagged as a significant attack source by DShield (dshield.org)
		fmt.Printf("is-dshield: %t\n", data["is-dshield"])

		// IP is hosting an exploit finding bot or is running exploit scanning software
		fmt.Printf("is-exploit-bot: %t\n", data["is-exploit-bot"])

		// IP is part of a hijacked netblock or a netblock controlled by a criminal organization
		fmt.Printf("is-hijacked: %t\n", data["is-hijacked"])

		// Is this IP on a blocklist
		fmt.Printf("is-listed: %t\n", data["is-listed"])

		// IP is involved in distributing or is running malware
		fmt.Printf("is-malware: %t\n", data["is-malware"])

		// IP has been detected as an anonymous web proxy or anonymous HTTP proxy
		fmt.Printf("is-proxy: %t\n", data["is-proxy"])

		// IP address is hosting a spam bot, comment spamming or any other spamming type software
		fmt.Printf("is-spam-bot: %t\n", data["is-spam-bot"])

		// IP is running a hostile web spider / web crawler
		fmt.Printf("is-spider: %t\n", data["is-spider"])

		// IP is involved in distributing or is running spyware
		fmt.Printf("is-spyware: %t\n", data["is-spyware"])

		// IP is a Tor node or running a Tor related service
		fmt.Printf("is-tor: %t\n", data["is-tor"])

		// IP belongs to a public VPN provider (only set if the 'vpn-lookup' option is enabled)
		fmt.Printf("is-vpn: %t\n", data["is-vpn"])

		// The unix time when this IP was last seen on any blocklist. IPs are automatically removed after 7
		// days therefor this value will never be older than 7 days
		fmt.Printf("last-seen: %.f\n", data["last-seen"])

		// The number of blocklists the IP is listed on
		fmt.Printf("list-count: %.f\n", data["list-count"])

		// An array of objects containing details on which specific sensors detected the IP
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

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
