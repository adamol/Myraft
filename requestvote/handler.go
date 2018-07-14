package requestvote

import (
	"net/http"
	"myraft/state"
)

type RequestVoteRequest struct {
	Term         int // candidate's term
	CandidateId  int // candidate requesting vote
	LastLogIndex int // index of candidate's last log entry
	LastLogTerm  int // term of candidate's last log entry
}

type RequestVoteReply struct {
	Term        int  // currentTerm, for candidate to update itself
	VoteGranted bool // true means candidate received vote
}

type Controller struct {
	NodeInfo state.NodeInfo
}

func (c Controller) Handler(w http.ResponseWriter, r *http.Request) {

}
