package reducer

import (
	"fmt"

	"github.com/Neyylo/noise-reducer/reducer/filters"
	"github.com/Neyylo/noise-reducer/reducer/utils/iohelper"
)

func ProcessLowPass(inputPath, outputPath string, alpha float64) error {

	samples, format, err := iohelper.ReadWavFile(inputPath)
	if err != nil {
		return fmt.Errorf("échec lecture fichier : %w", err)
	}

	filtered := filters.LowPassFilter(samples, alpha)

	err = iohelper.WriteWavFile(outputPath, filtered, format)
	if err != nil {
		return fmt.Errorf("échec écriture fichier : %w", err)
	}

	return nil
}
