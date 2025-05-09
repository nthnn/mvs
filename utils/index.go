/*
 * Copyright 2025 Nathanne Isip
 * This file is part of Minimal Versioning System (https://github.com/nthnn/mvs)
 * This code is licensed under MIT license (see LICENSE for details)
 */

package utils

import (
	"os"
	"strings"

	"github.com/nthnn/mvs/core"
)

func LoadIndex() map[string]bool {
	indices := map[string]bool{}
	bytes, _ := os.ReadFile(core.IndexFile)

	for _, l := range strings.Split(
		string(bytes),
		"\n",
	) {
		if l != "" {
			indices[l] = true
		}
	}

	return indices
}

func SaveIndex(idx map[string]bool) {
	var list []string
	for item := range idx {
		list = append(list, item)
	}

	AtomicWriteFile(
		core.IndexFile,
		[]byte(strings.Join(list, "\n")),
		0644,
	)
}
