package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var reader, text string

var arr *[2]string

func main() {
	writeFile("senha.txt", "senhas fodas daleeeee klasjdlksakdsa")
}

func inputKey() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	return text
}

func validInput(input string) bool {
	var emptyString bool

	if input == " " {
		fmt.Println("empty string")
		emptyString = true
	}

	return emptyString
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

func travelArray(arr *[2]string) {
	var count int

	count = 1

	for i := 0; i <= count; i++ {
		fmt.Println(arr[i])
	}
}

func oneUpToTen() {
	i := 1
	max := 10

	for i <= max {
		fmt.Println(i)
		i++
	}
}

func fileExists() {
	if _, err := os.Stat("senha.txt"); err == nil {
		fmt.Printf("File exists\n")
	} else {
		fmt.Printf("File does not exist\n")
	}
}

func readFile(path string) string {
	b, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
	}

	str := string(b)

	return str
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeFile(path string, words string) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(words)

	return nil
}