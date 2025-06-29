package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parsedStr := strings.Split(datastring, ",")

	if len(parsedStr) != 3 {
		return fmt.Errorf("the datastring does not have enough data or too big")
	}

	steps, err := strconv.Atoi(parsedStr[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return fmt.Errorf("the step count is zero or less")
	}

	parsedDuration, err := time.ParseDuration(parsedStr[2])
	if err != nil {
		return err
	}
	if parsedDuration <= 0 {
		return fmt.Errorf("wrong time duration")
	}
	t.Steps = steps
	t.TrainingType = parsedStr[1]
	t.Duration = parsedDuration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	var (
		resultCalories float64
		resultDuration float64
		resultDistance float64
		resultSpeed    float64
		err            error
	)
	switch t.TrainingType {
	case "Ходьба":
		resultCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Бег":
		resultCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	resultDuration = time.Duration(t.Duration).Hours()
	resultDistance = spentenergy.Distance(t.Steps, t.Height)
	resultSpeed = spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	result := fmt.Sprintf("Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n", t.TrainingType, resultDuration, resultDistance, resultSpeed, resultCalories,
	)
	return result, nil
}
