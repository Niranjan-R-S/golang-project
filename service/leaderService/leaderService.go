package leaderService

import (
	"../../repository/graphRepository"
	"fmt"
)

func CreateLeader(leaderName string) string{
	_, _ = graphRepository.DB.Execute(
		fmt.Sprintf("graph.addVertex(%s)", leaderName),
		map[string]string{},
		map[string]string{},
	)
	return "Leader created successfully"
}

func GetAllLeaders(){
	result, _ := graphRepository.DB.Execute(
		"g.V()",
		map[string]string{},
		map[string]string{},
	)
	fmt.Println(result)
}