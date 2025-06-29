package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parsedStr := strings.Split(datastring, ",")
	if len(parsedStr) != 2 {
		return fmt.Errorf("a datastring does not have enough data or too big")
	}

	ds.Steps, err = strconv.Atoi(parsedStr[0])
	if err != nil {
		return err
	}
	if ds.Steps <= 0 {
		return fmt.Errorf("the step count is zero or less")
	}

	ds.Duration, err = time.ParseDuration(parsedStr[1])
	if err != nil {
		return err
	}
	if ds.Duration <= 0 {
		return fmt.Errorf("wrong time duration")
	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Personal.Height <= 0 {
		return "", fmt.Errorf("height is zero or less")
	}
	distanceKM := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		log.Println(err)
		return "", err
	}

	result := fmt.Sprintf("Количество шагов: %d.\n"+
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n", ds.Steps, distanceKM, calories)

	return result, nil
}
