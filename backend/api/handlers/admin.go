package handlers

import (
	"friday/api/response"
	"friday/api/routing"
	"friday/api/util"
	"friday/database/models"
	"friday/database/repository"
	"friday/services/auth"
	"golang.org/x/crypto/bcrypt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type AdminAuthRequest struct {
	Password string `json:"password"`
}

type AdminAuthResponse struct {
	Token string `json:"token"`
}

var AuthenticateRoute = routing.Route{
	Method: "POST",
	Path:   "/authenticate", // Don't use admin prefix or AuthN middleware will block this route
	Handle: func(w http.ResponseWriter, r *http.Request) {

		// Create an AdminAuthRequest to unmarshal the request body into
		var adminAuthReq AdminAuthRequest
		if result, err := util.ReadRequestBody(w, r, &adminAuthReq); err != nil || !result {
			return
		}
		if adminAuthReq.Password == "" {
			response.WriteJSONErrorResponse(w, "Request is missing password", response.ErrorCodeMissingAdminPassword)
			return
		}

		passwordHash, exists := auth.GetCurrentHash()
		if !exists {
			response.WriteJSONErrorResponse(w, "Admin password hash not set", response.ErrorCodeAdminPasswordNotSet)
			return
		}

		now := time.Now()
		salt := now.Format("2006-01")
		if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(adminAuthReq.Password+salt)); err != nil {
			response.WriteJSONErrorResponse(w, "Invalid password", response.ErrorCodeInvalidPassword)
			return
		}

		token, err := auth.CreateAdminJWT()
		if err != nil {
			response.WriteJSONErrorResponse(w, "Failed to create JWT", response.ErrorCodeJWTGenerationFailed)
			return
		}

		response.WriteJSONSuccessResponse(w, AdminAuthResponse{Token: token})
	},
}

var UpsertShootRoute = routing.Route{
	Method: "POST",
	Path:   "/admin/shoot",
	Handle: func(w http.ResponseWriter, r *http.Request) {
		var shoot models.Shoot
		if result, err := util.ReadRequestBody(w, r, &shoot); err != nil || !result {
			return
		}

		// Create a database connection
		db := repository.Get()

		// We need an ID for updates, or a Name for creates
		if shoot.ID != 0 {
			// TODO: Update existing shoot
		} else if shoot.Name != "" {
			// Create new shoot
			err := db.CreateShoot(&shoot)
			if err != nil {
				response.WriteJSONErrorResponse(w, "Failed to create shoot", response.ErrorCodeFailedToCreateShoot)
				return
			}
		} else { // We need an ID for updates or a Name for creates - at least one is required
			response.WriteJSONErrorResponse(w, "Shoots requests require ID or Name", response.ErrorCodeShootsMissingRequiredFields)
			return
		}

		// This returns the updated shoot, or the created shoot with the new ID
		response.WriteJSONSuccessResponse(w, &shoot)
	},
}

var GetShootsRoute = routing.Route{
	Method: "GET",
	Path:   "/admin/shoots",
	Handle: func(w http.ResponseWriter, r *http.Request) {
		// Create a database connection
		db := repository.Get()

		shoots, err := db.GetShoots()
		if err != nil {
			response.WriteJSONErrorResponse(w, "Failed to get shoots", response.ErrorCodeFailedToGetShoots)
			return
		}

		response.WriteJSONSuccessResponse(w, &shoots)
	},
}

type DeleteShootRequest struct {
	ID int `json:"id"`
}

var DeleteShootRoute = routing.Route{
	Method: "DELETE",
	Path:   "/admin/unshoot",
	Handle: func(w http.ResponseWriter, r *http.Request) {
		var deleteShootRequest DeleteShootRequest
		if result, err := util.ReadRequestBody(w, r, &deleteShootRequest); err != nil || !result {
			return
		}
		if deleteShootRequest.ID == 0 {
			response.WriteJSONErrorResponse(w, "Request is missing shoot ID", response.ErrorCodeShootsMissingRequiredFields)
			return
		}

		// Create a database connection
		db := repository.Get()

		err := db.DeleteShoot(deleteShootRequest.ID)
		if err != nil {
			response.WriteJSONErrorResponse(w, "Failed to delete shoot", response.ErrorCodeFailedToDeleteShoot)
			return
		}

		response.WriteJSONSuccessResponse(w, nil)
	},
}

type AdminTestResponse struct {
	Test string `json:"test"`
}

var TestRoute = routing.Route{
	Method: "GET",
	Path:   "/admin/test",
	Handle: func(w http.ResponseWriter, r *http.Request) {
		response.WriteJSONSuccessResponse(w, &AdminTestResponse{Test: "asdf"})
	},
}

var ExportDatabaseRoute = routing.Route{
	Method: "GET",
	Path:   "/admin/export-database",
	Handle: func(w http.ResponseWriter, r *http.Request) {
		dbPath := repository.GetDBPath()

		file, err := os.Open(dbPath)
		if err != nil {
			response.WriteJSONErrorResponse(w, "Failed to read database file", response.ErrorCodeFailedToReadDatabaseFile)
			return
		}
		defer func(file *os.File) {
			_ = file.Close()
		}(file)

		w.Header().Set("Content-Disposition", "attachment; filename=friday.db")
		w.Header().Set("Content-Type", "application/x-sqlite3")

		_, err = io.Copy(w, file)
		if err != nil {
			response.WriteJSONErrorResponse(w, "Failed to write file to response", response.ErrorCodeFiledToWriteDatabaseFileResponse)
			return
		}
	},
}

var ImportDatabaseRoute = routing.Route{
	Method: "POST",
	Path:   "/admin/import-database",
	Handle: func(w http.ResponseWriter, r *http.Request) {
		// Limit upload size (200 MB)
		r.Body = http.MaxBytesReader(w, r.Body, 200<<20)
		if err := r.ParseMultipartForm(200 << 20); err != nil {
			response.WriteJSONErrorResponse(w, "Failed to parse form data: "+err.Error(), response.ErrorCodeFailedToReadDatabaseFileRequest)
			return
		}

		file, _, err := r.FormFile("file")
		if err != nil {
			response.WriteJSONErrorResponse(w, "Missing file in request: "+err.Error(), response.ErrorCodeMissingDatabaseFileRequest)
			return
		}
		defer func(file multipart.File) {
			_ = file.Close()
		}(file)

		dbPath := repository.GetDBPath()
		dir := filepath.Dir(dbPath)

		tmp, err := os.CreateTemp(dir, "friday.db.import.*")
		if err != nil {
			response.WriteJSONErrorResponse(w, "Failed to create temp file: "+err.Error(), response.ErrorCodeFailedToWriteDatabaseFile)
			return
		}

		// Ensure cleanup on failure
		tmpPath := tmp.Name()
		defer func() {
			_ = tmp.Close()
			if _, err := os.Stat(tmpPath); err == nil {
				_ = os.Remove(tmpPath)
			}
		}()

		if _, err := io.Copy(tmp, file); err != nil {
			response.WriteJSONErrorResponse(w, "Failed to save uploaded file: "+err.Error(), response.ErrorCodeFailedToWriteDatabaseFile)
			return
		}

		// Close before rename
		if err := tmp.Close(); err != nil {
			response.WriteJSONErrorResponse(w, "Failed to close temp file: "+err.Error(), response.ErrorCodeFailedToWriteDatabaseFile)
			return
		}

		// Replace existing DB atomically
		if err := os.Rename(tmpPath, dbPath); err != nil {
			response.WriteJSONErrorResponse(w, "Failed to replace database file: "+err.Error(), response.ErrorCodeFailedToWriteDatabaseFile)
			return
		}

		// Reload the database connection
		db := repository.Get()
		if err := db.ReloadFromDisk(); err != nil {
			response.WriteJSONErrorResponse(w, "Failed to reload database from disk: "+err.Error(), response.ErrorCodeDatabaseConnection)
		}

		response.WriteJSONSuccessResponse(w, nil)
	},
}

var AdminRoutes = []routing.Route{
	AuthenticateRoute,
	UpsertShootRoute,
	GetShootsRoute,
	DeleteShootRoute,
	TestRoute,
	ExportDatabaseRoute,
	ImportDatabaseRoute,
}

func init() {
	routing.RegisterRoutes(AdminRoutes)
}
