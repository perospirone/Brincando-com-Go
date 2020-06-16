package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader, text string

func main() {
	getArray()
}

func inputKey() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	//fmt.Println(text)

	return text
}

func scopeLocal() string {
	var a string

	a = "dasdasa"

	return a
}

func getArray() {
	var arr = new([2]string)

	arr[0] = scopeLocal()
	arr[1] = inputKey()

	print(arr[1])
}
