/*
 * Copyright 2025 Nathanne Isip
 * This file is part of Minimal Versioning System (https://github.com/nthnn/mvs)
 * This code is licensed under MIT license (see LICENSE for details)
 */

package commands

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/nthnn/mvs/core"
	"github.com/nthnn/mvs/logger"
	"github.com/nthnn/mvs/utils"
	"github.com/vmihailenco/msgpack/v5"
)

func CheckoutCommand(args []string) {
	if len(args) == 0 {
		logger.Error("Checkout requires name parameter.")
		return
	}

	target := args[0]
	refPath := filepath.Join(core.RefsPath, target)

	var hashVal string
	if data, err := os.ReadFile(refPath); err == nil {
		if err := utils.Verify(refPath); err != nil {
			logger.Error("Blob ref data tampering detected.")
			return
		}

		hashVal = strings.TrimSpace(string(data))
		utils.AtomicWriteFile(
			core.HeadFile,
			[]byte("ref: refs/heads/"+target),
			0644,
		)

		utils.Sign(core.HeadFile)
	} else {
		hashVal = target
	}

	path := filepath.Join(core.CommitsPath, hashVal)
	if err := utils.Verify(path); err != nil {
		logger.Error("Commit data tampering detected.")
		return
	}
	raw, _ := utils.ReadCompressed(path)

	var c core.Commit
	msgpack.Unmarshal(raw, &c)

	for _, fe := range c.Files {
		data, _ := utils.ReadCompressed(filepath.Join(
			core.ObjectsPath,
			utils.HashFile(fe.Path),
		))
		utils.AtomicWriteFile(fe.Path, data, fe.Mode)
	}

	logger.Log("Checked out: %s", args[0])
}
