package utils

import (
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

func ReadWavFile(path string) ([]float64, *audio.Format, error) {
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
	for i, s := range samples {
		floatSamples[i] = float64(s) / 32767.0
	}
	return floatSamples, floatBuf.Format, nil
}

func WriteWavFile(path string, samples []float64, format *audio.Format) error {
	intSamples := make([]int, len(samples))
	for i, s := range samples {
		intSamples[i] = int(s * 32767.0)
	}
	intBuf := &audio.IntBuffer{
		Data:           intSamples,
		Format:         format,
		SourceBitDepth: 16,
	}
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()
	e := wav.NewEncoder(out, format.SampleRate, 16, format.NumChannels, 1)
	if err := e.Write(intBuf); err != nil {
		return err
	}
	return e.Close()
}
