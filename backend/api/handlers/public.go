package handlers

import "friday/api/routing"

var ListFavoritePhotosRoute = routing.Route{
	Method: "GET",
	Path:   "/photos/favorites",
	Handle: GetListFavoritePhotosRouteHandler(),
}
