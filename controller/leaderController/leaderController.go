package leaderController

import (
	"../../service/leaderService"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateLeader(response http.ResponseWriter, request *http.Request){
	decoder := json.NewDecoder(request.Body)

	var data leaderData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	leaderName := data.name

	_, _ = fmt.Fprintf(response, leaderService.CreateLeader(leaderName))
}

func GetAllLeaders(response http.ResponseWriter, request *http.Request){
	leaderService.GetAllLeaders()
	_, _ = fmt.Fprintf(response, "")
}

type leaderData struct {
	name string
}
