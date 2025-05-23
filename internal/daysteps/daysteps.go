package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Ruzzt/fitness_tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	times := strings.Split(data, ",")
	if len(times) != 2 {
		return 0, 0, errors.New("Insufficient data")
	}

	steps, err := strconv.Atoi(times[0])
	if err != nil {
		return 0, 0, fmt.Errorf("conversion error: %w", err)
	}
	if steps <= 0 {
		return 0, 0, errors.New("The number of steps must be greater than zero")
	}

	duration, err := time.ParseDuration(times[1])
	if err != nil {
		return 0, 0, errors.New("Invalid duration format")
	}
	if duration <= 0 {
		return 0, 0, errors.New("Duration must be greater than zero")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		return fmt.Sprintf("Ошибка парсинга: %v", err)
	}

	if steps <= 0 {
		return "Количество шагов должно быть больше нуля"
	}

	distantion := stepLength * float64(steps)
	distantionKm := distantion / mInKm

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return fmt.Sprintf("Ошибка при расчёте калорий: %v", err)
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distantionKm, calories,
	)
}
