package response

type ErrorCode int

const (
	// ErrorCodeNone 0 = no error
	ErrorCodeNone ErrorCode = 0
	// ErrorCodeTest 1 = for unit tests
	ErrorCodeTest ErrorCode = 2

	// Example Error Code Category range:
	// ErrorCodeExample ErrorCode = 1000 + iota
)
