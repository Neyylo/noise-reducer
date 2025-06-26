// fichier : filters/fft.go

package filters

import (
	"math"

	"gonum.org/v1/gonum/dsp/fourier"
)

// FFTLowPass applique un filtre passe-bas complet
func FFTLowPass(samples []float64, sampleRate int, cutoffHz float64) []float64 {
	N := len(samples)

	// 1) Construire un slice complexe à partir de tes échantillons réels
	x := make([]complex128, N)
	for i, v := range samples {
		x[i] = complex(v, 0)
	}

	//FFT complexe (taille N → spectre taille N)
	fft := fourier.NewCmplxFFT(N)
	S := fft.Coefficients(nil, x) // []complex128 de longueur N

	//Coupe les fréquences > cutoffHz
	for k := range S {
		freq := float64(k) * float64(sampleRate) / float64(N)
		if freq > cutoffHz {
			S[k] = 0
		}
	}

	//iFFT complète vers domaine temporel
	Y := fft.Sequence(nil, S) //[]complex128 de longueur N

	//Récupère la partie réelle + normalisation
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

// FFTBandPass applique un filtre passe-bande complet (spectre complexe).
func FFTBandPass(samples []float64, sampleRate int, lowHz, highHz float64) []float64 {
	N := len(samples)

	x := make([]complex128, N)
	for i, v := range samples {
		x[i] = complex(v, 0)
	}

	fft := fourier.NewCmplxFFT(N)
	S := fft.Coefficients(nil, x)

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
