package calculator_sdk

type Calculation interface {
	Add(value []int64) int64
	Min(value []int64) int64
	Multiply(numberOne float64, numberTwo float64) float64
	Division(numberOne float64, numberTwo float64) float64
}
