// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

package fs

import (
	"os"
	"time"
)

type FileSystem interface {
	Create(name string) (File, error)
	Mkdir(name string, perm os.FileMode) error
	MkdirAll(path string, perm os.FileMode) error
	Open(name string) (File, error)
	OpenFile(name string, flag int, perm os.FileMode) (File, error)
	Remove(name string) error
	RemoveAll(path string) error
	Rename(previous, next string) error
	Stat(name string) (os.FileInfo, error)
}

type File interface {
	Chdir() error
	Chmod(mode os.FileMode) error
	Chown(uid, gid int) error
	Close() error
	Fd() uintptr
	Name() string
	Read(b []byte) (n int, err error)
	ReadAt(b []byte, off int64) (n int, err error)
	Readdir(n int) ([]os.FileInfo, error)
	Readdirnames(n int) (names []string, err error)
	Seek(offset int64, whence int) (ret int64, err error)
	SetDeadline(t time.Time) error
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
	Stat() (os.FileInfo, error)
	Sync() error
	Truncate(size int64) error
	Write(b []byte) (n int, err error)
	WriteAt(b []byte, off int64) (n int, err error)
	WriteString(s string) (n int, err error)
}
