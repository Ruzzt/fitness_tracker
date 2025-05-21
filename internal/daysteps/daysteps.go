package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
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
	stepsStr := strings.TrimSpace(times[0])
	if stepsStr == "" {
		return 0, 0, errors.New("steps value is empty")
	}

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, 0, errors.New("Error converting the number of steps")
	}
	if steps <= 0 {
		return 0, 0, errors.New("The number of steps must be greater than zero")
	}

	durationStr := strings.TrimSpace(times[1])
	if durationStr == "" {
		return 0, 0, errors.New("duration value is empty")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, 0, errors.New("Invalid duration format")
	}
	if duration <= 0 {
		return 0, 0, errors.New("Duration must be greater than zero")
	}

	log.Println("некорректный формат")

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Printf("Error parsing package: %v\n", err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distantion := stepLength * float64(steps)
	distantionKm := distantion / mInKm

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println("Ошибка при расчёте калорий:", err)
		return ""
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distantionKm, calories,
	)

}
