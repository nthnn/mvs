/*
 * Copyright 2025 Nathanne Isip
 * This file is part of Minimal Versioning System (https://github.com/nthnn/mvs)
 * This code is licensed under MIT license (see LICENSE for details)
 */

package core

import (
	"os"

	"github.com/nthnn/mvs/logger"
	"gopkg.in/yaml.v3"
)

const (
	RepoDir       = ".mvs"
	ObjectsPath   = ".mvs/objects"
	CommitsPath   = ".mvs/commits"
	RefsPath      = ".mvs/refs/heads"
	HeadFile      = ".mvs/HEAD"
	IndexFile     = ".mvs/index"
	KeyDirectory  = ".mvs/keys"
	PrivateKeyPem = ".mvs/keys/id_ed25519"
	PublicKeyPem  = ".mvs/keys/id_ed25519.pub"

	GlobalConf = "~/.local/mvs/globals.yaml"
)

var GlobalConfiguration GlobalConfig

func LoadGlobalConfig() {
	data, err := os.ReadFile(GlobalConf)
	if err != nil {
		GlobalConfiguration.Email = "noemail@noemail.com"
		GlobalConfiguration.PrivateKey = "MVS"
		GlobalConfiguration.PublicKey = "MVS"

		logger.Warning(
			"Error: %v",
			err,
		)
		return
	}

	var conf GlobalConfig
	if err := yaml.Unmarshal(data, &conf); err != nil {
		logger.Warning(
			"Could not parse global config: %v",
			err,
		)
		return
	}

	GlobalConfiguration = conf
}
