/*
 * Copyright 2025 Nathanne Isip
 * This file is part of Minimal Versioning System (https://github.com/nthnn/mvs)
 * This code is licensed under MIT license (see LICENSE for details)
 */

package commands

import (
	"github.com/nthnn/mvs/logger"
	"github.com/nthnn/mvs/utils"
)

func RemoveCommand(paths []string) {
	index := utils.LoadIndex()

	for _, path := range paths {
		if index[path] {
			delete(index, path)
			logger.Log("Removed from staging: %s", path)
		} else {
			logger.Error("Not staged: %s", path)
		}
	}

	utils.SaveIndex(index)
}
