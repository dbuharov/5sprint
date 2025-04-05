package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse() "678,0h50m"
func (ds *DaySteps) Parse(datastring string) (err error) {
	daySl := strings.Split(datastring, ",")
	if len(daySl) != 2 {
		return errors.New("not enough agruments")
	}
	steps, err := strconv.Atoi(daySl[0])
	if err != nil {
		return errors.New("incorect steps data")
	}
	ds.Steps = steps //шаги0
	duration, err := time.ParseDuration(daySl[1])
	if err != nil {
		return errors.New("incorect time data")
	}
	ds.Duration = duration //время1
	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", errors.New("incorect time data")
	}
	actDist := spentenergy.Distance(ds.Steps)
	actCalories := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if actCalories <= 0 {
		return "", errors.New("incorect calories")
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", ds.Steps, actDist, actCalories), nil
}
