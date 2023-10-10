package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-audio/aiff"
)

func main() {
	f, err := os.Open("audio/audio.aiff")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	dur, err := aiff.NewDecoder(f).Duration()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s duration: %s\n", f.Name(), dur)
}
