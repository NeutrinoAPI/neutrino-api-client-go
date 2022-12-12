package main

import (
	"fmt"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	params := make(url.Values, 2)

	// A phone number
	params.Add("number", "+12106100045")

	// ISO 2-letter country code, assume numbers are based in this country. If not set numbers are
	// assumed to be in international format (with or without the leading + sign)
	params.Add("country-code", "")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.HLRLookup(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// The phone number country
		fmt.Printf("country: \"%s\"\n", data["country"])

		// The number location as an ISO 2-letter country code
		fmt.Printf("country-code: \"%s\"\n", data["country-code"])

		// The number location as an ISO 3-letter country code
		fmt.Printf("country-code3: \"%s\"\n", data["country-code3"])

		// ISO 4217 currency code associated with the country
		fmt.Printf("currency-code: \"%s\"\n", data["currency-code"])

		// The currently used network/carrier name
		fmt.Printf("current-network: \"%s\"\n", data["current-network"])

		// The HLR lookup status, possible values are:
		// • ok - the HLR lookup was successful and the device is connected
		// • absent - the number was once registered but the device has been switched off or out of
		//   network range for some time
		// • unknown - the number is not known by the mobile network
		// • invalid - the number is not a valid mobile MSISDN number
		// • fixed-line - the number is a registered fixed-line not mobile
		// • voip - the number has been detected as a VOIP line
		// • failed - the HLR lookup has failed, we could not determine the real status of this number
		fmt.Printf("hlr-status: \"%s\"\n", data["hlr-status"])

		// Was the HLR lookup successful. If true then this is a working and registered cell-phone or mobile
		// device (SMS and phone calls will be delivered)
		fmt.Printf("hlr-valid: %t\n", data["hlr-valid"])

		// The mobile IMSI number (International Mobile Subscriber Identity)
		fmt.Printf("imsi: \"%s\"\n", data["imsi"])

		// The international calling code
		fmt.Printf("international-calling-code: \"%s\"\n", data["international-calling-code"])

		// The number represented in full international format
		fmt.Printf("international-number: \"%s\"\n", data["international-number"])

		// True if this is a mobile number (only true with 100% certainty, if the number type is unknown
		// this value will be false)
		fmt.Printf("is-mobile: %t\n", data["is-mobile"])

		// Has this number been ported to another network
		fmt.Printf("is-ported: %t\n", data["is-ported"])

		// Is this number currently roaming from its origin country
		fmt.Printf("is-roaming: %t\n", data["is-roaming"])

		// The number represented in local dialing format
		fmt.Printf("local-number: \"%s\"\n", data["local-number"])

		// The number location. Could be a city, region or country depending on the type of number
		fmt.Printf("location: \"%s\"\n", data["location"])

		// The mobile MCC number (Mobile Country Code)
		fmt.Printf("mcc: \"%s\"\n", data["mcc"])

		// The mobile MNC number (Mobile Network Code)
		fmt.Printf("mnc: \"%s\"\n", data["mnc"])

		// The mobile MSC number (Mobile Switching Center)
		fmt.Printf("msc: \"%s\"\n", data["msc"])

		// The mobile MSIN number (Mobile Subscription Identification Number)
		fmt.Printf("msin: \"%s\"\n", data["msin"])

		// The number type, possible values are:
		// • mobile
		// • fixed-line
		// • premium-rate
		// • toll-free
		// • voip
		// • unknown
		fmt.Printf("number-type: \"%s\"\n", data["number-type"])

		// True if this a valid phone number
		fmt.Printf("number-valid: %t\n", data["number-valid"])

		// The origin network/carrier name
		fmt.Printf("origin-network: \"%s\"\n", data["origin-network"])

		// The ported to network/carrier name (only set if the number has been ported)
		fmt.Printf("ported-network: \"%s\"\n", data["ported-network"])

		// If the number is currently roaming, the ISO 2-letter country code of the roaming in country
		fmt.Printf("roaming-country-code: \"%s\"\n", data["roaming-country-code"])

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
