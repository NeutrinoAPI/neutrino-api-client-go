package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	file, err := ioutil.TempFile("", "bin-list-download-*.png")
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

	// Include ISO 3-letter country codes and ISO 3-letter currency codes in the data. These will be
	// added to columns 10 and 11 respectively
	params.Add("include-iso3", "false")

	// Include 8-digit and higher BIN codes. This option includes all 6-digit BINs and all 8-digit and
	// higher BINs (including some 9, 10 and 11 digit BINs where available)
	params.Add("include-8digit", "false")

	// Include all BINs and all available fields in the CSV file (overrides any values set for
	// 'include-iso3' or 'include-8digit')
	params.Add("include-all", "false")

	// Set this option to 'gzip' to have the output file compressed using gzip
	params.Add("output-encoding", "")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.BINListDownload(params, file)
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
