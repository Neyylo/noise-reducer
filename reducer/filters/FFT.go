// file: filters/fft.go

package filters

import (
	"math"

	"gonum.org/v1/gonum/dsp/fourier"
)

// FFTLowPass applies a full-spectrum low-pass filter
func FFTLowPass(samples []float64, sampleRate int, cutoffHz float64) []float64 {
	N := len(samples)

	// 1) Build a complex slice from real samples
	x := make([]complex128, N)
	for i, v := range samples {
		x[i] = complex(v, 0)
	}

	// 2) Complex FFT (size N → spectrum size N)
	fft := fourier.NewCmplxFFT(N)
	S := fft.Coefficients(nil, x) // []complex128 of length N

	// 3) Remove frequencies above cutoffHz
	for k := range S {
		freq := float64(k) * float64(sampleRate) / float64(N)
		if freq > cutoffHz {
			S[k] = 0
		}
	}

	// 4) Full iFFT back to time domain
	Y := fft.Sequence(nil, S) // []complex128 of length N

	// 5) Extract real part + normalize
	out := make([]float64, N)
	var maxAmp float64
	for i, c := range Y {
		r := real(c)
		out[i] = r
		if a := math.Abs(r); a > maxAmp {
			maxAmp = a
		}
	}
	if maxAmp > 0 {
		scale := 1.0 / maxAmp
		for i := range out {
			out[i] *= scale
		}
	}
	return out
}

// FFTBandPass applies full spectrum band-pass filter (complex FFT)
func FFTBandPass(samples []float64, sampleRate int, lowHz, highHz float64) []float64 {
	N := len(samples)

	x := make([]complex128, N)
	for i, v := range samples {
		x[i] = complex(v, 0)
	}

	fft := fourier.NewCmplxFFT(N)
	S := fft.Coefficients(nil, x)

	// Keep only frequencies between lowHz | highHz
	for k := range S {
		freq := float64(k) * float64(sampleRate) / float64(N)
		if freq < lowHz || freq > highHz {
			S[k] = 0
		}
	}

	Y := fft.Sequence(nil, S)

	out := make([]float64, N)
	var maxAmp float64
	for i, c := range Y {
		r := real(c)
		out[i] = r
		if a := math.Abs(r); a > maxAmp {
			maxAmp = a
		}
	}
	if maxAmp > 0 {
		scale := 1.0 / maxAmp
		for i := range out {
			out[i] *= scale
		}
	}
	return out
}
