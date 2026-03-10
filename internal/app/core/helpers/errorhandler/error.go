package errorhandler

import (
	"fmt"
	"runtime"
)

func FailOnError(err error, msg string) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		message := fmt.Sprintf("%s: %s", msg, err.Error())
		fmt.Println(message)

		fmt.Printf("Ошибка произошла в файле: %s, строке: %d, сообщение: %s\n", file, line, message)
	}
}

func Fatal(err error, msg string) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		message := fmt.Sprintf("%s: %s", msg, err.Error())
		fmt.Println(message)

		fmt.Printf("Ошибка произошла в файле: %s, строке: %d, сообщение: %s\n", file, line, message)
	}
}
