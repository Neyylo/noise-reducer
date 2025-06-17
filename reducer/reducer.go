package reducer

import (
	"fmt"
	"math"
	"math/cmplx"
)

func fft(x []complex128) []complex128 {
	N := len(x)
	if N <= 1 {
		return x
	}
	even := make([]complex128, N/2)
	odd := make([]complex128, N/2)
	for i := 0; i < N/2; i++ {
		even[i] = x[2*i]
		odd[i] = x[2*i+1]
	}
	Feven := fft(even)
	Fodd := fft(odd)
	X := make([]complex128, N)
	for k := 0; k < N/2; k++ {
		t := cmplx.Exp(complex(0, -2*math.Pi*float64(k)/float64(N))) * Fodd[k]
		X[k] = Feven[k] + t
		X[k+N/2] = Feven[k] - t
	}
	return X
}

func ifft(X []complex128) []complex128 {
	N := len(X)
	conj := make([]complex128, N)
	for i := range X {
		conj[i] = cmplx.Conj(X[i])
	}
	fftRes := fft(conj)
	out := make([]complex128, N)
	for i := range fftRes {
		out[i] = cmplx.Conj(fftRes[i]) / complex(float64(N), 0)
	}
	return out
}

func ProcessFFTBandpass(inputPath, outputPath string, sampleRate int, lowHz, highHz float64) error {
	samples, format, err := iohelper.ReadWavFile(inputPath)
	if err != nil {
		return fmt.Errorf("read failed: %w", err)
	}
	N := len(samples)
	x := make([]complex128, N)
	for i := range samples {
		x[i] = complex(samples[i], 0)
	}
	X := fft(x)
	for i := 0; i < N; i++ {
		freq := float64(i) * float64(sampleRate) / float64(N)
		if freq < lowHz || freq > highHz {
			X[i] = 0
		}
	}
	xFiltered := ifft(X)
	output := make([]float64, N)
	max := 0.0
	for i := range xFiltered {
		output[i] = real(xFiltered[i])
		if math.Abs(output[i]) > max {
			max = math.Abs(output[i])
		}
	}
	if max > 0 {
		for i := range output {
			output[i] /= max
		}
	}
	err = iohelper.WriteWavFile(outputPath, output, format)
	if err != nil {
		return fmt.Errorf("write failed: %w", err)
	}
	return nil
}
