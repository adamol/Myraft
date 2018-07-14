package main

import (
	"flag"
	"myraft/join"
	"myraft/clusterstate"
	"myraft/state"
	"myraft/shared"
)

var nodeInfo *state.NodeInfo
var logger *shared.Logger

/**
* go run main.go --at 127.0.0.1:8080
* go run main.go --at 127.0.0.1:8081 --join 127.0.0.1:8080
*/
func init() {
	selfInfo := flag.String("at", "127.0.0.1:8080", "default port")
	serverName := flag.String("as", "foobar", "default serverName")
	leaderInfo := flag.String("join", "", "default host")

	flag.Parse()

	nodeInfo = state.GetNodeInfo()
	logger = shared.GetLogger()

	logger.SetNodeName(*serverName)

	nodeInfo.InitState(logger, selfInfo, leaderInfo, serverName)
}

func main() {
	if nodeInfo.GetSelf().State != state.LEADER {
		res := join.HttpClient{Logger: *logger}.SendJoinRpc(nodeInfo.GetSelf(), nodeInfo.GetLeader())

		if res.Success {
			nodeInfo = &res.NodeInfo
		}
	}

	routes := make([]shared.Route, 0)

	for _, route := range join.GetRoutes() {
		routes = append(routes, route)
	}

	for _, route := range clusterstate.GetRoutes() {
		routes = append(routes, route)
	}

	var router shared.Router

	router.SetRoutes(routes)
	router.ListenOnPort(nodeInfo.GetSelf().Port)
}