package main

import (
	"fmt"
	"net/url"
	. "neutrino_api_client_go/pkg"
	"os"
)

func main() {
	params := make(url.Values, 7)

	// The user-agent string to lookup. For client hints use the 'UA' header or the JSON data directly
	// from 'navigator.userAgentData.brands' or 'navigator.userAgentData.getHighEntropyValues()'
	params.Add("ua", "Mozilla/5.0 (Linux; Android 11; SM-G9980U1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.101 Mobile Safari/537.36")

	// For client hints this corresponds to the 'UA-Full-Version' header or 'uaFullVersion' from
	// NavigatorUAData
	params.Add("ua-version", "")

	// For client hints this corresponds to the 'UA-Platform' header or 'platform' from NavigatorUAData
	params.Add("ua-platform", "")

	// For client hints this corresponds to the 'UA-Platform-Version' header or 'platformVersion' from
	// NavigatorUAData
	params.Add("ua-platform-version", "")

	// For client hints this corresponds to the 'UA-Mobile' header or 'mobile' from NavigatorUAData
	params.Add("ua-mobile", "")

	// For client hints this corresponds to the 'UA-Model' header or 'model' from NavigatorUAData. You
	// can also use this parameter to lookup a device directly by its model name, model code or hardware
	// code, on android you can get the model name from:
	// https://developer.android.com/reference/android/os/Build.html#MODEL
	params.Add("device-model", "")

	// This parameter is only used in combination with 'device-model' when doing direct device lookups
	// without any user-agent data. Set this to the brand or manufacturer name, this is required for
	// accurate device detection with ambiguous model names. On android you can get the device brand
	// from: https://developer.android.com/reference/android/os/Build#MANUFACTURER
	params.Add("device-brand", "")

	neutrinoAPIClient := NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
	response := neutrinoAPIClient.UALookup(params)
	if response.IsOK() {
		data := response.Data
		fmt.Println("API Response OK:")

		// If the client is a web browser which underlying browser engine does it use
		fmt.Printf("browser-engine: \"%s\"\n", data["browser-engine"])

		// If the client is a web browser which year was this browser version released
		fmt.Printf("browser-release: \"%s\"\n", data["browser-release"])

		// The device brand / manufacturer
		fmt.Printf("device-brand: \"%s\"\n", data["device-brand"])

		// The device display height in CSS 'px'
		fmt.Printf("device-height-px: %.f\n", data["device-height-px"])

		// The device model
		fmt.Printf("device-model: \"%s\"\n", data["device-model"])

		// The device model code
		fmt.Printf("device-model-code: \"%s\"\n", data["device-model-code"])

		// The device display pixel ratio (the ratio of the resolution in physical pixels to the resolution
		// in CSS pixels)
		fmt.Printf("device-pixel-ratio: %.f\n", data["device-pixel-ratio"])

		// The device display PPI (pixels per inch)
		fmt.Printf("device-ppi: %.f\n", data["device-ppi"])

		// The average device price on release in USD
		fmt.Printf("device-price: %.f\n", data["device-price"])

		// The year when this device model was released
		fmt.Printf("device-release: \"%s\"\n", data["device-release"])

		// The device display resolution in physical pixels (e.g. 720x1280)
		fmt.Printf("device-resolution: \"%s\"\n", data["device-resolution"])

		// The device display width in CSS 'px'
		fmt.Printf("device-width-px: %.f\n", data["device-width-px"])

		// Is this a mobile device (e.g. a phone or tablet)
		fmt.Printf("is-mobile: %t\n", data["is-mobile"])

		// Is this a WebView / embedded software client
		fmt.Printf("is-webview: %t\n", data["is-webview"])

		// The client software name
		fmt.Printf("name: \"%s\"\n", data["name"])

		// The full operating system name
		fmt.Printf("os: \"%s\"\n", data["os"])

		// The operating system family. The major OS families are: Android, Windows, macOS, iOS, Linux
		fmt.Printf("os-family: \"%s\"\n", data["os-family"])

		// The operating system full version
		fmt.Printf("os-version: \"%s\"\n", data["os-version"])

		// The operating system major version
		fmt.Printf("os-version-major: \"%s\"\n", data["os-version-major"])

		// The user agent type, possible values are:
		// • desktop
		// • phone
		// • tablet
		// • wearable
		// • tv
		// • console
		// • email
		// • library
		// • robot
		// • unknown
		fmt.Printf("type: \"%s\"\n", data["type"])

		// The user agent string
		fmt.Printf("ua: \"%s\"\n", data["ua"])

		// The client software full version
		fmt.Printf("version: \"%s\"\n", data["version"])

		// The client software major version
		fmt.Printf("version-major: \"%s\"\n", data["version-major"])

	} else {
		// You should handle this gracefully!
		_, _ = fmt.Fprintln(os.Stderr, "API Error:", *response.ErrorMessage+",", "Error Code:", fmt.Sprintf("%d,", *response.ErrorCode), "HTTP Status Code:", response.HttpStatusCode)
		if response.ErrorCause != nil {
			_, _ = fmt.Fprintln(os.Stderr, *response.ErrorCause)
		}
	}
}
