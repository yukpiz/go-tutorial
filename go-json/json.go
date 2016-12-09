package main

import (
	"encoding/json"
	"fmt"
)

type compilationsAPI struct {
	SourceCode string
	Langid     int
}

func main() {
	var s compilationsAPI
	str := `{"sourceCode": "puts \"hello world\"","langid":1}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}