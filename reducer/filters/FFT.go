package filters

import (
	"math"
	"math/cmplx"
)

func FFT(x []complex128) []complex128 {
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
	Feven := FFT(even)
	Fodd := FFT(odd)

	X := make([]complex128, N)
	for k := 0; k < N/2; k++ {
		t := cmplx.Exp(complex(0, -2*math.Pi*float64(k)/float64(N))) * Fodd[k]
		X[k] = Feven[k] + t
		X[k+N/2] = Feven[k] - t
	}
	return X
}

func IFFT(X []complex128) []complex128 {
	N := len(X)
	conj := make([]complex128, N)
	for i := 0; i < N; i++ {
		conj[i] = cmplx.Conj(X[i])
	}
	fft := FFT(conj)
	out := make([]complex128, N)
	for i := 0; i < N; i++ {
		out[i] = cmplx.Conj(fft[i]) / complex(float64(N), 0)
	}
	return out
}

func FFTLowPass(samples []float64, sampleRate int, cutoffHz float64) []float64 {
	N := len(samples)
	x := make([]complex128, N)
	for i := 0; i < N; i++ {
		x[i] = complex(samples[i], 0)
	}

	X := FFT(x)

	// Appliquer le filtre
	for i := 0; i < N; i++ {
		freq := math.Abs(float64(i))
		if freq*float64(sampleRate)/float64(N) > cutoffHz {
			X[i] = 0
		}
	}

	xFiltered := IFFT(X)
	out := make([]float64, N)
	max := 0.0
	for i := 0; i < N; i++ {
		out[i] = real(xFiltered[i])
		if math.Abs(out[i]) > max {
			max = math.Abs(out[i])
		}
	}

	// Normalisation
	if max > 0 {
		for i := range out {
			out[i] /= max
		}
	}

	return out
}
