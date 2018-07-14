package clusterstate

import (
	"myraft/shared"
	"myraft/state"
)

func GetRoutes() []shared.Route {
	routes := make([]shared.Route, 0)

	routes = append(routes, shared.Route{Path: "/clusterstate", Handler: GetController().Handler, Method: "PUT"})

	return routes
}

var controller *Controller = nil

func GetController() *Controller {
	if controller == nil {
		controller = &Controller{Logger: *shared.GetLogger(), NodeInfo: state.GetNodeInfo()}
	}
	return controller
}
