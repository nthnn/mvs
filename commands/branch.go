/*
 * Copyright 2025 Nathanne Isip
 * This file is part of Minimal Versioning System (https://github.com/nthnn/mvs)
 * This code is licensed under MIT license (see LICENSE for details)
 */

package commands

import (
	"path/filepath"

	"github.com/nthnn/mvs/core"
	"github.com/nthnn/mvs/logger"
	"github.com/nthnn/mvs/utils"
)

func BranchRepo(name string) {
	head := utils.ResolveHead()
	utils.AtomicWriteFile(
		filepath.Join(core.RefsPath, name),
		[]byte(head),
		0644,
	)

	utils.Sign(filepath.Join(core.RefsPath, name))
}

func BranchCommand(args []string) {
	if len(args) == 0 {
		logger.Error("Branch requires name parameter.")
		return
	}

	name := args[0]
	BranchRepo(name)

	logger.Log("Branch created: %s", name)
}
