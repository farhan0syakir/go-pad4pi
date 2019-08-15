package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/farhan0syakir/go-pad4pi"
)

var myKeypad = [][]string{
	{"1", "2"},
	{"3", "4"},
}

var rowPins = []int{35, 36}
var colPins = []int{37, 38}

func handler(key string) {
	log.Println(key)
}

func main() {
	mykeypad := pad4pi.NewKeypad(rowPins, colPins, myKeypad)
	mykeypad.RegisterKeyPressHandler(handler)
	defer mykeypad.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	defer signal.Stop(quit)

	// In a real application the main thread would do something useful here.
	// But we'll just run for a minute then exit.
	fmt.Println("Watching Pin")
	select {
	case <-time.After(time.Minute):
	case <-quit:
	}

}
