package main

import (
	"fmt"
	"io"
	"os"

	"github.com/tcolgate/mp3"
)

func main() {
	t := 0.0

	r, err := os.Open("audio/audio.mp3")
	if err != nil {
		fmt.Println(err)
		return
	}

	d := mp3.NewDecoder(r)
	var f mp3.Frame
	skipped := 0

	for {

		if err := d.Decode(&f, &skipped); err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}

		t = t + f.Duration().Seconds()
	}

	fmt.Printf("%s duration: %v\n", r.Name(), t)

}
