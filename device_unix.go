// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

//go:build !windows

package atomicfs

import (
	"errors"
	"os"
	"syscall"
)

// deviceID returns the filesystem device identifier of the file or directory
// at path. It is used to ensure the tmp and destination directories reside on
// the same device, which is required for an atomic rename.
func deviceID(path string) (uint64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return 0, errors.New("atomicfs: unable to determine device id")
	}

	return uint64(stat.Dev), nil
}
