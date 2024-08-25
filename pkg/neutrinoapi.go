package neutrinoapi

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"syscall"
	"time"
)

// Servers
const (
	MulticloudEndpoint = "https://neutrinoapi.net/"
	AwsEndpoint = "https://aws.neutrinoapi.net/"
	GcpEndpoint = "https://gcp.neutrinoapi.net/"
	BackupEndpoint = "https://neutrinoapi.com/"
	ConnectTimeoutInSeconds = 10
)

// Make a request to the Neutrino API
type Client struct {
	UserID  string
	APIKey  string
	BaseURL string
}

// NewNeutrinoAPIClient - initializer for NeutrinoAPIClient
func NewNeutrinoAPIClient(userID string, APIKey string) *Client {
	return &Client{userID, APIKey, MulticloudEndpoint}
}

// NewNeutrinoAPIClientWithBaseURL - initializer for NeutrinoAPIClient
func NewNeutrinoAPIClientWithBaseURL(userID string, APIKey string, baseURL string) *Client {
	return &Client{userID, APIKey, baseURL}
}

// BadWordFilter - Detect bad words, swear words and profanity in a given text
//
// The parameters this API accepts are:
// * censor-character - The character to use to censor out the bad words found
// * catalog - Which catalog of bad words to use
// * content - The content to scan
//
// link: https://www.neutrinoapi.com/api/bad-word-filter
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) BadWordFilter(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("POST", "bad-word-filter", params, nil, 30)
}

// BINListDownload - Download our entire BIN database for direct use on your own systems
//
// The parameters this API accepts are:
// * include-iso3 - Include ISO 3-letter country codes and ISO 3-letter currency codes in the data
// * include-8digit - Include 8-digit and higher BIN codes
// * include-all - Include all BINs and all available fields in the CSV file (overrides any values set for 'include-iso3' or 'include-8digit')
// * output-encoding - Set this option to 'gzip' to have the output file compressed using gzip
//
// link: https://www.neutrinoapi.com/api/bin-list-download
// param: params net/url.Values type, a collection of key/value pairs
// param: file *os.File, where to save the response
// returns *APIResponse
func (neutrinoAPIClient Client) BINListDownload(params url.Values, file *os.File) *APIResponse {
	return neutrinoAPIClient.ExecRequest("POST", "bin-list-download", params, file, 30)
}

// BINLookup - Perform a BIN (Bank Identification Number) or IIN (Issuer Identification Number) lookup
//
// The parameters this API accepts are:
// * bin-number - The BIN or IIN number
// * customer-ip - Pass in the customers IP address and we will return some extra information about them
//
// link: https://www.neutrinoapi.com/api/bin-lookup
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) BINLookup(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("GET", "bin-lookup", params, nil, 10)
}

// BrowserBot - Browser bot can extract content, interact with keyboard and mouse events, and execute JavaScript on a website
//
// The parameters this API accepts are:
// * delay - Delay in seconds to wait before capturing any page data
// * ignore-certificate-errors - Ignore any TLS/SSL certificate errors and load the page anyway
// * selector - Extract content from the page DOM using this selector
// * url - The URL to load
// * timeout - Timeout in seconds
// * exec - Execute JavaScript on the website
// * user-agent - Override the browsers default user-agent string with this one
//
// link: https://www.neutrinoapi.com/api/browser-bot
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) BrowserBot(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("POST", "browser-bot", params, nil, 300)
}

// Convert - A currency and unit conversion tool
//
// The parameters this API accepts are:
// * from-value - The value to convert from (e.g. 10.95)
// * from-type - The type of the value to convert from (e.g. USD)
// * to-type - The type to convert to (e.g. EUR)
//
// link: https://www.neutrinoapi.com/api/convert
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) Convert(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("GET", "convert", params, nil, 10)
}

// DomainLookup - Retrieve domain name details and detect potentially malicious or dangerous domains
//
// The parameters this API accepts are:
// * host - A domain name
// * live - For domains that we have never seen before then perform various live checks and realtime reconnaissance
//
// link: https://www.neutrinoapi.com/api/domain-lookup
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) DomainLookup(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("GET", "domain-lookup", params, nil, 120)
}

// EmailValidate - Parse, validate and clean an email address
//
// The parameters this API accepts are:
// * email - An email address
// * fix-typos - Automatically attempt to fix typos in the address
//
// link: https://www.neutrinoapi.com/api/email-validate
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) EmailValidate(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("GET", "email-validate", params, nil, 30)
}

// EmailVerify - SMTP based email address verification
//
// The parameters this API accepts are:
// * email - An email address
// * fix-typos - Automatically attempt to fix typos in the address
//
// link: https://www.neutrinoapi.com/api/email-verify
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) EmailVerify(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("GET", "email-verify", params, nil, 120)
}

// GeocodeAddress - Geocode an address, partial address or just the name of a place
//
// The parameters this API accepts are:
// * address - The full address
// * house-number - The house/building number to locate
// * street - The street/road name to locate
// * city - The city/town name to locate
// * county - The county/region name to locate
// * state - The state name to locate
// * postal-code - The postal code to locate
// * country-code - Limit result to this country (the default is no country bias)
// * language-code - The language to display results in
// * fuzzy-search - If no matches are found for the given address
//
// link: https://www.neutrinoapi.com/api/geocode-address
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) GeocodeAddress(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("GET", "geocode-address", params, nil, 30)
}

// GeocodeReverse - Convert a geographic coordinate (latitude and longitude) into a real world address
//
// The parameters this API accepts are:
// * latitude - The location latitude in decimal degrees format
// * longitude - The location longitude in decimal degrees format
// * language-code - The language to display results in
// * zoom - The zoom level to respond with: address - the most precise address available street - the street level city - the city level state - the state level country - the country level 
//
// link: https://www.neutrinoapi.com/api/geocode-reverse
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) GeocodeReverse(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("GET", "geocode-reverse", params, nil, 30)
}

// HLRLookup - Connect to the global mobile cellular network and retrieve the status of a mobile device
//
// The parameters this API accepts are:
// * number - A phone number
// * country-code - ISO 2-letter country code
//
// link: https://www.neutrinoapi.com/api/hlr-lookup
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) HLRLookup(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("GET", "hlr-lookup", params, nil, 30)
}

// HostReputation - Check the reputation of an IP address, domain name or URL against a comprehensive list of blacklists and blocklists
//
// The parameters this API accepts are:
// * host - An IP address
// * list-rating - Only check lists with this rating or better
// * zones - Only check these DNSBL zones/hosts
//
// link: https://www.neutrinoapi.com/api/host-reputation
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) HostReputation(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("GET", "host-reputation", params, nil, 120)
}

// HTMLClean - Clean and sanitize untrusted HTML
//
// The parameters this API accepts are:
// * output-type - The level of sanitization
// * content - The HTML content
//
// link: https://www.neutrinoapi.com/api/html-clean
// param: params net/url.Values type, a collection of key/value pairs
// param: file *os.File, where to save the response
// returns *APIResponse
func (neutrinoAPIClient Client) HTMLClean(params url.Values, file *os.File) *APIResponse {
	return neutrinoAPIClient.ExecRequest("POST", "html-clean", params, file, 30)
}

// HTMLRender - Render HTML content to PDF, JPG or PNG
//
// The parameters this API accepts are:
// * css - Inject custom CSS into the HTML
// * footer - The footer HTML to insert into each page
// * title - The document title
// * content - The HTML content
// * page-width - Set the PDF page width explicitly (in mm)
// * timeout - Timeout in seconds
// * grayscale - Render the final document in grayscale
// * margin-left - The document left margin (in mm)
// * page-size - Set the document page size
// * ignore-certificate-errors - Ignore any TLS/SSL certificate errors
// * page-height - Set the PDF page height explicitly (in mm)
// * margin-top - The document top margin (in mm)
// * bg-color - For image rendering set the background color in hexadecimal notation (e.g. #0000ff)
// * margin - The document margin (in mm)
// * image-width - If rendering to an image format (PNG or JPG) use this image width (in pixels)
// * format - Which format to output
// * zoom - Set the zoom factor when rendering the page (2.0 for double size
// * margin-right - The document right margin (in mm)
// * delay - Number of seconds to wait before rendering the page (can be useful for pages with animations etc)
// * image-height - If rendering to an image format (PNG or JPG) use this image height (in pixels)
// * header - The header HTML to insert into each page
// * margin-bottom - The document bottom margin (in mm)
// * landscape - Set the document to landscape orientation
// * exec - Execute JavaScript on the website
// * user-agent - Override the browsers default user-agent string with this one
//
// link: https://www.neutrinoapi.com/api/html-render
// param: params net/url.Values type, a collection of key/value pairs
// param: file *os.File, where to save the response
// returns *APIResponse
func (neutrinoAPIClient Client) HTMLRender(params url.Values, file *os.File) *APIResponse {
	return neutrinoAPIClient.ExecRequest("POST", "html-render", params, file, 300)
}

// ImageResize - Resize an image and output as either JPEG or PNG
//
// The parameters this API accepts are:
// * resize-mode - The resize mode to use
// * width - The width to resize to (in px)
// * format - The output image format
// * image-url - The URL or Base64 encoded Data URL for the source image
// * bg-color - The image background color in hexadecimal notation (e.g. #0000ff)
// * height - The height to resize to (in px)
//
// link: https://www.neutrinoapi.com/api/image-resize
// param: params net/url.Values type, a collection of key/value pairs
// param: file *os.File, where to save the response
// returns *APIResponse
func (neutrinoAPIClient Client) ImageResize(params url.Values, file *os.File) *APIResponse {
	return neutrinoAPIClient.ExecRequest("POST", "image-resize", params, file, 20)
}

// ImageWatermark - Watermark one image with another image
//
// The parameters this API accepts are:
// * resize-mode - The resize mode to use
// * format - The output image format
// * width - If set resize the resulting image to this width (in px)
// * image-url - The URL or Base64 encoded Data URL for the source image
// * position - The position of the watermark image
// * watermark-url - The URL or Base64 encoded Data URL for the watermark image
// * opacity - The opacity of the watermark (0 to 100)
// * bg-color - The image background color in hexadecimal notation (e.g. #0000ff)
// * height - If set resize the resulting image to this height (in px)
//
// link: https://www.neutrinoapi.com/api/image-watermark
// param: params net/url.Values type, a collection of key/value pairs
// param: file *os.File, where to save the response
// returns *APIResponse
func (neutrinoAPIClient Client) ImageWatermark(params url.Values, file *os.File) *APIResponse {
	return neutrinoAPIClient.ExecRequest("POST", "image-watermark", params, file, 20)
}

// IPBlocklist - The IP Blocklist API will detect potentially malicious or dangerous IP addresses
//
// The parameters this API accepts are:
// * ip - An IPv4 or IPv6 address
// * vpn-lookup - Include public VPN provider IP addresses
//
// link: https://www.neutrinoapi.com/api/ip-blocklist
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) IPBlocklist(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("GET", "ip-blocklist", params, nil, 10)
}

// IPBlocklistDownload - This API is a direct feed to our IP blocklist data
//
// The parameters this API accepts are:
// * format - The data format
// * cidr - Output IPs using CIDR notation
// * ip6 - Output the IPv6 version of the blocklist
// * category - The category of IP addresses to include in the download file
// * output-encoding - Set this option to 'gzip' to have the output file compressed using gzip
// * checksum - Do not download the file but just return the current files MurmurHash3 checksum
//
// link: https://www.neutrinoapi.com/api/ip-blocklist-download
// param: params net/url.Values type, a collection of key/value pairs
// param: file *os.File, where to save the response
// returns *APIResponse
func (neutrinoAPIClient Client) IPBlocklistDownload(params url.Values, file *os.File) *APIResponse {
	return neutrinoAPIClient.ExecRequest("POST", "ip-blocklist-download", params, file, 30)
}

// IPInfo - Get location information about an IP address and do reverse DNS (PTR) lookups
//
// The parameters this API accepts are:
// * ip - An IPv4 or IPv6 address
// * reverse-lookup - Do a reverse DNS (PTR) lookup
//
// link: https://www.neutrinoapi.com/api/ip-info
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) IPInfo(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("GET", "ip-info", params, nil, 10)
}

// IPProbe - Execute a realtime network probe against an IPv4 or IPv6 address
//
// The parameters this API accepts are:
// * ip - An IPv4 or IPv6 address
//
// link: https://www.neutrinoapi.com/api/ip-probe
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) IPProbe(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("GET", "ip-probe", params, nil, 120)
}

// PhonePlayback - Make an automated call to any valid phone number and playback an audio message
//
// The parameters this API accepts are:
// * number - The phone number to call
// * limit - Limit the total number of calls allowed to the supplied phone number
// * audio-url - A URL to a valid audio file
// * limit-ttl - Set the TTL in number of days that the 'limit' option will remember a phone number (the default is 1 day and the maximum is 365 days)
//
// link: https://www.neutrinoapi.com/api/phone-playback
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) PhonePlayback(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("POST", "phone-playback", params, nil, 30)
}

// PhoneValidate - Parse, validate and get location information about a phone number
//
// The parameters this API accepts are:
// * number - A phone number
// * country-code - ISO 2-letter country code
// * ip - Pass in a users IP address and we will assume numbers are based in the country of the IP address
//
// link: https://www.neutrinoapi.com/api/phone-validate
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) PhoneValidate(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("GET", "phone-validate", params, nil, 10)
}

// PhoneVerify - Make an automated call to any valid phone number and playback a unique security code
//
// The parameters this API accepts are:
// * number - The phone number to send the verification code to
// * country-code - ISO 2-letter country code
// * security-code - Pass in your own security code
// * language-code - The language to playback the verification code in
// * code-length - The number of digits to use in the security code (between 4 and 12)
// * limit - Limit the total number of calls allowed to the supplied phone number
// * playback-delay - The delay in milliseconds between the playback of each security code
// * limit-ttl - Set the TTL in number of days that the 'limit' option will remember a phone number (the default is 1 day and the maximum is 365 days)
//
// link: https://www.neutrinoapi.com/api/phone-verify
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) PhoneVerify(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("POST", "phone-verify", params, nil, 30)
}

// QRCode - Generate a QR code as a PNG image
//
// The parameters this API accepts are:
// * code-format - The barcode format to output
// * width - The width of the QR code (in px)
// * fg-color - The QR code foreground color
// * bg-color - The QR code background color
// * content - The content to encode into the QR code (e.g. a URL or a phone number)
// * height - The height of the QR code (in px)
//
// link: https://www.neutrinoapi.com/api/qr-code
// param: params net/url.Values type, a collection of key/value pairs
// param: file *os.File, where to save the response
// returns *APIResponse
func (neutrinoAPIClient Client) QRCode(params url.Values, file *os.File) *APIResponse {
	return neutrinoAPIClient.ExecRequest("POST", "qr-code", params, file, 20)
}

// SMSVerify - Send a unique security code to any mobile device via SMS
//
// The parameters this API accepts are:
// * number - The phone number to send a verification code to
// * country-code - ISO 2-letter country code
// * security-code - Pass in your own security code
// * language-code - The language to send the verification code in
// * code-length - The number of digits to use in the security code (must be between 4 and 12)
// * limit - Limit the total number of SMS allowed to the supplied phone number
// * brand-name - Set a custom brand or product name in the verification message
// * limit-ttl - Set the TTL in number of days that the 'limit' option will remember a phone number (the default is 1 day and the maximum is 365 days)
//
// link: https://www.neutrinoapi.com/api/sms-verify
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) SMSVerify(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("POST", "sms-verify", params, nil, 30)
}

// UALookup - Parse, validate and get detailed user-agent information from a user agent string or from client hints
//
// The parameters this API accepts are:
// * ua - The user-agent string to lookup
// * ua-version - For client hints this corresponds to the 'UA-Full-Version' header or 'uaFullVersion' from NavigatorUAData
// * ua-platform - For client hints this corresponds to the 'UA-Platform' header or 'platform' from NavigatorUAData
// * ua-platform-version - For client hints this corresponds to the 'UA-Platform-Version' header or 'platformVersion' from NavigatorUAData
// * ua-mobile - For client hints this corresponds to the 'UA-Mobile' header or 'mobile' from NavigatorUAData
// * device-model - For client hints this corresponds to the 'UA-Model' header or 'model' from NavigatorUAData
// * device-brand - This parameter is only used in combination with 'device-model' when doing direct device lookups without any user-agent data
//
// link: https://www.neutrinoapi.com/api/ua-lookup
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) UALookup(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("GET", "ua-lookup", params, nil, 10)
}

// URLInfo - Parse, analyze and retrieve content from the supplied URL
//
// The parameters this API accepts are:
// * url - The URL to probe
// * fetch-content - If this URL responds with html
// * ignore-certificate-errors - Ignore any TLS/SSL certificate errors and load the URL anyway
// * timeout - Timeout in seconds
// * retry - If the request fails for any reason try again this many times
//
// link: https://www.neutrinoapi.com/api/url-info
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) URLInfo(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("GET", "url-info", params, nil, 30)
}

// VerifySecurityCode - Check if a security code sent via SMS Verify or Phone Verify is valid
//
// The parameters this API accepts are:
// * security-code - The security code to verify
// * limit-by - If set then enable additional brute-force protection by limiting the number of attempts by the supplied value
//
// link: https://www.neutrinoapi.com/api/verify-security-code
// param: params net/url.Values type, a collection of key/value pairs
// returns *APIResponse
func (neutrinoAPIClient Client) VerifySecurityCode(params url.Values) *APIResponse {
	return neutrinoAPIClient.ExecRequest("GET", "verify-security-code", params, nil, 30)
}

// ExecRequest Make a request to the Neutrino API
func (neutrinoAPIClient Client) ExecRequest(httpMethod string, endpoint string, params url.Values, file *os.File, timeoutInSeconds time.Duration) *APIResponse {

	var client http.Client
	var err error
	var req *http.Request
	var resp *http.Response
	var apiUrl *url.URL
	var apiResponse = APIResponseOfErrorCode(NoStatus, NoContentType, NoData)

	apiUrl, err = apiUrl.Parse(fmt.Sprintf("%s%s", neutrinoAPIClient.BaseURL, endpoint))
	if err != nil {
		errStr := fmt.Sprintln(err)
		return APIResponseOfErrorCause(UrlParsingError, errStr)
	}

	if strings.EqualFold(httpMethod, "GET") {
		apiUrl.RawQuery = params.Encode()
		req, err = http.NewRequest("GET", apiUrl.String(), nil)
	} else {
		req, err = http.NewRequest("POST", apiUrl.String(), strings.NewReader(params.Encode()))
	}
	if err != nil {
		errStr := fmt.Sprintln(err)
		return APIResponseOfErrorCause(NetworkIoError, errStr)
	}
	req.Header.Add("Content-type", "application/x-www-form-urlencoded")
	req.Header.Add("User-ID", neutrinoAPIClient.UserID)
	req.Header.Add("API-Key", neutrinoAPIClient.APIKey)

	client = http.Client{
		Timeout: timeoutInSeconds * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: ConnectTimeoutInSeconds * time.Second,
			}).DialContext,
			TLSHandshakeTimeout: ConnectTimeoutInSeconds * time.Second,
		},
	}
	resp, err = client.Do(req)
	contentType := NoContentType
	statusCode := NoStatus
	if resp != nil {
		contentType = resp.Header.Get("content-type")
		statusCode = resp.StatusCode
	}

	switch t := err.(type) {
	case *net.OpError:
		errStr := fmt.Sprintln(err)
		return APIResponseOfErrorCause(NetworkIoError, errStr)
	case syscall.Errno:
		if t == syscall.ECONNREFUSED {
			return APIResponseOfErrorCode(statusCode, contentType, ConnectTimeout)
		}
		errStr := fmt.Sprintln(err)
		return APIResponseOfErrorCause(NetworkIoError, errStr)
	case net.Error:
		if strings.Contains(t.Error(), "Client.Timeout exceeded") {
			return APIResponseOfErrorCause(ReadTimeout, t.Error())
		}
		if strings.Contains(t.Error(), "connection refused") {
			return APIResponseOfErrorCause(ConnectTimeout, t.Error())
		}
		if strings.Contains(t.Error(), "unsupported protocol scheme") {
			return APIResponseOfErrorCause(UrlParsingError, t.Error())
		}
		if strings.Contains(t.Error(), "http: server gave HTTP response to HTTPS client") {
			return APIResponseOfErrorCause(NetworkIoError, err.Error())
			// return APIResponseOfErrorCause(TlsProtocolError, t.Error()) // could be this
		}
		return APIResponseOfErrorCause(NetworkIoError, err.Error())
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if statusCode == 200 {
		if strings.Contains(contentType, "application/json") {
			var result map[string]interface{}
			decoder := json.NewDecoder(resp.Body)
			if err := decoder.Decode(&result); err != nil {
				var errStr = fmt.Sprintln(err)
				return APIResponseOfErrorCause(InvalidJsonResponse, errStr)
			}
			return APIResponseOfData(statusCode, contentType, result)
		} else if file != nil {
			_, err := io.Copy(file, resp.Body)
			if err != nil {
				var errStr = fmt.Sprintln(err)
				return APIResponseOfErrorCause(NoData, errStr)
			}
			apiResponse = APIResponseOfOutputFilePath(statusCode, contentType, file.Name())
		} else {
			if body, err := ioutil.ReadAll(resp.Body); err == nil {
				apiResponse = APIResponseOfHTTPStatus(statusCode, contentType, NetworkIoError, string(body))
			} else {
				apiResponse = APIResponseOfErrorCode(statusCode, contentType, NetworkIoError)
			}
		}
	} else {
		if strings.Contains(contentType, "application/json") {
			var result map[string]interface{}
			decoder := json.NewDecoder(resp.Body)
			if err := decoder.Decode(&result); err != nil {
				var errStr = fmt.Sprintln(err)
				return APIResponseOfErrorCause(InvalidJsonResponse, errStr)
			}
			apiError := 0
			switch result["api-error"].(type) {
			case int:
				apiError = result["api-error"].(int)
			case float64:
				apiError = int(result["api-error"].(float64))
			}
			apiErrMsg, ok := result["api-error-msg"].(string)
			if ok {
				apiResponse = APIResponseOfHTTPStatus(statusCode, contentType, apiError, apiErrMsg)
			} else {
				apiResponse = APIResponseOfErrorCode(statusCode, contentType, InvalidJsonResponse)
			}
			return apiResponse
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				var errStr = fmt.Sprintln(err)
				return APIResponseOfErrorCause(NetworkIoError, errStr)
			}
			apiResponse = APIResponseOfHTTPStatus(statusCode, contentType, ApiGatewayError, string(body))
		}
	}
	return apiResponse
}
