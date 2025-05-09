/*
 * Copyright 2025 Nathanne Isip
 * This file is part of Minimal Versioning System (https://github.com/nthnn/mvs)
 * This code is licensed under MIT license (see LICENSE for details)
 */

package core

import "os"

type GlobalConfig struct {
	Name          string `yaml:"name"`
	Email         string `yaml:"email"`
	PublicKey     string `yaml:"public_key"`
	PrivateKey    string `yaml:"private_key"`
	DefaultBranch string `yaml:"branch"`
}

type FileEntry struct {
	Path string      `msgpack:"path"`
	Mode os.FileMode `msgpack:"mode"`
	Hash string      `msgpack:"hash"`
}

type Commit struct {
	Hash      string      `msgpack:"hash"`
	Parent    string      `msgpack:"parent"`
	Message   string      `msgpack:"message"`
	Author    string      `msgpack:"author"`
	Email     string      `msgpack:"email"`
	Timestamp string      `msgpack:"when"`
	Files     []FileEntry `msgpack:"files"`
}
