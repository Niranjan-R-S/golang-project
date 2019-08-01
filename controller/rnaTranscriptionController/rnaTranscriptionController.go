package rnaTranscriptionController

import (
  "fmt"
  "net/http"
  "encoding/json"
  "../../model/rnaTranscriptionModel"
)

func SayHelloWorld(response http.ResponseWriter, request *http.Request){
    fmt.Fprintf(response, "Hello World")
}

func RnaTranscriptionProblem(response http.ResponseWriter, request *http.Request){
    decoder := json.NewDecoder(request.Body)

    var data myData
    err := decoder.Decode(&data)
    if err != nil {
      panic(err)
    }

    dnaString := data.DNA

    var userInputDna string = dnaString
    fmt.Fprintf(response, rnaTranscriptionModel.ConvertDNAToRNA(userInputDna))
}

type myData struct {
	DNA string
}
