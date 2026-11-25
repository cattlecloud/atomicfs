atomicfs
========

A library for atomic filesystem operations in Go.

[![Go Reference](https://pkg.go.dev/badge/cattlecloud.net/go/atomicfs.svg)](https://pkg.go.dev/cattlecloud.net/go/atomicfs)
[![License](https://img.shields.io/github/license/cattlecloud/atomicfs?color=7C00D8&style=flat-square&label=License)](https://github.com/cattlecloud/atomicfs/blob/main/LICENSE)
[![Build](https://img.shields.io/github/actions/workflow/status/cattlecloud/atomicfs/ci.yaml?style=flat-square&color=0FAA07&label=Tests)](https://github.com/cattlecloud/atomicfs/actions/workflows/ci.yaml)

# Overview

The `cattlecloud.net/go/atomicfs/v2` module provides a package for performing
atomic filesystem operations.

#### Reading material

- https://rcrowley.org/2010/01/06/things-unix-can-do-atomically.html

# Getting Started

The `atomicfs` package can be added to a project with `go get`.

```shell
go get cattlecloud.net/go/atomicfs/v2@latest
```

```go
import "cattlecloud.net/go/atomicfs/v2"
```

#### Examples

Written files default to mode `0660` (read-write for the owner and group).
Pass `atomicfs.WithMode` to override it:

```go
writer := atomicfs.New(
    atomicfs.WithTempDir("/tmp"),
    atomicfs.WithMode(0600), // custom file mode
)

input := strings.NewReader("hello, world")
_ = writer.WriteFile("/path/to/output.txt", input)
```

# Contributing

The `cattlecloud.net/go/atomicfs/v2` module is always improving with new
features and error corrections. For contributing bug fixes and new features
please file an issue.

# License

The `cattlecloud.net/go/atomicfs/v2` module is open source under the
[BSD-3-Clause](LICENSE) license.
