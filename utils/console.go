/*
 * Copyright 2025 Nathanne Isip
 * This file is part of Minimal Versioning System (https://github.com/nthnn/mvs)
 * This code is licensed under MIT license (see LICENSE for details)
 */

package utils

import (
	"os"

	"golang.org/x/term"
)

func GetSingleChar() (byte, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return 0, err
	}

	defer term.Restore(
		int(os.Stdin.Fd()),
		oldState,
	)

	var buf [1]byte
	_, err = os.Stdin.Read(buf[:])

	if err != nil {
		return 0, err
	}

	return buf[0], nil
}
