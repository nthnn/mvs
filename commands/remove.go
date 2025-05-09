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
