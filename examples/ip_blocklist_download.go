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

	params := make(url.Values, 4)

	// The data format. Can be either CSV or TXT
	params.Add("format", "csv")

	// Include public VPN provider IP addresses, this option is only available for Tier 3 or higher
	// accounts. WARNING: This option will add at least an additional 8 million IP addresses to the
	// download if not using CIDR notation
	params.Add("include-vpn", "false")

	// Output IPs using CIDR notation. This option should be preferred but is off by default for
	// backwards compatibility
	params.Add("cidr", "false")

	// Output the IPv6 version of the blocklist, the default is to output IPv4 only. Note that this
	// option enables CIDR notation too as this is the only notation currently supported for IPv6
	params.Add("ip6", "false")

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
