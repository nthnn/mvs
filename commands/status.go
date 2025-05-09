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
	"slices"
	"strings"

	"github.com/nthnn/mvs/core"
	"github.com/nthnn/mvs/logger"
	"github.com/nthnn/mvs/utils"

	"github.com/vmihailenco/msgpack/v5"
)

func StatusCommand() {
	repo, err := utils.FindRepoDir()
	if err != nil {
		logger.Error("Not an MVS repository (up to mount point)")
		return
	}

	if err := os.Chdir(repo); err != nil {
		logger.Error("Failed to change directory: %v", err)
		return
	}

	staged := utils.LoadIndex()
	headHash := utils.ResolveHead()

	var lastCommit core.Commit
	if headHash != "" {
		blob, err := utils.ReadCompressed(
			filepath.Join(core.CommitsPath, headHash),
		)

		if err != nil {
			logger.Error(
				"Failed to read commit %s: %v",
				headHash,
				err,
			)
			return
		}

		if err := msgpack.Unmarshal(
			blob,
			&lastCommit,
		); err != nil {
			logger.Error(
				"Failed to parse commit %s: %v",
				headHash,
				err,
			)
			return
		}
	}

	modes := make(map[string]os.FileMode, len(lastCommit.Files))
	hashes := make(map[string]string, len(lastCommit.Files))

	for _, file := range lastCommit.Files {
		modes[file.Path] = file.Mode
		hashes[file.Path] = file.Hash
	}

	var modified, untracked []string
	filepath.Walk(".", func(
		path string,
		info os.FileInfo,
		err error,
	) error {
		if err != nil ||
			info.IsDir() ||
			strings.HasPrefix(path, core.RepoDir) ||
			strings.HasSuffix(path, ".sig") {
			return nil
		}

		data, err := os.ReadFile(path)

		if err != nil {
			modified = append(modified, path)
			return nil
		}

		hashNow := utils.Hash(data)
		if origHash, tracked := hashes[path]; tracked || staged[path] {
			if hashNow != origHash || info.Mode() != modes[path] {
				modified = append(modified, path)
			}
		} else {
			untracked = append(untracked, path)
		}

		return nil
	})

	for path := range hashes {
		if staged[path] {
			continue
		}

		if _, err := os.Stat(path); os.IsNotExist(err) {
			modified = append(modified, path)
		}
	}

	currentBranch, err := utils.GetCurrentBranch()
	if err != nil {
		logger.Error(
			"Error retrieving current branch: %v",
			err,
		)
		return
	}

	deleted := []string{}
	for path := range hashes {
		if staged[path] {
			continue
		}

		if _, err := os.Stat(path); os.IsNotExist(err) {
			modified = append(modified, path)
			deleted = append(deleted, path)
		}
	}

	fmt.Println("On " + stylizeText(currentBranch) + " branch:\r\n")

	fmt.Println("\x1b[1mStaged\x1b[0m:")
	stagedNum := len(staged)

	if stagedNum == 0 {
		fmt.Println("  \x1b[3m(No staged files found)\x1b[0m")
	} else {
		paths := make([]string, 0, len(staged))
		for p := range staged {
			paths = append(paths, p)
		}

		for index, path := range paths {
			prefix := "├─"
			if len(paths) == 1 {
				prefix = "──"
			} else if index == len(paths)-1 {
				prefix = "└─"
			}

			fmt.Printf(
				"  \x1b[1m%s\x1b[0m %s\r\n",
				prefix,
				stylizeText(path),
			)
		}
	}
	fmt.Println()

	fmt.Println("\x1b[1mModified\x1b[0m:")
	modifiedNum := len(modified)

	if modifiedNum == 0 {
		fmt.Println("  \x1b[3m(No modified files found)\x1b[0m")
	} else {
		for index, modifiedFile := range modified {
			prefix := "├─"
			if modifiedNum == 1 {
				prefix = "──"
			} else if index == modifiedNum-1 {
				prefix = "└─"
			}

			if slices.Contains(deleted, modifiedFile) {
				fmt.Printf(
					"  \x1b[1m%s\x1b[0m %s \x1b[3m(deleted)\x1b[0m\r\n",
					prefix,
					stylizeText(modifiedFile),
				)
			} else {
				fmt.Printf(
					"  \x1b[1m%s\x1b[0m %s\r\n",
					prefix,
					stylizeText(modifiedFile),
				)
			}
		}
	}
	fmt.Println()

	fmt.Println("\x1b[1mUntracked\x1b[0m:")
	untrackedNum := len(untracked)

	if untrackedNum == 0 {
		fmt.Println("  \x1b[3m(No modified files found)\x1b[0m")
	} else {
		for index, untrackedFile := range untracked {
			prefix := "├─"
			if untrackedNum == 1 {
				prefix = "──"
			} else if index == untrackedNum-1 {
				prefix = "└─"
			}

			fmt.Printf(
				"  \x1b[1m%s\x1b[0m %s\r\n",
				prefix,
				stylizeText(untrackedFile),
			)
		}
	}
	fmt.Println()

}
