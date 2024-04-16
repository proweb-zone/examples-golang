package main

import (
	"fmt"
	"strings"
)

func main(){
sentence := "the fiRsT sentence"
result := capitalizeWords(sentence)
fmt.Println(result)
}

func capitalizeWords(sentence string) string {

words := strings.Fields(sentence)
capitalizedSentence := ""
for _, word := range words {
capitalizedWord := strings.Title(strings.ToLower(word))
capitalizedSentence += capitalizedWord + " "
}

return capitalizedSentence
}
