package logger

import (
	"fmt"
	"time"
)

func Error(data string, args ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()

	fmt.Printf(
		"%s \u001b[41m ERROR \u001b[0m ",
		time.Now().Format(time.RFC3339),
	)

	fmt.Printf(data, args...)
	fmt.Println()
}
