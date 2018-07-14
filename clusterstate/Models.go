package clusterstate

import (
	"myraft/state"
)

type UpdateClusterStateResponse struct {
	Success bool
}

type UpdateClusterStateRequest struct {
	NodeInfo state.NodeInfo
}
