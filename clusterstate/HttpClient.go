package clusterstate

import (
	"myraft/shared"
	"encoding/json"
	"fmt"
	"myraft/state"
)

type HttpClient struct {
	shared.HttpClient
}

func (c HttpClient) SendUpdateRpc(index int, server *state.Server, nodeInfo *state.NodeInfo) {
	reqJson, _ := json.Marshal(UpdateClusterStateRequest{NodeInfo: state.NodeInfo{
		Servers: nodeInfo.Servers,
		LeaderIndex: nodeInfo.SelfIndex,
		SelfIndex: index,
	}})

	url := fmt.Sprintf("%v/clusterstate", server.GetUrl())

	r := c.SendRequestRpc("PUT", url, reqJson)

	var res UpdateClusterStateResponse

	c.ParseBody(r.Body, &res)

	if res.Success != true {
		panic("failed to update follower cluster state")
	}
}



