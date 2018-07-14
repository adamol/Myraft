package join

import (
	"myraft/shared"
	"fmt"
	"encoding/json"
	"myraft/state"
)

type HttpClient struct {
	shared.HttpClient
	Logger shared.Logger
}

func (c HttpClient) SendJoinRpc(self *state.Server, leader *state.Server) JoinResponse {
	url := fmt.Sprintf("%v/join", leader.GetUrl())

	reqJson, _ := json.Marshal(JoinRequest{Server: self})

	r := c.SendRequestRpc("POST", url, reqJson)

	var res JoinResponse

	c.ParseBody(r.Body, res)

	return res
}
