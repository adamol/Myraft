package join

import (
	"myraft/state"
)

type JoinRequest struct {
	Server *state.Server
}

type JoinResponse struct {
	Success bool
	NodeInfo state.NodeInfo
}
