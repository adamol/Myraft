package clusterstate

import (
	"myraft/shared"
	"net/http"
	"myraft/state"
)

type Controller struct {
	shared.BaseController
	Logger shared.Logger
	NodeInfo *state.NodeInfo
}

func (c *Controller) Handler(w http.ResponseWriter, r *http.Request) {
	var req UpdateClusterStateRequest

	c.ParseBody(r.Body, &req)

	shared.PrintStruct(req.NodeInfo, "nodeinfo from request: %v")

	c.NodeInfo = &req.NodeInfo

	shared.PrintStruct(c.NodeInfo, "updated nodeinfo: %v")

	c.RespondWithJson(w, UpdateClusterStateResponse{Success: true}, 200)
}
