/*
 * Copyright 2025 Nathanne Isip
 * This file is part of Minimal Versioning System (https://github.com/nthnn/mvs)
 * This code is licensed under MIT license (see LICENSE for details)
 */

package logger

import (
	"fmt"
	"math"
	"strings"
	"sync"
	"time"
)

var logMutex sync.Mutex

func Colorize(
	text string,
	startColor, endColor [3]int,
	isBold bool,
	isItalic bool,
) string {
	var coloredText strings.Builder
	textLength := len(text)

	if textLength == 0 {
		return ""
	}

	if isBold {
		coloredText.WriteString("\x1b[1m")
	}

	if isItalic {
		coloredText.WriteString("\x1b[3m")
	}

	for i, char := range text {
		ratio := float64(i) / float64(textLength-1)

		r := int(math.Round(
			float64(startColor[0]) +
				ratio*float64(endColor[0]-startColor[0]),
		))
		g := int(math.Round(
			float64(startColor[1]) +
				ratio*float64(endColor[1]-startColor[1]),
		))
		b := int(math.Round(
			float64(startColor[2]) +
				ratio*float64(endColor[2]-startColor[2]),
		))

		coloredText.WriteString(fmt.Sprintf(
			"\x1b[38;2;%d;%d;%dm%c",
			r, g, b,
			char,
		))
	}

	coloredText.WriteString("\x1b[0m")
	return coloredText.String()
}

func Log(data string, args ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()

	fmt.Printf(
		"%s \u001b[44m  LOG  \u001b[0m ",
		time.Now().Format(time.RFC3339),
	)

	fmt.Printf(data, args...)
	fmt.Println()
}
