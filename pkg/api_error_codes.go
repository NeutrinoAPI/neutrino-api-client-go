package neutrinoapi

import "fmt"

const (
	InvalidParameter = 1
	MaxCallLimit = 2
	BadUrl = 3
	AbuseDetected = 4
	NotResponding = 5
	Concurrent = 6
	NotVerified = 7
	TelephonyLimit = 8
	InvalidJson = 9
	AccessDenied = 10
	MaxPhoneCalls = 11
	BadAudio = 12
	HlrLimitReached = 13
	TelephonyBlocked = 14
	TelephonyRateExceeded = 15
	FreeLimit = 16
	RenderingFailed = 17
	DeprecatedApi = 18
	CreditLimitReached = 19
	NotMultiEnabled = 21
	NoBatchMode = 22
	BatchLimitExceeded = 23
	BatchInvalid = 24
	UserDefinedDailyLimit = 31
	AccessForbidden = 43
	RequestTooLarge = 44
	NoEndpoint = 45
	InternalServerError = 51
	ServerOffline = 52
	ConnectTimeout = 61
	ReadTimeout = 62
	Timeout = 63
	DnsLookupFailed = 64
	TlsProtocolError = 65
	UrlParsingError = 66
	NetworkIoError = 67
	FileIoError = 68
	InvalidJsonResponse = 69
	NoData = 70
	ApiGatewayError = 71
)

// Get description of error code
func GetErrorMessage(code int) string {
	switch code {
	case InvalidParameter:
		return "MISSING OR INVALID PARAMETER"
	case MaxCallLimit:
		return "DAILY API LIMIT EXCEEDED"
	case BadUrl:
		return "INVALID URL"
	case AbuseDetected:
		return "ACCOUNT OR IP BANNED"
	case NotResponding:
		return "NOT RESPONDING. RETRY IN 5 SECONDS"
	case Concurrent:
		return "TOO MANY CONNECTIONS"
	case NotVerified:
		return "ACCOUNT NOT VERIFIED"
	case TelephonyLimit:
		return "TELEPHONY NOT ENABLED ON YOUR ACCOUNT. PLEASE CONTACT SUPPORT FOR HELP"
	case InvalidJson:
		return "INVALID JSON. JSON CONTENT TYPE SET BUT NON-PARSABLE JSON SUPPLIED"
	case AccessDenied:
		return "ACCESS DENIED. PLEASE CONTACT SUPPORT FOR ACCESS TO THIS API"
	case MaxPhoneCalls:
		return "MAXIMUM SIMULTANEOUS PHONE CALLS"
	case BadAudio:
		return "COULD NOT LOAD AUDIO FROM URL"
	case HlrLimitReached:
		return "HLR LIMIT REACHED. CARD DECLINED"
	case TelephonyBlocked:
		return "CALLS AND SMS TO THIS NUMBER ARE LIMITED"
	case TelephonyRateExceeded:
		return "CALL IN PROGRESS"
	case FreeLimit:
		return "FREE PLAN LIMIT EXCEEDED"
	case RenderingFailed:
		return "RENDERING FAILED. COULD NOT GENERATE OUTPUT FILE"
	case DeprecatedApi:
		return "THIS API IS DEPRECATED. PLEASE USE THE LATEST VERSION"
	case CreditLimitReached:
		return "MAXIMUM ACCOUNT CREDIT LIMIT REACHED. PAYMENT METHOD DECLINED"
	case NotMultiEnabled:
		return "BATCH PROCESSING NOT ENABLED FOR THIS ENDPOINT"
	case NoBatchMode:
		return "BATCH PROCESSING NOT AVAILABLE ON YOUR PLAN"
	case BatchLimitExceeded:
		return "BATCH PROCESSING REQUEST LIMIT EXCEEDED"
	case BatchInvalid:
		return "INVALID BATCH REQUEST. DOES NOT CONFORM TO SPEC"
	case UserDefinedDailyLimit:
		return "DAILY API LIMIT EXCEEDED. SET BY ACCOUNT HOLDER"
	case AccessForbidden:
		return "ACCESS DENIED. USER ID OR API KEY INVALID"
	case RequestTooLarge:
		return "REQUEST TOO LARGE. MAXIMUM SIZE IS 5MB FOR DATA AND 25MB FOR UPLOADS"
	case NoEndpoint:
		return "ENDPOINT DOES NOT EXIST"
	case InternalServerError:
		return "FATAL EXCEPTION. REQUEST COULD NOT BE COMPLETED"
	case ServerOffline:
		return "SERVER OFFLINE. MAINTENANCE IN PROGRESS"
	case ConnectTimeout:
		return "TIMEOUT OCCURRED CONNECTING TO SERVER"
	case ReadTimeout:
		return "TIMEOUT OCCURRED READING API RESPONSE"
	case Timeout:
		return "TIMEOUT OCCURRED DURING API REQUEST"
	case DnsLookupFailed:
		return "ERROR RECEIVED FROM YOUR DNS RESOLVER"
	case TlsProtocolError:
		return "ERROR DURING TLS PROTOCOL HANDSHAKE"
	case UrlParsingError:
		return "ERROR PARSING REQUEST URL"
	case NetworkIoError:
		return "IO ERROR DURING API REQUEST"
	case FileIoError:
		return "IO ERROR WRITING TO OUTPUT FILE"
	case InvalidJsonResponse:
		return "INVALID JSON DATA RECEIVED"
	case NoData:
		return "NO PAYLOAD DATA RECEIVED"
	case ApiGatewayError:
		return "API GATEWAY ERROR"
	default:
		return fmt.Sprintf("API Error: %d", code)
	}
}
