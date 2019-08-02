package main

import (
  "fmt"
  "bufio"
  "os"
  "net/http"
  "log"
  "github.com/qasaur/gremgo"
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

func connectToDatabase(){
    errs := make(chan error)
    go func(chan error) {
      err := <-errs
      log.Fatal("Lost connection to the database: " + err.Error())
    }(errs) // Example of connection error handling logic

    dialer := gremgo.NewDialer("ws://127.0.0.1:8182") // Returns a WebSocket dialer to connect to Gremlin Server
    g, err := gremgo.Dial(dialer, errs) // Returns a gremgo client to interact with
    if err != nil {
      fmt.Println(err)
        return
    }
    res, err := g.Execute( // Sends a query to Gremlin Server with bindings
      "g.V().has('name', 'neptune')",
      map[string]string{},
      map[string]string{},
    )
    if err != nil {
      fmt.Println(err)
        return
    }
    fmt.Println(res)
}

func main(){
    connectToDatabase()
    http.HandleFunc("/", rnaTranscriptionController.SayHelloWorld)
    http.HandleFunc("/rnaTranscription", rnaTranscriptionController.RnaTranscriptionProblem)
    http.ListenAndServe(":8080", nil)
}
