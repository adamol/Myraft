package shared

import (
	"fmt"
	"strings"
	"encoding/json"
)

const (
	FOLLOWER     = "follower"
	CANDIDATE    = "candidate"
	LEADER       = "leader"
)

type Server struct {
	Port string
	Host string
	State string // FOLLOWER, CANDIDATE, LEADER
}

func (s Server) GetUrl() string {
	return fmt.Sprintf("%v:%v", s.Host, s.Port)
}

type NodeInfo struct {
	Servers []*Server
	SelfIndex int
	LeaderIndex int
}

func (n NodeInfo) GetSelf() *Server {
	fmt.Println(fmt.Sprintf("trying to access index %v from length %v", n.SelfIndex, len(n.Servers)))

	jsonState, _ := json.Marshal(n)
	fmt.Println(string(jsonState))

	return n.Servers[n.SelfIndex]
}

func (n NodeInfo) GetLeader() *Server {
	return n.Servers[n.LeaderIndex]
}

func (n *NodeInfo) InitState(selfInfo *string, leaderInfo *string) {
	self := Server{
		Host: fmt.Sprintf("http://%v", strings.Split(*selfInfo, ":")[0]),
		Port: strings.Split(*selfInfo, ":")[1],
	}

	n.Servers = append(n.Servers, &self)
	n.SelfIndex = 0

	fmt.Println(fmt.Sprintf("Self appended to servers, length of servers: %v", len(n.Servers)))

	if *leaderInfo == "" {
		fmt.Println(fmt.Sprintf("leaderinfo not provided, initializing as leader"))
		self.State = LEADER
		n.LeaderIndex = 0
	} else {
		fmt.Println(fmt.Sprintf("leaderinfo was provided, initializing as follower"))
		self.State = FOLLOWER

		leader := Server{
			State: LEADER,
			Host:  fmt.Sprintf("http://%v", strings.Split(*leaderInfo, ":")[0]),
			Port:  strings.Split(*leaderInfo, ":")[1],
		}

		n.Servers = append(n.Servers, &leader)
		n.LeaderIndex = 1
	}
}