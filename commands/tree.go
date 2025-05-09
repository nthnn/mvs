/*
 * Copyright 2025 Nathanne Isip
 * This file is part of Minimal Versioning System (https://github.com/nthnn/mvs)
 * This code is licensed under MIT license (see LICENSE for details)
 */

package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nthnn/mvs/core"
	"github.com/nthnn/mvs/logger"
	"github.com/nthnn/mvs/utils"
)

func padString(str string, length uint8) string {
	if len(str) >= int(length) {
		return str
	}

	return str + strings.Repeat(" ", 9-len(str))
}

func currentBranch() string {
	bytes, _ := os.ReadFile(core.HeadFile)
	str := string(bytes)

	if strings.HasPrefix(str, "ref:") {
		return filepath.Base(strings.TrimSpace(
			strings.TrimPrefix(str, "ref: "),
		))
	}

	return ""
}

func TreeCommand() {
	repo, err := utils.FindRepoDir()
	if err != nil {
		logger.Error("Not an MVS repository")
		return
	}
	os.Chdir(repo)

	current := currentBranch()
	entries, _ := os.ReadDir(core.RefsPath)

	var branches []os.DirEntry
	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".sig") {
			continue
		}

		branches = append(branches, entry)
	}

	fmt.Println("\x1b[1mBranches\x1b[0m:")

	num := len(branches)
	if num == 0 {
		fmt.Println("  \x1b[3m(No branches found)\x1b[0m")
		return
	}

	for index, file := range branches {
		prefix := "├─"
		if num == 1 {
			prefix = "──"
		} else if index == len(branches)-1 {
			prefix = "└─"
		}

		branchName := file.Name()
		name := logger.Colorize(
			branchName,
			[3]int{111, 66, 193},
			[3]int{0, 123, 255},
			true, false,
		)

		if branchName == current {
			fmt.Printf(
				"\x1b[1m%s\x1b[0m %s \x1b[3m(current)\x1b[0m\r\n",
				prefix,
				padString(name, 12),
			)
		} else {
			fmt.Printf(
				"\x1b[1m%s\x1b[0m %s\r\n",
				prefix,
				padString(name, 12),
			)
		}
	}
}
