package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoConvert(input string) (string, error) {
	if strings.ContainsAny(input, ".-") { // проверяем наличие символов точки и тире
		return morse.ToText(input), nil
	}
	return morse.ToMorse(input), nil
}
