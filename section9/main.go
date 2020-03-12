package main

import (
	"fmt"
	"io"
	"os"
)

func open() {
	file, err := os.Create("textfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "Hello Ikuhiroo !\n")
}

func read() {
	file, err := os.Open("textfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("read file:")
	io.Copy(os.Stdout, file)
}

func main() {
	// open()
	// read()
	os.Mkdir("setting", 0755)
	os.MkdirAll("setting/myapp/networksettings", 0755)
}
