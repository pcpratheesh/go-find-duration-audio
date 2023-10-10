package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-audio/wav"
)

func main() {
	f, err := os.Open("audio/audio.wav")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	dur, err := wav.NewDecoder(f).Duration()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s duration: %s\n", f.Name(), dur)
}
