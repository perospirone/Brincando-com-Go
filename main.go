package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader, text string

func main() {
	fmt.Print(getArray())
}

func inputKey() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	return text
}

func scopeLocal() string {
	var local string

	local = "Sou uma variavel local então só existo nesse bloco de codigo\n"

	return local
}

func getArray() *[2]string {
	var arr = new([2]string)

	arr[0] = scopeLocal()
	arr[1] = inputKey()

	return arr
}
