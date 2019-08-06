package rnaTranscriptionService

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

func ConvertDNAToRNA(userInputDna string) string{
    var componentMap = getComponentMap(getDNAComponents(), getRNAComponents())
    var finalString = ""
    for _, value := range userInputDna{
        finalString += componentMap[string(value)]
    }
    return finalString
}
