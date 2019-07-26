package main

import (
  "fmt"
  "bufio"
  "os"
  "net/http"
  "encoding/json"
  "github.com/qasaur/gremgo"
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

func getDNAComponents() []string{
    var dnaComponents = []string{"G", "C", "T", "A"}
    return dnaComponents
}

func getRNAComponents() []string{
    var rnaComponents = []string{"C", "G", "A", "U"}
    return rnaComponents
}

func getComponentMap(dnaComponents []string, rnaComponents []string) map[string] string{
  var componentMap map[string] string
  componentMap = make(map[string]string)
  for index, value := range dnaComponents{
    componentMap[value] = rnaComponents[index]
  }
  return componentMap
}

func convertDNAToRNA(userInputDna string) string{
    var componentMap map[string] string = getComponentMap(getDNAComponents(), getRNAComponents())
    var finalString = ""
    for _, value := range userInputDna{
        finalString += componentMap[string(value)]
    }
    return finalString
}

type myData struct {
	DNA string
}

func rnaTranscriptionProblem(response http.ResponseWriter, request *http.Request){
    decoder := json.NewDecoder(request.Body)

    var data myData
    err := decoder.Decode(&data)
    if err != nil {
      panic(err)
    }

    dnaString := data.DNA

    var userInputDna string = dnaString
    fmt.Fprintf(response, convertDNAToRNA(userInputDna))
}

func sayHelloWorld(response http.ResponseWriter, request *http.Request){
    fmt.Fprintf(response, "Hello World")
}

func main(){
  http.HandleFunc("/", sayHelloWorld)
  http.HandleFunc("/rnaTranscription", rnaTranscriptionProblem)
  http.ListenAndServe(":8080", nil)
}
