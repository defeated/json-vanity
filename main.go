package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	output := JsonPrettyPrint(string(input))
	fmt.Print(output)
}
