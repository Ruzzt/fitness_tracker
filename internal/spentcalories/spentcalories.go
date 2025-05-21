package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	datas := strings.Split(data, ",")
	if len(datas) != 3 {
		return 0, "", 0, errors.New("The length is less than 3")
	}
	steps, err := strconv.Atoi(datas[0])
	if err != nil {
		return 0, "", 0, errors.New("Invalid steps count")
	}

	if steps <= 0 {
		return 0, "", 0, errors.New("Steps must be greater than zero")
	}

	activity := strings.TrimSpace(datas[1])

	duration, err := time.ParseDuration(datas[2])
	if err != nil {
		return 0, "", 0, errors.New("invalid duration format")
	}

	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	lenSteps := stepLengthCoefficient * height
	return float64(steps) * lenSteps / mInKm

}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distantion := distance(steps, height)
	hours := duration.Hours()
	return distantion / hours
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		return "", errors.New("Invalid steps count")
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	hours := duration.Hours()

	var calories float64

	switch strings.ToLower(activity) {
	case "бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	case "ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("unknown type of training: " + activity)
	}

	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity, hours, dist, speed, calories,
	)

	return result, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("the number of steps is less than zero")
	}
	if weight <= 0 {
		return 0, errors.New("the number of weight is less than zero")
	}
	if height <= 0 {
		return 0, errors.New("the number of height is less than zero")
	}
	if duration <= 0 {
		return 0, errors.New("the number of duration is less than zero")
	}
	average := meanSpeed(steps, height, duration)
	munites := duration.Minutes()
	calories := weight * average * munites

	return calories / minInH, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("the number of steps is less than zero")
	}
	if weight <= 0 {
		return 0, errors.New("the number of weight is less than zero")
	}
	if height <= 0 {
		return 0, errors.New("the number of height is less than zero")
	}
	if duration <= 0 {
		return 0, errors.New("the number of duration is less than zero")
	}
	average := meanSpeed(steps, height, duration)
	munites := duration.Minutes()
	calories := (weight * average * munites) / minInH

	return calories * walkingCaloriesCoefficient, nil
}
