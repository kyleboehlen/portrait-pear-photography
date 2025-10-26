package handlers

import (
	"friday/api/routing"
)

var AuthenticateRoute = routing.Route{
	Method: "POST",
	Path:   "/authenticate", // Don't use admin prefix or AuthN middleware will block this route
	Handle: GetAuthenticateRouteHandler(),
}

var UpsertShootRoute = routing.Route{
	Method: "POST",
	Path:   "/admin/shoot",
	Handle: GetUpsertShootRouteHandler(),
}

var ListShootsRoute = routing.Route{
	Method: "GET",
	Path:   "/admin/shoots",
	Handle: GetListShootsRouteHandler(),
}

var DeleteShootRoute = routing.Route{
	Method: "DELETE",
	Path:   "/admin/unshoot",
	Handle: GetDeleteShootRouteHandler(),
}

var TestRoute = routing.Route{
	Method: "GET",
	Path:   "/admin/test",
	Handle: GetTestRouteHandler(),
}

var ExportDatabaseRoute = routing.Route{
	Method: "GET",
	Path:   "/admin/export-database",
	Handle: GetExportDatabaseRouteHandler(),
}

var ImportDatabaseRoute = routing.Route{
	Method: "POST",
	Path:   "/admin/import-database",
	Handle: GetImportDatabaseRouteHandler(),
}

var UploadPhotoRoute = routing.Route{
	Method: "POST",
	Path:   "/admin/upload-photo",
	Handle: GetUploadPhotoRouteHandler(),
}

var DeletePhotoRoute = routing.Route{
	Method: "DELETE",
	Path:   "/admin/delete-photo",
	Handle: GetDeletePhotoRouteHandler(),
}

var ListPhotosFilteredRoute = routing.Route{
	Method: "POST",
	Path:   "/admin/photos",
	Handle: GetListPhotosFilteredRouteHandler(),
}

var AdminRoutes = []routing.Route{
	AuthenticateRoute,
	UpsertShootRoute,
	ListShootsRoute,
	DeleteShootRoute,
	TestRoute,
	ExportDatabaseRoute,
	ImportDatabaseRoute,
	UploadPhotoRoute,
	DeletePhotoRoute,
	ListPhotosFilteredRoute,
}

func init() {
	routing.RegisterRoutes(AdminRoutes)
}
