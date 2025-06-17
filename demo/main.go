package main

import (
	"log"

	"github.com/Neyylo/noise-reducer/reducer"
)

func main() {
	input := "./audio/heart.wav"
	output := "./audio/outputFFT.wav"

	// OU filtre fréquentiel (FFT)
	err := reducer.ProcessFFTLowPass(input, output, 200.0)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Traitement terminé avec succès !")
}
