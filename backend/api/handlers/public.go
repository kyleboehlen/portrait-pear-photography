package handlers

import "friday/api/routing"

var ListFavoritePhotosRoute = routing.Route{
	Method: "GET",
	Path:   "/photos/favorites",
	Handle: GetListFavoritePhotosRouteHandler(),
}

var ListShootPhotosRoute = routing.Route{
	Method: "POST",
	Path:   "/photos/shoot",
	Handle: GetListShootPhotosRouteHandler(),
}

var PublicRoutes = []routing.Route{
	ListFavoritePhotosRoute,
	ListShootPhotosRoute,
}

func init() {
	routing.RegisterRoutes(PublicRoutes)
}
