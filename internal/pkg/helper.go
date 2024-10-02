package pkg

import (
	"fmt"
	"regexp"
	"strings"
)

func FormatPhoneNumber(phone string) (string, error) {
	re := regexp.MustCompile(`[^\d+]`)
	formattedPhone := re.ReplaceAllString(phone, "")

	if !strings.HasPrefix(formattedPhone, "+998") {
		return "", fmt.Errorf("telefon raqami +998 bilan boshlanishi kerak")
	}

	if len(formattedPhone) != 13 {
		return "", fmt.Errorf("telefon raqami noto'g'ri uzunlikda")
	}

	return formattedPhone, nil
}
