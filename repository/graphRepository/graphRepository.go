package graphRepository

import(
  "github.com/qasaur/gremgo"
  "log"
  "fmt"
)

var DB gremgo.Client
var err error

func ConnectToDatabase(){
    errs := make(chan error)
    go func(chan error) {
      err := <-errs
      log.Fatal("Lost connection to the database: " + err.Error())
    }(errs) // Connection error handling logic

    dialer := gremgo.NewDialer("ws://127.0.0.1:8182") // Returns a WebSocket dialer to connect to Gremlin Server
    DB, err = gremgo.Dial(dialer, errs) // Returns a gremgo client to interact with
    if err != nil {
      fmt.Println(err)
        return
    }
}

func queryFromDatabase(){
    res, err := DB.Execute( // Sends a query to Gremlin Server with bindings
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
