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
	// ErrorCodeFailedToReadDatabaseFile Error code for API response on failed database file export
	ErrorCodeFailedToReadDatabaseFile ErrorCode = 13
	// ErrorCodeFiledToWriteDatabaseFileResponse Error code for API response on failed database file export write
	ErrorCodeFiledToWriteDatabaseFileResponse ErrorCode = 14
	// ErrorCodeFailedToReadDatabaseFileRequest Error code for API response on failed database file import read
	ErrorCodeFailedToReadDatabaseFileRequest ErrorCode = 15
	// ErrorCodeMissingDatabaseFileRequest Error code for API response on missing database file in import request
	ErrorCodeMissingDatabaseFileRequest ErrorCode = 16
	// ErrorCodeFailedToWriteDatabaseFile Error for any failures in writing database files
	ErrorCodeFailedToWriteDatabaseFile ErrorCode = 17

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
	// ErrorCodeFailedToGetShoots API return error for repository.GetShoots failures
	ErrorCodeFailedToGetShoots ErrorCode = 43
	// ErrorCodeFailedToDeleteShoot API return error for repository.DeleteShoot failures
	ErrorCodeFailedToDeleteShoot ErrorCode = 44
	// ErrorCodeFailedToUpdateShoot API return error for repository.UpdateShoot failures
	ErrorCodeFailedToUpdateShoot ErrorCode = 45

	// ErrorCodePhotosGeneral General photos error code - starts the Photos error block
	ErrorCodePhotosGeneral ErrorCode = 60
	// ErrorCodeFailedToReadUploadPhotoRequest Error code for failed reading form data, or photo is too large
	ErrorCodeFailedToReadUploadPhotoRequest ErrorCode = 61
	// ErrorCodeMissingPhotoFileRequest Error code for missing photo file in upload request
	ErrorCodeMissingPhotoFileRequest ErrorCode = 62
	// ErrorCodeFailedToWritePhoto Error code for failures with the photo temp file
	ErrorCodeFailedToWritePhoto ErrorCode = 63
	// ErrorCodeFailedToUploadToCloudflare Error code for failures uploading photo to Cloudflare
	ErrorCodeFailedToUploadToCloudflare ErrorCode = 64
	// ErrorCodeMissingExtraPhotoData Error code for missing shoot_id or categories in upload photo request
	ErrorCodeMissingExtraPhotoData ErrorCode = 65
	// ErrorCodeInvalidExtraPhotoData Error code for invalid shoot_id or categories in upload photo request
	ErrorCodeInvalidExtraPhotoData ErrorCode = 66
	// ErrorCodeFailedToCreatePhoto API return error for failure during photo database operations
	ErrorCodeFailedToCreatePhoto ErrorCode = 67
	// ErrorCodeFailedToGetPhotos API return error for failure during photo retrieval database operations
	ErrorCodeFailedToGetPhotos ErrorCode = 68

	// Example Error Code Category range:
	// ErrorCodeExample ErrorCode = 1000 + iota
)
