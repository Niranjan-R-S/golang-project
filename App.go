package main

import (
  "fmt"
  "bufio"
  "os"
  "net/http"
  "./controller/rnaTranscriptionController"
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

func main(){
  http.HandleFunc("/", rnaTranscriptionController.SayHelloWorld)
  http.HandleFunc("/rnaTranscription", rnaTranscriptionController.RnaTranscriptionProblem)
  http.ListenAndServe(":8080", nil)
}
