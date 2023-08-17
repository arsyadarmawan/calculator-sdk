package usecase_test

import (
	"github.com/arsyadarmawan/calculator-sdk/usecase"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddSuccess(t *testing.T) {
	option := usecase.CalculationOptions{}

	numbers := []int64{1, 2, 3, 4, 5}
	expectedNumber := int64(15)
	val := option.Add(numbers)
	assert.Equal(t, val, expectedNumber)
}

func TestAddFail(t *testing.T) {
	option := usecase.CalculationOptions{}

	numbers := []int64{1, 2, 3, 4, 5}
	expectedNumber := int64(16)
	val := option.Add(numbers)
	assert.NotEqual(t, val, expectedNumber)
}

func TestMinSuccess(t *testing.T) {
	option := usecase.CalculationOptions{}

	numbers := []int64{20, 2, 3, 5, 5}
	expectedNumber := int64(5)
	val := option.Min(numbers)
	assert.Equal(t, val, expectedNumber)
}

func TestMinSuccessOnlyOneNumber(t *testing.T) {
	option := usecase.CalculationOptions{}
	numbers := []int64{20}
	expectedNumber := int64(20)
	val := option.Min(numbers)
	assert.Equal(t, val, expectedNumber)
}

func TestMultiplySuccess(t *testing.T) {
	opt := usecase.CalculationOptions{}
	numberOne := float64(20.1)
	numberTwo := float64(4)
	expectedNum := float64(80.4)
	actualRes := opt.Multiply(numberOne, numberTwo)
	assert.Equal(t, actualRes, expectedNum)
}

func TestMultiplyErr(t *testing.T) {
	opt := usecase.CalculationOptions{}
	numberOne := float64(20.1)
	numberTwo := float64(4)
	expectedNum := float64(80)
	actualRes := opt.Multiply(numberOne, numberTwo)
	assert.NotEqual(t, actualRes, expectedNum)
}

func TestDivisionSuccess(t *testing.T) {
	opt := usecase.CalculationOptions{}
	numberOne := float64(20)
	numberTwo := float64(4)
	expectedNum := float64(5)
	actualRes := opt.Division(numberOne, numberTwo)
	assert.Equal(t, actualRes, expectedNum)
}
