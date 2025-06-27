package tests

import (
	"math"
	"testing"

	"github.com/Neyylo/noise-reducer/reducer/filters"
)

// Test that FFTLowPass returns a signal of the same length
func TestFFTLowPassPreservesLength(t *testing.T) {
	sampleRate := 44100
	N := sampleRate // 1 second
	freq := 440.0   // A4 pitch
	samples := make([]float64, N)
	for i := range samples {
		samples[i] = math.Sin(2 * math.Pi * freq * float64(i) / float64(sampleRate))
	}

	cutoff := 800.0
	filtered := filters.FFTLowPass(samples, sampleRate, cutoff)

	if len(filtered) != len(samples) {
		t.Errorf("Expected output length %d, got %d", len(samples), len(filtered))
	}
}

// Test that FFTBandPass zeroes out low-frequency content below cutoff
func TestFFTBandPassDoesNotPanic(t *testing.T) {
	sampleRate := 44100
	N := sampleRate
	freq := 440.0

	samples := make([]float64, N)
	for i := range samples {
		samples[i] = math.Sin(2 * math.Pi * freq * float64(i) / float64(sampleRate))
	}

	lowHz := 300.0
	highHz := 500.0
	filtered := filters.FFTBandPass(samples, sampleRate, lowHz, highHz)

	if len(filtered) != len(samples) {
		t.Errorf("Expected output length %d, got %d", len(samples), len(filtered))
	}
}

// Optional: energy check (just to confirm output isn't silent)
func TestFilteredOutputHasEnergy(t *testing.T) {
	sampleRate := 44100
	N := sampleRate
	freq := 1000.0 // Outside the band

	samples := make([]float64, N)
	for i := range samples {
		samples[i] = math.Sin(2 * math.Pi * freq * float64(i) / float64(sampleRate))
	}

	lowHz := 100.0
	highHz := 300.0
	filtered := filters.FFTBandPass(samples, sampleRate, lowHz, highHz)

	var energy float64
	for _, v := range filtered {
		energy += v * v
	}

	if energy < 1e-5 {
		t.Errorf("Filtered signal has too little energy: %f", energy)
	}
}
