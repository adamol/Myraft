package appendentries

import (
	"net/http"
	"myraft/state"
)

type AppendEntryRequest struct {
	Term         int
	Leader_id    int
	PrevLogIndex int
	PrevLogTerm  int
	Entries      []LogEntry
	LeaderCommit int
}

type AppendEntryReply struct {
	Term        int
	Success     bool
	CommitIndex int
}

type LogEntry struct {
	Command interface{}
	Term    int
}

type Controller struct {
	NodeInfo state.NodeInfo
}

func (c Controller) Handler(w http.ResponseWriter, r *http.Request) {

}
