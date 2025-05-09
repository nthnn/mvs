/*
 * Copyright 2025 Nathanne Isip
 * This file is part of Minimal Versioning System (https://github.com/nthnn/mvs)
 * This code is licensed under MIT license (see LICENSE for details)
 */

package utils

import (
	"io"
	"os"

	"github.com/klauspost/pgzip"
)

func WriteCompressed(path string, data []byte) error {
	file, _ := os.Create(path)
	defer file.Close()

	pgzipWriter := pgzip.NewWriter(file)
	pgzipWriter.SetConcurrency(100000, 10)
	defer pgzipWriter.Close()

	_, e := pgzipWriter.Write(data)
	return e
}

func ReadCompressed(path string) ([]byte, error) {
	file, _ := os.Open(path)
	defer file.Close()

	pgzipReader, _ := pgzip.NewReader(file)
	defer pgzipReader.Close()

	return io.ReadAll(pgzipReader)
}
