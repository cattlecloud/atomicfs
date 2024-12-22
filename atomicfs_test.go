// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

package atomicfs

import (
	"path/filepath"
	"strings"
	"testing"

	"cattlecloud.net/go/atomicfs/fs"
	"cattlecloud.net/go/atomicfs/sys"
	"github.com/shoenig/test/must"
)

func TestFileWriter_WriteFile(t *testing.T) {
	tmpDir := t.TempDir()

	writer := New(Options{
		TmpDirectory: tmpDir,
		TmpExtension: ".tempfile",
		Mode:         0600,
		FS:           fs.New(),
		Sys:          sys.New(),
	})

	input := strings.NewReader("foobar")
	filePath := filepath.Join(tmpDir, "out.txt")
	err := writer.WriteFile(filePath, input)
	must.NoError(t, err)
}
