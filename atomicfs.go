// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

// Package atomicfs provides tools for doing all-or-nothing atomic filesystem
// operations in Go.
package atomicfs

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// A Writer writes the contents of an io.Reader to a destination file
// atomically, or fails without leaving corrupt data in place of the
// destination file.
//
// The primary design goal of a Writer is correctness and reliability.
//
// Performance is not a consideration.
type Writer struct {
	tempDir string
	tempExt string
	mode    os.FileMode
}

// An Option configures the behavior of a Writer.
type Option func(*Writer)

// WithTempDir sets the directory where temporary files are created during the
// process of writing a file in preparation for an atomic rename.
//
// The temp dir must be on the same filesystem device as the destination file,
// otherwise errors will be returned during the write.
//
// Defaults to os.TempDir().
func WithTempDir(dir string) Option {
	return func(w *Writer) {
		if dir != "" {
			w.tempDir = dir
		}
	}
}

// WithTempExtension sets the extension used to demarcate the temporary files
// created during a write from other normal files. The extension is appended to
// the end of the destination file name; any leading dots are stripped.
//
// Defaults to "tmp".
func WithTempExtension(ext string) Option {
	ext = strings.TrimLeft(ext, ".")
	return func(w *Writer) {
		if ext != "" {
			w.tempExt = ext
		}
	}
}

// WithMode sets the file mode of the destination file to be written.
//
// Defaults to 0660 (read-write for the owner and group).
func WithMode(mode os.FileMode) Option {
	return func(w *Writer) {
		if mode != 0 {
			w.mode = mode
		}
	}
}

// New creates a Writer configured with the given options. The options are
// applied in order; later options override earlier ones.
func New(options ...Option) *Writer {
	w := &Writer{
		tempDir: os.TempDir(),
		tempExt: "tmp",
		mode:    os.FileMode(0660),
	}
	for _, option := range options {
		option(w)
	}
	return w
}

// WriteFile writes the contents of source into a file at filePath, atomically.
func (w *Writer) WriteFile(filePath string, source io.Reader) error {
	fileDir := filepath.Dir(filePath)
	fileName := filepath.Base(filePath)

	// 1) check the target directory is on the same filesystem device
	// as the configured tmp directory - otherwise atomic operations
	// are not possible
	if err := w.checkDevice(fileDir); err != nil {
		return err
	}

	// 2) write the content to a tmp file, also triggering
	// a complete flush (fsync)
	tmpPath, err := w.writeTmp(source, fileName)
	if err != nil {
		if tmpPath != "" {
			_ = os.Remove(tmpPath)
		}
		return err
	}

	// 3) since we know the tmp file and destination file exist on the
	// same device, a filesystem rename will be an atomic operation
	if err := w.rename(tmpPath, filePath); err != nil {
		_ = os.Remove(tmpPath)
		return err
	}

	return nil
}

func (w *Writer) checkDevice(fileDir string) error {
	fileDirDeviceID, err := deviceID(fileDir)
	if err != nil {
		return fmt.Errorf("atomicfs: unable to stat destination directory %s: %w", fileDir, err)
	}

	tmpDirDeviceID, err := deviceID(w.tempDir)
	if err != nil {
		return fmt.Errorf("atomicfs: unable to stat tmp directory %s: %w", w.tempDir, err)
	}

	if fileDirDeviceID != tmpDirDeviceID {
		return errors.New("atomicfs: tmp & destination directories not on same device")
	}

	return nil
}

func (w *Writer) rename(previous, next string) error {
	if err := os.Rename(previous, next); err != nil {
		return fmt.Errorf("atomicfs: unable to rename tmp file %s to %s: %w", previous, next, err)
	}
	return w.syncDir(next)
}

func (w *Writer) syncDir(name string) error {
	directory := filepath.Dir(name)
	f, err := os.Open(directory)
	if err != nil {
		return fmt.Errorf("atomicfs: unable to open directory %s for syncing: %w", directory, err)
	}

	if err := f.Sync(); err != nil {
		return fmt.Errorf("atomicfs: unable to sync directory %s: %w", directory, err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("atomicfs: unable to close directory %s after syncing: %w", directory, err)
	}

	return nil
}

func (w *Writer) writeTmp(source io.Reader, name string) (string, error) {
	tmpName := name + "." + w.tempExt

	f, err := os.CreateTemp(w.tempDir, tmpName)
	if err != nil {
		return "", fmt.Errorf("atomicfs: unable to create tmp file in %s: %w", w.tempDir, err)
	}
	tmpPath := f.Name()

	if err := f.Chmod(w.mode); err != nil {
		_ = f.Close()
		return tmpPath, fmt.Errorf("atomicfs: unable to chmod tmp file in %s: %w", w.tempDir, err)
	}

	if _, err := io.Copy(f, source); err != nil {
		_ = f.Close()
		return tmpPath, fmt.Errorf("atomicfs: unable to copy source into tmp file in %s: %w", w.tempDir, err)
	}

	if err := f.Sync(); err != nil {
		_ = f.Close()
		return tmpPath, fmt.Errorf("atomicfs: unable to fsync tmp file in %s: %w", w.tempDir, err)
	}

	if err := f.Close(); err != nil {
		return tmpPath, fmt.Errorf("atomicfs: unable to close tmp file in %s: %w", w.tempDir, err)
	}

	return tmpPath, nil
}
