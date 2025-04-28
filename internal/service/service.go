package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// AutoConvert автоматически определяет тип строки и возвращает результат
func AutoConvert(input string) (string, error) {
	if strings.ContainsAny(input, ".- ") { // проверяем наличие точек и тире
		return morse.ToText(input), nil // конвертируем код Морзе в текст
	}

	return morse.ToMorse(input), nil // конвертируем текст в код Морзе
}
