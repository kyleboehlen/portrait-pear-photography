package handlers

import (
	"friday/api/response"
	"friday/database/repository"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func GetExportDatabaseRouteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}

func GetImportDatabaseRouteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Limit upload size (200 MB)
		r.Body = http.MaxBytesReader(w, r.Body, 200<<20)
		if err := r.ParseMultipartForm(200 << 20); err != nil {
			response.WriteJSONErrorResponse(w, "Failed to parse database form data: "+err.Error(), response.ErrorCodeFailedToReadDatabaseFileRequest)
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
	}
}
