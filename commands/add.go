/*
 * Copyright 2025 Nathanne Isip
 * This file is part of Minimal Versioning System (https://github.com/nthnn/mvs)
 * This code is licensed under MIT license (see LICENSE for details)
 */

package commands

import (
	"os"

	"github.com/nthnn/mvs/logger"
	"github.com/nthnn/mvs/utils"
)

func AddCommand(paths []string) {
	index := utils.LoadIndex()
	for _, path := range paths {
		info, err := os.Stat(path)

		if err == nil && !info.IsDir() {
			index[path] = true
			logger.Log("Added: %s", path)
		} else {
			logger.Warning("Skipped: %s", path)
		}
	}

	utils.SaveIndex(index)
}
