/*
 * Copyright 2025 Nathanne Isip
 * This file is part of Minimal Versioning System (https://github.com/nthnn/mvs)
 * This code is licensed under MIT license (see LICENSE for details)
 */

package logger

import (
	"fmt"
	"time"
)

func Warning(data string, args ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()

	fmt.Printf(
		"%s \u001b[43m WARN  \u001b[0m ",
		time.Now().Format(time.RFC3339),
	)

	fmt.Printf(data, args...)
	fmt.Println()
}
