package message

import (
	"runtime"
)

func GetMessage() (string, string, string) {
	var message string
	var message2 string

	switch runtime.GOOS {
	case "linux":
		message = "You're a good person"

	case "openbsd":
		message = "You're a good interesting person"

	case "freebsd":
		message = "You're a good interesting elderly person"

	case "windows":
		message = "You're not a good person"
		message2 = "file path backslashes..."

	case "darwin":
		message = "You're caught in the mind virus"

	case "ios":
		message = "You're caught in the mind virus"

	case "android":
		message = "You do with what you can, so you are a good person"

	default:
		message = "You're probably a good interesting person"
	}

	return message, message2, runtime.GOOS
}
