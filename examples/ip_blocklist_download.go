package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	file, err := ioutil.TempFile("", "ip-blocklist-download-*.csv")
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

	// The data format. Can be either CSV or TXT
	params.Add("format", "csv")

	// Output IPs using CIDR notation. This option should be preferred but is off by default for
	// backwards compatibility
	params.Add("cidr", "false")

	// Output the IPv6 version of the blocklist, the default is to output IPv4 only. Note that this
	// option enables CIDR notation too as this is the only notation currently supported for IPv6
	params.Add("ip6", "false")

	// The category of IP addresses to include in the download file, possible values are:
	// • all - all IPs available on your current plan (excludes VPN providers for any plans lower than
	//   Tier 3)
	// • bot - all IPs hosting a malicious bot or part of a botnet. This is a broad category which
	//   includes brute-force crackers
	// • exploit-bot - all IPs hosting an exploit finding bot or running exploit scanning software
	// • hijacked - all IPs that are part of a hijacked netblock or a netblock controlled by a
	//   criminal organization
	// • malware - all IPs involved in distributing or running malware or spyware
	// • proxy - all IPs detected as an anonymous web proxy or anonymous HTTP proxy
	// • spam-bot - all IPs hosting a spam bot, comment spamming or any other spamming type software
	// • spider - all IPs running a hostile web spider / web crawler
	// • tor - all IPs that are Tor nodes or running a Tor related service
	// • vpn - all IPs belonging to public VPN providers (only available for Tier 3 or higher
	//   accounts)
	params.Add("category", "all")

	// Set this option to 'gzip' to have the output file compressed using gzip
	params.Add("output-encoding", "")

	// Do not download the file but just return the current files MurmurHash3 checksum. You can use this
	// feature to check if the file has changed since a previous check
	params.Add("checksum", "false")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.IPBlocklistDownload(params, file)
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
