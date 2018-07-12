package join

import (
	"net/http"
	"encoding/json"
	"myraft/shared"
	"fmt"
	"bytes"
)

type JoinRequest struct {
	Server *shared.Server
}

type JoinResponse struct {
	Success bool
	NodeInfo shared.NodeInfo
}

type Controller struct {
	NodeInfo *shared.NodeInfo
}

func (c Controller) Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received join request rpc")

	var req JoinRequest

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Errorf(err.Error())
	}

	c.NodeInfo.Servers = append(c.NodeInfo.Servers, req.Server)

	jsonState, _ := json.Marshal(c.NodeInfo)
	fmt.Println("updated nodeinfo:", string(jsonState))

	lastIndex := len(c.NodeInfo.Servers) - 1

	for index, server := range c.NodeInfo.Servers {
		if index == c.NodeInfo.SelfIndex || index == lastIndex {
			continue
		}

		// go func
		go UpdateClusterState(index, server, c.NodeInfo)
	}

	res, _ := json.Marshal(JoinResponse{Success: true, NodeInfo: shared.NodeInfo{
		Servers: c.NodeInfo.Servers,
		LeaderIndex: c.NodeInfo.SelfIndex,
		SelfIndex: len(c.NodeInfo.Servers) - 1,
	}})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(res)
}

func SendJoinRpc(self *shared.Server, leader *shared.Server) JoinResponse {
	reqJson, _ := json.Marshal(JoinRequest{Server: self})

	url := fmt.Sprintf("%v/join", leader.GetUrl())
	fmt.Println("sending request to url:", url)

	fmt.Println(string(reqJson))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqJson))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	r, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	fmt.Println(r)
	defer r.Body.Close()
	var res JoinResponse

	if err = json.NewDecoder(r.Body).Decode(&res); err != nil {
		fmt.Println("error decoding join rpc")
		fmt.Errorf(err.Error())
	}

	return res
}









func (c Controller) UpdateStateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received update cluster state request request rpc")

	var req UpdateClusterStateRequest

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Errorf(err.Error())
	}

	c.NodeInfo = &req.NodeInfo

	jsonState, _ := json.Marshal(c.NodeInfo)
	fmt.Println("updated nodeinfo:", string(jsonState))

	res, _ := json.Marshal(UpdateClusterStateResponse{Success: true})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(res)
}

type UpdateClusterStateResponse struct {
	Success bool
}

type UpdateClusterStateRequest struct {
	NodeInfo shared.NodeInfo
}

func UpdateClusterState(index int, server *shared.Server, nodeInfo *shared.NodeInfo) {
	reqJson, _ := json.Marshal(UpdateClusterStateRequest{NodeInfo: shared.NodeInfo{
		Servers: nodeInfo.Servers,
		LeaderIndex: nodeInfo.SelfIndex,
		SelfIndex: index,
	}})

	url := fmt.Sprintf("%v/clusterstate", server.GetUrl())
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(reqJson))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	r, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()
	var res UpdateClusterStateResponse

	if err = json.NewDecoder(r.Body).Decode(&res); err != nil {
		fmt.Println("error decoding join rpc")
		fmt.Errorf(err.Error())
	}

	if res.Success != true {
		panic("failed to update follower cluster state")
	}
}

