package commands

import (
	"flag"
	"os"
	"path/filepath"
	"time"

	"github.com/nthnn/mvs/core"
	"github.com/nthnn/mvs/logger"
	"github.com/nthnn/mvs/utils"
	"github.com/vmihailenco/msgpack/v5"
)

func CommitCommand() {
	flags := flag.NewFlagSet("commit", flag.ExitOnError)
	message := flags.String("m", "", "commit message")

	flags.Parse(os.Args[2:])
	if *message == "" {
		logger.Error("Commit message required: -m <message>")
		return
	}

	index := utils.LoadIndex()
	if len(index) == 0 {
		logger.Warning("No changes to commit.")
		return
	}

	parent := utils.ResolveHead()
	now := time.Now().UTC().Format(time.RFC3339)

	author := core.GlobalConfiguration.Name
	email := core.GlobalConfiguration.Email

	if author == "" {
		author = os.Getenv("USER")
	}

	commit := core.Commit{
		Parent:    parent,
		Message:   *message,
		Author:    author,
		Email:     email,
		Timestamp: now,
	}

	baseFiles := make(map[string]core.FileEntry)
	if parent != "" {
		parentBlob, err := utils.ReadCompressed(
			filepath.Join(
				core.CommitsPath,
				parent,
			),
		)

		if err != nil {
			logger.Error(
				"Failed to read parent commit %s: %v",
				parent,
				err,
			)
			return
		}

		var parentCommit core.Commit
		if err := msgpack.Unmarshal(
			parentBlob,
			&parentCommit,
		); err != nil {
			logger.Error(
				"Failed to parse parent commit %s: %v",
				parent,
				err,
			)
			return
		}

		for _, fe := range parentCommit.Files {
			baseFiles[fe.Path] = fe
		}
	}

	for path := range index {
		info, err := os.Stat(path)
		if err != nil {
			logger.Error("Cannot stat %s: %v", path, err)
			return
		}

		data, err := os.ReadFile(path)
		if err != nil {
			logger.Error("Cannot read %s: %v", path, err)
			return
		}

		blobHash := utils.Hash(data)
		if err := utils.WriteCompressed(
			filepath.Join(core.ObjectsPath, blobHash),
			data,
		); err != nil {
			logger.Error(
				"Failed to write object %s: %v",
				blobHash,
				err,
			)
			return
		}

		baseFiles[path] = core.FileEntry{
			Path: path,
			Mode: info.Mode(),
			Hash: blobHash,
		}
	}

	for _, fe := range baseFiles {
		commit.Files = append(commit.Files, fe)
	}

	initialBlob, err := msgpack.Marshal(&commit)
	if err != nil {
		logger.Error("Failed to marshal commit: %v", err)
		return
	}

	commit.Hash = utils.Hash(initialBlob)
	finalBlob, err := msgpack.Marshal(&commit)

	if err != nil {
		logger.Error(
			"Failed to marshal commit with hash: %v",
			err,
		)
		return
	}

	commitPath := filepath.Join(core.CommitsPath, commit.Hash)
	if err := utils.WriteCompressed(commitPath, finalBlob); err != nil {
		logger.Error("Failed to write commit object: %v", err)
		return
	}

	if err := utils.Sign(commitPath); err != nil {
		logger.Error("Failed to sign commit object: %v", err)
		return
	}

	ref := utils.CurrentRefPath()
	if err := utils.AtomicWriteFile(ref, []byte(commit.Hash), 0644); err != nil {
		logger.Error("Failed to update ref: %v", err)
		return
	}
	if err := utils.Sign(ref); err != nil {
		logger.Error("Failed to sign ref: %v", err)
		return
	}

	// 7) Clear the index
	if err := utils.AtomicWriteFile(core.IndexFile, []byte{}, 0644); err != nil {
		logger.Error("Failed to clear index: %v", err)
		return
	}

	logger.Log(
		"Committed: %s",
		stylizeText(commit.Hash[:16]+"…"),
	)
}
