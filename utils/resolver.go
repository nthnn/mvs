/*
 * Copyright 2025 Nathanne Isip
 * This file is part of Minimal Versioning System (https://github.com/nthnn/mvs)
 * This code is licensed under MIT license (see LICENSE for details)
 */

package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nthnn/mvs/core"
)

func FindRepoDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(
			dir,
			core.RepoDir,
		)); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}

		dir = parent
	}

	return "", fmt.Errorf("not an MVS repository")
}

func ResolveHead() string {
	bytes, _ := os.ReadFile(core.HeadFile)
	str := string(bytes)

	if strings.HasPrefix(str, "ref:") {
		ref := strings.TrimSpace(strings.TrimPrefix(
			str,
			"ref: ",
		))
		data, _ := os.ReadFile(filepath.Join(
			core.RepoDir,
			ref,
		))

		return strings.TrimSpace(string(data))
	}

	return strings.TrimSpace(str)
}

func CurrentRefPath() string {
	bytes, _ := os.ReadFile(core.HeadFile)
	str := string(bytes)

	if strings.HasPrefix(str, "ref:") {
		ref := strings.TrimSpace(strings.TrimPrefix(
			str,
			"ref: ",
		))

		return filepath.Join(core.RepoDir, ref)
	}

	return filepath.Join(
		core.RefsPath,
		core.GlobalConfiguration.DefaultBranch,
	)
}

func GetCurrentBranch() (string, error) {
	data, err := os.ReadFile(core.HeadFile)
	if err != nil {
		return "", fmt.Errorf(
			"failed to read head file: %w",
			err,
		)
	}

	refLine := strings.TrimSpace(string(data))
	const prefix = "ref: refs/heads/"
	if strings.HasPrefix(refLine, prefix) {
		return strings.TrimPrefix(refLine, prefix), nil
	}

	return "HEAD", nil
}
