package main

import (
	"fmt"
	"time"

	"github.com/go-flac/go-flac"
)

func main() {
	file, err := flac.ParseFile("audio/audio.flac")
	if err != nil {
		panic(err)
	}
	stream, err := file.GetStreamInfo()
	if err != nil {
		panic(err)
	}

	duration := time.Duration((float64(stream.SampleCount) / float64(stream.SampleRate)) * float64(time.Second))
	fmt.Printf("%s duration: %s\n", "audio.flac", duration)

}
