package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
)

var reader, text string

/*
type Robo struct {
	name string
	age  int
	gun  string
}*/

type Response struct {
	TradeID int
	Price   string
	Size    string
	Bid     string
	Ask     string
	Volume  string
	Time    string
}

func main() {
	getContent()
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

	local := "Sou uma variavel local então só existo nesse bloco de codigo\n"

	return local
}

func getArray() *[2]string {
	arr := new([2]string)

	arr[0] = scopeLocal()
	arr[1] = inputKey()

	return arr
}

func travelArray(arr *[2]string) {

	count := 1

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

func renameFile(path string, newPath string) error {
	err := os.Rename(path, newPath)

	if err != nil {
		return err
	}

	return nil
}

func numbersInFull() {
	numbers := make(map[int]string)

	numbers[1] = "one"
	numbers[2] = "two"
	numbers[3] = "three"
	numbers[4] = "four"
	numbers[5] = "five"

	fmt.Println(numbers)
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func pointers() {
	x := 5

	var pointer *int

	pointer = &x

	fmt.Printf("Memory address of variable x: %x\n", &x)
	fmt.Printf("Memory address stored in ipointer variable: %x\n", pointer)
	fmt.Printf("Contents of *ipointer variable: %d\n", *pointer)
}

func splitHelloWorld() map[string]string {
	split := make(map[string]string)

	text := "hello world"

	split["h"] = text[0:5]
	split["w"] = text[6:11]

	return split
}

func hi(name string) {
	defer fmt.Print(name)
	fmt.Print("Hi ")
}

func sum(numbers ...int) {
	fmt.Println(numbers)
}

func getContent() Response {
	// json data
	url := "https://api.gdax.com/products/BTC-EUR/ticker"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var data Response

	json.Unmarshal(body, &data)

	return data
}
