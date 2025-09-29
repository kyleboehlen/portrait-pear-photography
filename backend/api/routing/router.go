package routing

import (
	"fmt"
	"net/http"
)

type Route struct {
	Method string
	Path   string
	Handle http.HandlerFunc
}

var registeredRoutes []Route

func RegisterRoutes(routes []Route) {
	registeredRoutes = append(registeredRoutes, routes...)
}

// GetRegisteredRoutesCount Helper function for testing functionality
func GetRegisteredRoutesCount() int {
	return len(registeredRoutes)
}
func GetNewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	fmt.Printf("Registering %d routes:\n", len(registeredRoutes))
	for _, route := range registeredRoutes {
		fmt.Printf("  Route: %s %s\n", route.Method, route.Path)

		// Pattern matching for mux is finicky, checking the http method explicitly
		mux.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			// Denying the request if it doesn't match the handler's specified method
			if r.Method != route.Method {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
			// Calling the method passed in from the handler
			route.Handle(w, r)
		})
	}
	return mux
}
