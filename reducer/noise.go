package noise

func LowPassFilter(input []float64, alpha float64) []float64 {
	output := make([]float64, len(input))
	if len(input) == 0 {
		return output
	}
	for i := 1; i < len(input); i++ {
		output[i] = alpha*input[i] + (float64(1)-alpha)*output[i-1]
	}

	return output
}
