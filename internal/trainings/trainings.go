package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	slData := strings.Split(datastring, ",")
	if len(slData) != 3 {
		return errors.New("not enough agruments")
	}
	steps, err := strconv.Atoi(slData[0])
	if err != nil {
		return errors.New("incorect steps data")
	}
	t.Steps = steps
	if slData[1] != "Бег" && slData[1] != "Ходьба" {
		return errors.New("unknown training")
	}
	t.TrainingType = slData[1]

	duration, err := time.ParseDuration(slData[2])
	if err != nil {
		return errors.New("incorect time data")
	}
	t.Duration = duration
	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	trDistance := spentenergy.Distance(t.Steps)
	if trDistance <= 0 {
		return "", errors.New("distanse will be more than zero")
	}
	aveSpeedAct := spentenergy.MeanSpeed(t.Steps, t.Duration)
	var calories float64
	switch t.TrainingType {

	case "Бег":
		calories = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
		if calories <= 0 {
			return "", errors.New("error in calculating calories for running")
		}
	case "Ходьба":
		calories = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if calories <= 0 {
			return "", errors.New("error in calculating calories for walk")
		}
	default:
		return "неизвестный тип тренировки", errors.New("unknown training type")
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		t.TrainingType, t.Duration.Hours(), trDistance, aveSpeedAct, calories), nil
}
