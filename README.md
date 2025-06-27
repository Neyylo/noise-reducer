# noise-reducer

> Go library for noise reduction and audio component extraction using FFT-based filters (low-pass, band-pass).

## 📦 Installation

```bash
go get github.com/Neyylo/noise-reducer
go get gonum.org/v1/gonum/dsp/fourier
```

## 🚀 Features

- Low-pass filtering via `filters.FFTLowPass`
- Band-pass filtering via `filters.FFTBandPass`
- WAV file read/write support (mono & stereo) with `iohelper`
- CLI demo and Go import usage
- Automatic signal normalization

## 🗂️ Project Structure

```
noise-reducer/
├── audio/      # Input/output WAV files
├── demo/      # Example CLI entry point (main.go)
├── reducer/      # Core library
│ ├── filters/        # Signal processing (FFT, IIR filters)
│ │ ├── FFT.go
│ │ └── LowPass.go
│ ├── tests/         # Unit tests
│ │ └── noise_test.go
│ ├── utils/
│ │ └── iohelper/       # Audio I/O utils
│ │ ├── reader.go
│ │ └── writer.go
│ ├── reducer.go     # High-level orchestration (Process functions)
├── go.mod
├── go.sum
└── README.md

## 📖 Example Usage

### Go Application

```go
package main

import (
  "log"
  "github.com/Neyylo/noise-reducer/reducer"
)

func main() {
  input := "./audio/heart.wav"
  output := "./audio/output_bandpass.wav"

  // Apply a 20–150 Hz band-pass filter
  err := reducer.ProcessFFTBandPass(input, output, 20.0, 150.0)
  if err != nil {
    log.Fatal(err)
  }
  log.Println("Processing complete!")
}
```

### CLI Demo

```bash
cd demo
go run main.go
```

## 🔍 API Overview

### filters.FFTLowPass
```go
func FFTLowPass(samples []float64, sampleRate int, cutoffHz float64) []float64
```
Applies a low-pass FFT filter and returns normalized signal.

### filters.FFTBandPass
```go
func FFTBandPass(samples []float64, sampleRate int, lowHz, highHz float64) []float64
```
Applies a band-pass FFT filter between `lowHz` and `highHz`.

### reducer.ProcessLowPass
```go
func ProcessLowPass(inputPath, outputPath string, alpha float64) error
```
Applies an exponential IIR low-pass filter to a WAV file.

### reducer.ProcessFFTLowPass
```go
func ProcessFFTLowPass(inputPath, outputPath string, cutoffHz float64) error
```
Applies an FFT-based low-pass filter to a WAV file.

### reducer.ProcessFFTBandPass
```go
func ProcessFFTBandPass(inputPath, outputPath string, lowHz, highHz float64) error
```
Applies an FFT-based band-pass filter to a WAV file.

## ⚙️ Configuration Notes

- Works with mono or stereo WAV files.
- Stereo files are processed per channel.

## 🛠️ Testing

```bash
go test ./...
```

## 📄 License

Nolan Dugué | Neyylo

