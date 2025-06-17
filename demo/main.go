package main

import (
	"log"

	"github.com/Neyylo/noise-reducer/reducer"
)

func main() {
	err := reducer.ProcessFFTBandpass("./audio/heart.wav", "./audio/output_filtered.wav", 44100, 40, 250)
	if err != nil {
		log.Fatal(err)
	}
}
