package main

import (
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
    _ = http.ListenAndServe(":8080", nil)
}
