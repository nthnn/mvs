/*
 * Copyright 2025 Nathanne Isip
 * This file is part of Minimal Versioning System (https://github.com/nthnn/mvs)
 * This code is licensed under MIT license (see LICENSE for details)
 */

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

func AmendCommand() {
	flags := flag.NewFlagSet("amend", flag.ExitOnError)
	msg := flags.String("m", "", "new commit message")

	flags.Parse(os.Args[2:])
	if *msg == "" {
		logger.Error(
			"Amend requires a new message: -m <message>",
		)
		return
	}

	headHash := utils.ResolveHead()
	if headHash == "" {
		logger.Error("No commit to amend")
		return
	}

	commitPath := filepath.Join(core.CommitsPath, headHash)
	if err := utils.Verify(commitPath); err != nil {
		logger.Error(
			"Cannot verify existing commit: %v",
			err,
		)
		return
	}

	raw, err := utils.ReadCompressed(commitPath)
	if err != nil {
		logger.Error(
			"Failed to read commit %s: %v",
			headHash,
			err,
		)
		return
	}

	var old core.Commit
	if err := msgpack.Unmarshal(raw, &old); err != nil {
		logger.Error(
			"Failed to parse commit %s: %v",
			headHash,
			err,
		)
		return
	}

	amended := core.Commit{
		Parent:    old.Parent,
		Message:   *msg,
		Author:    old.Author,
		Email:     old.Email,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Files:     old.Files,
	}

	blob1, err := msgpack.Marshal(&amended)
	if err != nil {
		logger.Error(
			"Failed to marshal amended commit: %v",
			err,
		)
		return
	}

	amended.Hash = utils.Hash(blob1)
	blob2, err := msgpack.Marshal(&amended)
	if err != nil {
		logger.Error(
			"Failed to marshal amended commit with hash: %v",
			err,
		)
		return
	}

	newPath := filepath.Join(
		core.CommitsPath,
		amended.Hash,
	)

	if err := utils.WriteCompressed(
		newPath,
		blob2,
	); err != nil {
		logger.Error("Failed to write amended commit: %v", err)
		return
	}

	if err := utils.Sign(newPath); err != nil {
		logger.Error("Failed to sign amended commit: %v", err)
		return
	}

	ref := utils.CurrentRefPath()
	if err := utils.AtomicWriteFile(
		ref,
		[]byte(amended.Hash),
		0644,
	); err != nil {
		logger.Error("Failed to update ref: %v", err)
		return
	}

	if err := utils.Sign(ref); err != nil {
		logger.Error("Failed to sign ref: %v", err)
		return
	}

	logger.Log(
		"Amended commit: %s → %s",
		stylizeText(headHash[:16]+"…"),
		stylizeText(amended.Hash[:16]+"…"),
	)
}
