atomicfs
========

A library for atomic filesystem operations in Go.

[![Go Reference](https://pkg.go.dev/badge/cattlecloud.net/go/atomicfs.svg)](https://pkg.go.dev/cattlecloud.net/go/atomicfs)
[![License](https://img.shields.io/github/license/cattlecloud/atomicfs?color=7C00D8&style=flat-square&label=License)](https://github.com/cattlecloud/atomicfs/blob/main/LICENSE)
[![Build](https://img.shields.io/github/actions/workflow/status/cattlecloud/atomicfs/ci.yaml?style=flat-square&color=0FAA07&label=Tests)](https://github.com/cattlecloud/atomicfs/actions/workflows/ci.yaml)

# Overview

The `cattlecloud.net/go/atomicfs` module provides a package for performing atomic
filesystem operations.

#### Reading material

- https://rcrowley.org/2010/01/06/things-unix-can-do-atomically.html

# Getting Started

The `atomicfs` package can be added to a project with `go get`.

```shell
go get cattlecloud.net/go/atomicfs@latest
```

```go
import "cattlecloud.net/go/atomicfs"
```

#### Examples

```go
writer := atomicfs.NewFileWriter(atomicfs.Options{
    TmpDirectory: "/tmp",
    Mode:         0600,
})

_ = writer.Write(input, output)
```

# Contributing

The `cattlecloud.net/go/atomicfs` module is always improving with new features and
error corrections. For contributing bug fixes and new features please file an
issue.

# License

The `cattlecloud.net/go/atomicfs` module is open source under the [BSD-3-Clause](LICENSE)
license.
