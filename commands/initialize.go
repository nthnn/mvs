package commands

import (
	"os"
	"path/filepath"

	"github.com/nthnn/mvs/core"
	"github.com/nthnn/mvs/logger"
	"github.com/nthnn/mvs/utils"
)

func InitializeCommand() {
	dirs := []string{
		core.RepoDir,
		core.ObjectsPath,
		core.CommitsPath,
		core.RefsPath,
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			logger.Error("Failed to create %s: %v", dir, err)
			return
		}
	}

	if err := utils.AtomicWriteFile(
		core.IndexFile,
		[]byte{},
		0o644,
	); err != nil {
		logger.Error("Failed to write index: %v", err)
		return
	}

	branch := core.GlobalConfiguration.DefaultBranch
	if branch == "" {
		branch = "main"
	}

	branchPath := filepath.Join(core.RefsPath, branch)
	if err := utils.AtomicWriteFile(
		branchPath,
		[]byte(""),
		0o644,
	); err != nil {
		logger.Error("Failed to write branch ref: %v", err)
		return
	}
	utils.Sign(branchPath)

	headContent := "ref: refs/heads/" + branch
	if err := utils.AtomicWriteFile(
		core.HeadFile,
		[]byte(headContent),
		0o644,
	); err != nil {
		logger.Error("Failed to write HEAD: %v", err)
		return
	}
	utils.Sign(core.HeadFile)

	if err := utils.EnsureKeys(); err != nil {
		logger.Error("Failed to generate signing keys: %v", err)
		return
	}
	logger.Log(
		"Ed25519 signing keys generated at %s/keys/",
		core.RepoDir,
	)

	logger.Log("Initialized empty mvs repository in .mvs")
	if core.GlobalConfiguration.DefaultBranch != "" {
		BranchRepo(core.GlobalConfiguration.DefaultBranch)
	} else {
		BranchRepo("main")
	}
}
