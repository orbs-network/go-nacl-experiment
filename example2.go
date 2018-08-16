package main

import "io/ioutil"

func main() {
	data := []byte("hello world\n")
	ioutil.WriteFile("./output.txt", data, 0644)
}
