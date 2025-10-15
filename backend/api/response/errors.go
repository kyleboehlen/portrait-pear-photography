package response

type ErrorCode int

const (
	// ErrorCodeNone 0 = no error
	ErrorCodeNone ErrorCode = 0
	// ErrorCodeTest 2 = for unit tests
	ErrorCodeTest ErrorCode = 2

	// ErrorCodeDatabaseGeneral General database error code - starts the Database error block
	ErrorCodeDatabaseGeneral ErrorCode = 10
	// ErrorCodeDatabaseConnection used for issues with the database setting up the repo
	ErrorCodeDatabaseConnection ErrorCode = 11
	// ErrorCodeDatabaseNotHealthy Error code for a failed ping to the database
	ErrorCodeDatabaseNotHealthy ErrorCode = 12

	// ErrorCodeAdminAuthorizationGeneral General admin authorization error code - starts the Admin Authorization error block
	ErrorCodeAdminAuthorizationGeneral ErrorCode = 20
	// ErrorCodeMissingAuthorizationHeader For admin routes where the Authorization header is missing
	ErrorCodeMissingAuthorizationHeader       ErrorCode = 21
	ErrorCodeInvalidAuthorizationHeaderFormat ErrorCode = 22
	ErrorCodeInvalidOrExpiredToken            ErrorCode = 23
	// ErrorCodeJWTGenerationFailed JWT service errored out when generating the admin token
	ErrorCodeJWTGenerationFailed ErrorCode = 24
	// ErrorCodeInvalidPassword An error if the provided password doesn't match the bcrypt hash
	ErrorCodeInvalidPassword ErrorCode = 25
	// ErrorCodeAdminPasswordNotSet An error if the ADMIN_PASSWORD_HASH env var isn't set
	ErrorCodeAdminPasswordNotSet ErrorCode = 26
	// ErrorCodeMissingAdminPassword The request body does not contain the password
	ErrorCodeMissingAdminPassword ErrorCode = 27

	// ErrorCodeRequestGeneral General request error code - starts the Request error block
	ErrorCodeRequestGeneral ErrorCode = 30
	// ErrorCodeRequestBodyUnreadable Error code for when the request body can't be read using io
	ErrorCodeRequestBodyUnreadable ErrorCode = 31
	// ErrorCodeRequestJSONUnmarshalFailed Failed to unmarshal a request body into a struct, usually due to invalid JSON
	ErrorCodeRequestJSONUnmarshalFailed ErrorCode = 32

	// ErrorCodeShootsGeneral General shoots error code - starts the Shoots error block
	ErrorCodeShootsGeneral               ErrorCode = 40
	ErrorCodeShootsMissingRequiredFields ErrorCode = 41
	// ErrorCodeFailedToCreateShoot API return error for repository.CreateShoot failures
	ErrorCodeFailedToCreateShoot ErrorCode = 42

	// Example Error Code Category range:
	// ErrorCodeExample ErrorCode = 1000 + iota
)
