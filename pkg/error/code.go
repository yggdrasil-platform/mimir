package _error

const (
	Unknown = 0

	MethodNotAllowed = 100
	RequestBodyEmpty = 101
	RequestBodyInvalid = 102
	RequestQueryInvalid = 103
)

var errDict = map[int]string{
	Unknown: "an unknown error has occurred",

	MethodNotAllowed: "this method is not allowed for this endpoint",
	RequestBodyEmpty: "a request body is required",
	RequestBodyInvalid: "the request body is invalid",
	RequestQueryInvalid: "the URL query is invalid",
}

// GetErrMessage returns the text for the error code. It returns the empty string if the code is unknown.
func GetErrMessage(code int) string {
	return errDict[code]
}
