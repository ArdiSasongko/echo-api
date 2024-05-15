package helper

func DefaultEmpty(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
