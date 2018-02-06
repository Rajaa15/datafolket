package main

import (
	"./infinite"
	"os/signal"
	"os"
	"fmt"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	go func() {
		<-sigs
		fmt.Printf("You pressed ctrl + C. User interrupted infinite loop.")
		os.Exit(0)
	}()
	Loop()
}

func Loop() {
	infinite.Infinite01("Look, I can count forever:")
}
