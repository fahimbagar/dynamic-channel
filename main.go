package main

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"
)

type MyChannels map[string]chan int

var myChan = make(MyChannels)

func NextMessage(wantChannels []string) int64 {
	// create select array with dynamic size
	cases := make([]reflect.SelectCase, len(wantChannels) + 1)

	// insert select default so the receiver can wait
	cases[0] = reflect.SelectCase{Dir: reflect.SelectDefault}

	for i, wantChan := range wantChannels {
		fmt.Printf("want from Chan %s\n", wantChan)
		cases[i+1] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(myChan[wantChan])}
	}

	for {
		// select the available one from select case
		index, value, _ := reflect.Select(cases)
		if index == 0 {
			fmt.Printf("%+v waiting for 2 seconds before checking again\n", wantChannels)
			// wait for 2 seconds
			<-time.After(2 * time.Second)
		} else {
			// if it's not select default (index 0), return the value
			return value.Int()
		}
	}
}

func main() {
	/*
	From the documentation (http://golang.org/doc/effective_go.html#channels):
	If the channel is unbuffered, the sender blocks until the receiver has received the value.
	If the channel has a buffer, the sender blocks only until the value has been copied to the buffer;
	if the buffer is full, this means waiting until some receiver has retrieved a value.
	*/
	myChan["CH1"] = make(chan int, 10)
	myChan["CH2"] = make(chan int, 10)

	fmt.Println("inserting 2 to CH2")
	myChan["CH2"]<-2

	fmt.Println("inserting 1 to CH1")
	myChan["CH1"]<-1

	go func() {
		fmt.Printf("result from CH1: %d\n", NextMessage([]string{"CH1"}))
	}()

	fmt.Println("inserting 2 to CH2")
	myChan["CH2"]<-2

	fmt.Println("inserting 2 to CH2")
	myChan["CH2"]<-2

	go func() {
		fmt.Printf("result from CH1 & CH2: %d\n", NextMessage([]string{"CH1", "CH2"}))
	}()

	go func() {
		fmt.Printf("result from CH1: %d\n", NextMessage([]string{"CH1"}))
	}()

	fmt.Println("waiting for 5 minute before sending int to CH1")
	<-time.After(5 * time.Second)
	fmt.Println("inserting 1 to CH1")
	myChan["CH1"]<-1

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Press CTRL+C to exit")
	<-c
	fmt.Println("Exited")
	os.Exit(0)
}