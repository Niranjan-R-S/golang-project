package main

import (
    "./controller/leaderController"
    "./controller/rnaTranscriptionController"
    "./repository/graphRepository"
    "bufio"
    "fmt"
    "net/http"
    "os"
)

func getUserInput() string{
    fmt.Println("Enter the DNA string")
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    return text
}

func printFinalValue(rnaString string){
    fmt.Println(rnaString)
}

func init(){
    graphRepository.ConnectToDatabase()
}

func main(){
    http.HandleFunc("/", rnaTranscriptionController.SayHelloWorld)
    http.HandleFunc("/rnaTranscription", rnaTranscriptionController.RnaTranscriptionProblem)
    http.HandleFunc("/getAllLeaders", leaderController.GetAllLeaders)
    http.HandleFunc("/createLeader", leaderController.CreateLeader)
    _ = http.ListenAndServe(":8080", nil)
}
