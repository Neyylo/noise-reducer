package iohelper

import (
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

func ReadFile(path string) ([]float64, *audio.Format, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	d := wav.NewDecoder(f)
	buf, err := d.FullPCMBuffer()
	if err != nil {
		return nil, nil, err
	}

	floatBuf := buf.AsFloatBuffer()
	samples := floatBuf.Data
	floatSamples := make([]float64, len(samples))

	for i, s := range floatBuf.Data {
		floatSamples[i] = s / scale
	}

	format := floatBuf.Format

	return floatSamples, format, nil
}
