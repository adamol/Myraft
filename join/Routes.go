package join

import (
	"myraft/shared"
	"myraft/state"
)

func GetRoutes() []shared.Route {
	routes := make([]shared.Route, 0)

	routes = append(routes, shared.Route{Path: "/join", Handler: GetController().Handler, Method: "Post"})

	return routes
}

var controller *Controller = nil

func GetController() Controller {
	if controller == nil {
		controller = &Controller{NodeInfo: state.GetNodeInfo()}
	}
	return *controller
}