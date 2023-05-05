package main

import (
	"fmt"
	"time"
)

func say(s string) {
	time.Sleep(10000 * time.Millisecond)
	fmt.Println(s)
}

func main(){
	say("hello")
	go say("world")
}