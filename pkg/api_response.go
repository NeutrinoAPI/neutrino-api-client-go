package neutrinoapi

const (
	NoStatus      = 0
	NoContentType = ""
)

// API response payload, holds the response data along with any error details
type APIResponse struct {
    // The response data for JSON based APIs
	Data map[string]interface{}
    // The local file path storing the output for file based APIs
	File *string
    // The response content type (MIME type)
	ContentType string
    // The HTTP status code returned
	HttpStatusCode int
    // The API error code if any error has occurred
	ErrorCode *int
    // The API error message if any error has occurred
	ErrorMessage *string
    // For client-side errors or exceptions get the underlying cause
	ErrorCause *string
}

// Was this request successul
func (apiResponse APIResponse) IsOK() bool {
	return apiResponse.Data != nil || apiResponse.File != nil
}

// Create an API response for JSON data
func APIResponseOfData(statusCode int, contentType string, data map[string]interface{}) *APIResponse {
	return &APIResponse{data, nil, contentType, statusCode, nil, nil, nil}
}

// Create an API response for file data
func APIResponseOfOutputFilePath(statusCode int, contentType string, file string) *APIResponse {
	return &APIResponse{nil, &file, contentType, statusCode, nil, nil, nil}
}

// Create an API response for error code
func APIResponseOfErrorCode(statusCode int, contentType string, errorCode int) *APIResponse {
	var errorMsg = GetErrorMessage(errorCode)
	return &APIResponse{nil, nil, contentType, statusCode, &errorCode, &errorMsg, nil}
}

// Create an API response for status code
func APIResponseOfHTTPStatus(statusCode int, contentType string, errorCode int, errorMsg string) *APIResponse {
	return &APIResponse{nil, nil, contentType, statusCode, &errorCode, &errorMsg, nil}
}

// Create an API response for error cause
func APIResponseOfErrorCause(errorCode int, errorCause string) *APIResponse {
	var errorMsg = GetErrorMessage(errorCode)
	return &APIResponse{nil, nil, NoContentType, NoStatus, &errorCode, &errorMsg, &errorCause}
}
