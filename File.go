package main

import (
	// "os"
	"io/ioutil"
	"log"
	"fmt"
)

func main() {
	// file, _ := os.Create("test.txt")
	// // _ = os.Rename("test.txt", "test2.txt")
	// // _ = os.Remove("test2.txt")
	// file.Close()

	// err := ioutil.WriteFile("test.txt", []byte("Hello\n"), 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
	fmt.Printf("Fmt read: %s\n", data)
	log.Printf("Data read: %s\n", data)
}