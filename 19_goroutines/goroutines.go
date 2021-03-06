package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

/*
	The goroutine functions will always be executed
	after the command function is executed
	But we won't know the goroutine functions executing orders.
	They ara running asynchronously.

	I guess the functions which are called in the usual way are put
	into a synchronous queue.
	And the goroutine functions will be put into an asynchronous queue.
	And the asynchronous queue won't be executed in a sequence way.

*/
func main() {

	go f("hello1")

	f("hello2")

	go f("hello3")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
