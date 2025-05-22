package main

import (
	"fmt"
	"os"

	noise "github.com/Neyylo/noise-reducer/reducer"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

var (
	path  string  = "./audio/heart.wav"
	scale float64 = 32767.0
)

func main() {
	f, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("Couldn't open the file - %v", err))
	}

	d := wav.NewDecoder(f)
	buf, err := d.FullPCMBuffer()
	if err != nil {
		panic(err)
	}

	os.Remove("output.wav")

	floatBuf := buf.AsFloatBuffer()
	samples := floatBuf.Data
	floatSamples := make([]float64, len(samples))

	for i, s := range floatBuf.Data {
		floatSamples[i] = s / scale
	}

	//fmt.Println(floatSamples)
	format := floatBuf.Format
	filtered := noise.LowPassFilter(floatSamples, 0.1)

	f.Close()

	intSamples := make([]int, len(filtered))

	scale := 32767.0

	for i, sample := range filtered {
		intSamples[i] = int(sample * scale)
	}

	intBuf := &audio.IntBuffer{
		Data:           intSamples,
		Format:         format,
		SourceBitDepth: 16,
	}

	out, err := os.Create("./audio/test-Output/output.wav")

	if err != nil {
		panic(fmt.Sprintf("Couldn't create the OUTPUT - %v", err))
	}

	e := wav.NewEncoder(out,
		intBuf.Format.SampleRate,
		int(intBuf.SourceBitDepth),
		intBuf.Format.NumChannels,
		int(d.WavAudioFormat))

	if err = e.Write(intBuf); err != nil {
		panic(err)
	}

	if err = e.Close(); err != nil {
		panic(err)
	}

	out.Close()

	out, err = os.Open("./audio/test-Output/output.wav")

	if err != nil {
		panic(err)
	}

	d2 := wav.NewDecoder(out)
	d2.ReadInfo()
	fmt.Println("New file ->", d2)
	out.Close()

}
