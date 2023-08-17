package usecase

type CalculationOptions struct{}

func (c CalculationOptions) Add(values []int64) int64 {
	result := int64(0)
	if len(values) > 0 {
		for _, v := range values {
			result += v
		}
	}
	return result
}

func (c CalculationOptions) Min(values []int64) int64 {
	result := int64(0)
	if len(values) == 1 {
		return values[0]
	}

	for k, v := range values {
		if k == 0 {
			result = v
			continue
		}
		result -= v
	}
	return result
}

func (c CalculationOptions) Multiply(numberOne float64, numberTwo float64) float64 {
	return float64(numberOne * numberTwo)
}

func (c CalculationOptions) Division(numberOne float64, numberTwo float64) float64 {
	return numberOne / numberTwo
}
