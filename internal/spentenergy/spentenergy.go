package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	switch {
	case weight <= 0:
		return 0, fmt.Errorf("weight is zero or less")
	case height <= 0:
		return 0, fmt.Errorf("weight is zero or less")
	case steps <= 0:
		return 0, fmt.Errorf("steps is zero or less") // useless check
	case duration <= 0:
		return 0, fmt.Errorf("duration is zero or less") // useless check again
	}
	return ((weight * MeanSpeed(steps, height, duration) * time.Duration(duration).Minutes()) / minInH) * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	switch {
	case weight <= 0:
		return 0, fmt.Errorf("weight is zero or less")
	case height <= 0:
		return 0, fmt.Errorf("weight is zero or less")
	case steps <= 0:
		return 0, fmt.Errorf("steps is zero or less") // useless check
	case duration <= 0:
		return 0, fmt.Errorf("duration is zero or less") // useless check again
	}

	return (weight * MeanSpeed(steps, height, duration) * time.Duration(duration).Minutes()) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	return Distance(steps, height) / time.Duration(duration).Hours()
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 {
		return 0
	}
	if height <= 0 {
		return 0
	}
	return (float64(steps) * (height * stepLengthCoefficient)) / mInKm
}
