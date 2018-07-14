package shared

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type Route struct {
	Path string
	Handler func(http.ResponseWriter, *http.Request)
	Method string
}

type Router struct {
	routes []Route
}

func (r *Router) SetRoutes(routes []Route) {
	r.routes = routes
}

func (r Router) ListenOnPort(port string) {
	muxRouter := mux.NewRouter();

	logger := GetLogger()
	logger.Log(fmt.Sprintf("listening on port %v", port))

	for _, route := range r.routes {
		debuggedRoute := DebuggedRoute{RouteName: route.Path, RouteHandler: route.Handler}
		muxRouter.HandleFunc(route.Path, debuggedRoute.Handler).Methods(route.Method)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), muxRouter); err != nil {
		logger.Log(err.Error())
	}
}

type DebuggedRoute struct {
	RouteName string
	RouteHandler func(http.ResponseWriter, *http.Request)
}

func (dr DebuggedRoute) Handler(w http.ResponseWriter, r *http.Request) {
	logger.Log(fmt.Sprintf("Route %v received a request", dr.RouteName))

	dr.RouteHandler(w, r)
}