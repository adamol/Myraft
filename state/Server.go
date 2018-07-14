package state

import "fmt"

const (
	FOLLOWER     = "follower"
	CANDIDATE    = "candidate"
	LEADER       = "leader"
)

type Server struct {
	Name string
	Port string
	Host string
	State string // FOLLOWER, CANDIDATE, LEADER
}

func (s Server) GetUrl() string {
	return fmt.Sprintf("%v:%v", s.Host, s.Port)
}
