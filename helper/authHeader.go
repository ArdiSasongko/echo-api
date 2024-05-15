package helper

import (
	"errors"
	"strings"
)

func ValidToken(h string) (string, error) {
	if h == "" || !strings.HasPrefix(h, "Bearer") {
		return "", errors.New("missing authHeader")
	}

	token := strings.TrimSpace(strings.TrimPrefix(h, "Bearer"))

	if token == "" {
		return "", errors.New("token missing")
	}

	return token, nil
}
