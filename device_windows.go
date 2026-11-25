// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

//go:build windows

package atomicfs

import "errors"

// deviceID is unsupported on Windows, where same-device atomic rename and
// directory fsync do not apply.
func deviceID(string) (uint64, error) {
	return 0, errors.New("atomicfs: atomic writes are not supported on windows")
}
