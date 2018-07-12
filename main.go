package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"flag"
	"fmt"
	"myraft/appendentries"
	"myraft/requestvote"
	"myraft/join"
	"myraft/shared"
	"encoding/json"
)

var nodeInfo shared.NodeInfo

/**
* go run main.go --at 127.0.0.1:8080
* go run main.go --at 127.0.0.1:8081 --join 127.0.0.1:8080
*/
func init() {
	selfInfo := flag.String("at", "127.0.0.1:8080", "default port")
	leaderInfo := flag.String("join", "", "default host")

	flag.Parse()

	fmt.Println(fmt.Sprintf("'at' flag parsed as %v", *selfInfo))
	fmt.Println(fmt.Sprintf("'join' flag parsed as %v", *leaderInfo))

	nodeInfo.InitState(selfInfo, leaderInfo)

	jsonState, _ := json.Marshal(nodeInfo)
	fmt.Println(string(jsonState))
}

func main() {
	if nodeInfo.GetSelf().State != shared.LEADER {
		res := join.SendJoinRpc(nodeInfo.GetSelf(), nodeInfo.GetLeader())

		if res.Success {
			nodeInfo = res.NodeInfo
		}
	}

	r := mux.NewRouter();

	r.HandleFunc("/join", join.Controller{NodeInfo: &nodeInfo}.Handler).Methods("POST")
	r.HandleFunc("/clusterstate", join.Controller{NodeInfo: &nodeInfo}.UpdateStateHandler).Methods("PUT")
	r.HandleFunc("/request-vote", requestvote.Controller{NodeInfo: nodeInfo}.Handler).Methods("POST")
	r.HandleFunc("/append-entry", appendentries.Controller{NodeInfo: nodeInfo}.Handler).Methods("POST")

	fmt.Println("routing configured...")
	listenPort := fmt.Sprintf(":%v", nodeInfo.GetSelf().Port)

	fmt.Println("listening on port", listenPort)

	if err := http.ListenAndServe(listenPort, r); err != nil {
		log.Fatal(err)
	}
}