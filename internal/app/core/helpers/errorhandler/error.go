package errorhandler

import (
	"fmt"
	"net/http"
	"runtime"
	"user-flow/internal/app/core/helpers/response"
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

func AbortWithError(abort func(int, any), err error, errStatuses map[error]int) {
	if code, ok := errStatuses[err]; ok {
		abort(code, response.ErrorResponse(err.Error()))

		return
	}

	abort(http.StatusBadRequest, response.ErrorResponse(response.BadRequest))

	return
}
