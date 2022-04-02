package main

import (
	"fmt"
	"os"
)

const trial string = "trial.txt"

func writeToFile() {
	f, err := os.OpenFile(trial, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	d := "A new line to the file\n"
	_, err = fmt.Fprint(f, d)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func readFile() {
	data, err := os.ReadFile("trial.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func main() {
	for i := 0; i < 100; i++ {
		writeToFile()
	}
	readFile()
}
