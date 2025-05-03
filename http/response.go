package http

func Error(message string) interface{} {
	return map[string]string{"error": message}
}
