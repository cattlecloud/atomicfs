// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

package atomicfs

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/shoenig/test/must"
)

func TestWriter_WriteFile(t *testing.T) {
	dir := t.TempDir()

	writer := New(
		WithTempDir(dir),
		WithTempExtension(".tempfile"),
		WithMode(0600),
	)

	input := strings.NewReader("foobar")
	filePath := filepath.Join(dir, "out.txt")
	must.NoError(t, writer.WriteFile(filePath, input))

	content, err := os.ReadFile(filePath)
	must.NoError(t, err)
	must.Eq(t, "foobar", string(content))

	info, err := os.Stat(filePath)
	must.NoError(t, err)
	must.Eq(t, os.FileMode(0600), info.Mode())
}

func TestWriter_WriteFile_defaults(t *testing.T) {
	writer := New()
	must.Eq(t, os.TempDir(), writer.tempDir)
	must.Eq(t, "tmp", writer.tempExt)
	must.Eq(t, os.FileMode(0660), writer.mode)
}

func TestWriter_WriteFile_missingDirectory(t *testing.T) {
	writer := New(WithTempDir(t.TempDir()))

	filePath := filepath.Join(t.TempDir(), "missing", "out.txt")
	err := writer.WriteFile(filePath, strings.NewReader("x"))
	must.Error(t, err)
}
