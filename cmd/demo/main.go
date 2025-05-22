package main

import (
	"log"

	processing "github.com/Neyylo/noise-reducer/reducer/process"
)

func main() {
	input := "./audio/heart.wav"
	output := "./audio/test-Output/output0002.wav"
	alpha := 0.002

	err := processing.Process(input, output, alpha)
	if err != nil {
		log.Fatal("Erreur pendant le traitement :", err)
	}

	log.Println("Traitement terminé avec succès !")
}
