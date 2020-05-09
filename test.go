package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("err 1", err)
		return
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("err 2", err)
		return
	}
	fmt.Println(string(content))
}
