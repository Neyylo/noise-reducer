package main

import (
	"fmt"
	"os"

	"github.com/go-audio/wav"
)

var (
	path string = "./audio/heart.wav"
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
	f.Close()
	fmt.Println("Old audio ->", d)

	out, err := os.Create("./audio/test-Output/output.wav")
	if err != nil {
		panic(fmt.Sprintf("Couldn't create the OUTPUT - %v", err))
	}

	e := wav.NewEncoder(out,
		buf.Format.SampleRate,
		int(d.BitDepth),
		buf.Format.NumChannels,
		int(d.WavAudioFormat))
	if err = e.Write(buf); err != nil {
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
