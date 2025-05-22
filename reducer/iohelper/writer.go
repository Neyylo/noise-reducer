package iohelper

import (
	"fmt"
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

var scale float64 = 32767.0

func WriteWavFile(path string, samples []float64, format *audio.Format) error {
	intSamples := make([]int, len(samples))
	for i, sample := range samples {
		if sample > 1.0 {
			sample = 1.0
		} else if sample < -1.0 {
			sample = -1.0
		}
		intSamples[i] = int(sample * scale)
	}

	intBuf := &audio.IntBuffer{
		Data:           intSamples,
		Format:         format,
		SourceBitDepth: 16,
	}

	out, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("erreur création fichier : %w", err)
	}
	defer out.Close()

	encoder := wav.NewEncoder(out,
		format.SampleRate,
		16,
		format.NumChannels,
		1,
	)

	if err = encoder.Write(intBuf); err != nil {
		return fmt.Errorf("erreur écriture audio : %w", err)
	}

	if err = encoder.Close(); err != nil {
		return fmt.Errorf("erreur fermeture fichier : %w", err)
	}

	return nil
}
