package state

import (
	"fmt"
	"strings"
	"encoding/json"
	"myraft/shared"
)

type NodeInfo struct {
	Servers []*Server
	SelfIndex int
	LeaderIndex int
}

var nodeInfo NodeInfo

func GetNodeInfo() *NodeInfo {
	return &nodeInfo
}

func (n NodeInfo) GetSelf() *Server {
	logger := shared.GetLogger()
	jsonState, _ := json.Marshal(n.Servers)
	logger.Log(fmt.Sprintf("trying to find self at index %v from servers %v", n.SelfIndex, string(jsonState)))

	return n.Servers[n.SelfIndex]
}

func (n NodeInfo) GetLeader() *Server {
	logger := shared.GetLogger()
	jsonState, _ := json.Marshal(n.Servers)
	logger.Log(fmt.Sprintf("trying to find leader at index %v from servers %v", n.LeaderIndex, string(jsonState)))

	return n.Servers[n.LeaderIndex]
}

func (n *NodeInfo) SetServers(servers []*Server) {
	shared.PrintStruct(n, "before server update: %v")

	n.Servers = servers

	shared.PrintStruct(n, "updated servers, new nodeinfo state: %v")
}

func (n *NodeInfo) InitState(logger *shared.Logger, selfInfo *string, leaderInfo *string, serverName *string) {
	self := Server{
		Name: *serverName,
		Host: fmt.Sprintf("http://%v", strings.Split(*selfInfo, ":")[0]),
		Port: strings.Split(*selfInfo, ":")[1],
	}

	n.Servers = append(n.Servers, &self)
	n.SelfIndex = 0

	logger.Log(fmt.Sprintf("Self appended to servers, length of servers: %v", len(n.Servers)))

	if *leaderInfo == "" {
		logger.Log(fmt.Sprintf("leaderinfo not provided, initializing as leader"))
		self.State = LEADER
		n.LeaderIndex = 0
	} else {
		logger.Log(fmt.Sprintf("leaderinfo was provided, initializing as follower"))
		self.State = FOLLOWER

		leader := Server{
			State: LEADER,
			Host:  fmt.Sprintf("http://%v", strings.Split(*leaderInfo, ":")[0]),
			Port:  strings.Split(*leaderInfo, ":")[1],
		}

		n.Servers = append(n.Servers, &leader)
		n.LeaderIndex = 1
	}

	jsonState, _ := json.Marshal(n)
	logger.Log(fmt.Sprintf("initializing state: %v ", string(jsonState)))
}