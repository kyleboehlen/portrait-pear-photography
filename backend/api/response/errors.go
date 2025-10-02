package response

type ErrorCode int

const (
	// ErrorCodeNone 0 = no error
	ErrorCodeNone ErrorCode = 0
	// ErrorCodeTest 2 = for unit tests
	ErrorCodeTest ErrorCode = 2

	// ErrorCodeDatabaseConnection DATABASE ERRORS - 1000
	// ErrorCodeDatabaseConnection Error code used for issues with the database setting up the repo
	ErrorCodeDatabaseConnection = 1000 + iota
	// ErrorCodeDatabaseNotHealthy Error code for a failed ping to the database
	ErrorCodeDatabaseNotHealthy = 1000 + iota

	// Example Error Code Category range:
	// ErrorCodeExample ErrorCode = 1000 + iota
)
