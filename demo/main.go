package main

import (
	"log"

	"github.com/Neyylo/noise-reducer/reducer"
)

func main() {
	input := "./audio/heart.wav"
	output := "./audio/test-Output/outyut.wav"
	alpha := 0.002

	err := reducer.ProcessLowPass(input, output, alpha)
	if err != nil {
		log.Fatal("Erreur pendant le traitement :", err)
	}

	log.Println("Traitement terminé avec succès !")
}
