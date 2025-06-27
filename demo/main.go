package main

import (
	"log"

	"github.com/Neyylo/noise-reducer/reducer"
)

func main() {
	in := "./audio/heart.wav"
	out := "./audio/outputBand800Hz.wav"

	if err := reducer.ProcessFFTBandPass(in, out, 20.0, 200.0); err != nil {
		log.Fatal(err)
	}

	log.Println("Traitement band-pass terminé !")
}
