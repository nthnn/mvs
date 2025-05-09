package utils

import (
	"os"
	"path/filepath"
)

func AtomicWriteFile(
	path string,
	data []byte,
	perm os.FileMode,
) error {
	dir := filepath.Dir(path)
	temp, err := os.CreateTemp(dir, ".__mvs_tmp_*")

	if err != nil {
		return err
	}

	tempName := temp.Name()
	defer os.Remove(tempName)

	if _, err := temp.Write(data); err != nil {
		temp.Close()
		return err
	}

	if err := temp.Close(); err != nil {
		return err
	}

	return os.Rename(tempName, path)
}
