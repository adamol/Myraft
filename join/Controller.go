package join

import (
	"net/http"
	"myraft/shared"
	"myraft/clusterstate"
	"myraft/state"
)

type Controller struct {
	shared.BaseController
	NodeInfo *state.NodeInfo
}

func (c Controller) Handler(w http.ResponseWriter, r *http.Request) {
	var req JoinRequest
	c.ParseBody(r.Body, &req)

	servers := append(c.NodeInfo.Servers, req.Server)

	shared.PrintStruct(req, "append new server: %v")

	c.NodeInfo.SetServers(servers)

	lastIndex := len(c.NodeInfo.Servers) - 1

	for index, server := range c.NodeInfo.Servers {
		if index == c.NodeInfo.SelfIndex || index == lastIndex {
			continue
		}

		go clusterstate.HttpClient{}.SendUpdateRpc(index, server, c.NodeInfo)
	}

	c.RespondWithJson(w, JoinResponse{Success: true, NodeInfo: state.NodeInfo{
		Servers: c.NodeInfo.Servers,
		LeaderIndex: c.NodeInfo.SelfIndex,
		SelfIndex: len(c.NodeInfo.Servers) - 1,
	}}, 200)
}

