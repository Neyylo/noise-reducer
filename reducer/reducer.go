package reducer

import (
	"fmt"

	"github.com/Neyylo/noise-reducer/reducer/filters"
	"github.com/Neyylo/noise-reducer/reducer/utils/iohelper"
)

func ProcessLowPass(inputPath, outputPath string, alpha float64) error {

	samples, format, err := iohelper.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("échec lecture fichier : %w", err)
	}

	filtered := filters.LowPassFilter(samples, alpha)

	err = iohelper.WriteFile(outputPath, filtered, format)
	if err != nil {
		return fmt.Errorf("échec écriture fichier : %w", err)
	}

	return nil
}

func ProcessFFTBandPass(inputPath, outputPath string, lowHz, highHz float64) error {
	samples, format, err := iohelper.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("lecture fichier : %w", err)
	}

	filtered := filters.FFTBandPass(samples, format.SampleRate, lowHz, highHz)

	if err := iohelper.WriteFile(outputPath, filtered, format); err != nil {
		return fmt.Errorf("écriture fichier : %w", err)
	}
	return nil
}
