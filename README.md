# Go thrift parser

[![unittest](https://github.com/curoky/go-thrift-parser/actions/workflows/unittest.yaml/badge.svg)](https://github.com/curoky/go-thrift-parser/actions/workflows/unittest.yaml)
![GitHub license](https://img.shields.io/github/license/curoky/go-thrift-parser)
[![Go Report Card](https://goreportcard.com/badge/github.com/curoky/go-thrift-parser)](https://goreportcard.com/report/github.com/curoky/go-thrift-parser)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/curoky/go-thrift-parser)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/curoky/go-thrift-parser)
[![GoDoc](https://pkg.go.dev/badge/github.com/curoky/go-thrift-parser?status.svg)](https://pkg.go.dev/github.com/curoky/go-thrift-parser?tab=doc)

`go-thrift-parser` is a Go package which parses a thrift file.

## Getting started

import package

```golang
import (
  "fmt"

  "github.com/curoky/go-thrift-parser/parser"
  "github.com/curoky/go-thrift-parser/parser/ast"
)
```

parse thrift file

```golang
p := parser.CreateParser(false, []string{"."})
err := p.RecursiveParse("types.thrift")
if err != nil {
  panic(err)
}
```

visit ast

```golang
doc := p.Thrift.Documents["types.thrift"]
for _, body := range doc.Body {
  switch v := body.([]interface{})[0].(type) {
  case *ast.Include:
    fmt.Printf("include: %s\n", v.Name)
  case *ast.Type:
    fmt.Printf("type: %s\n", v.Name)
  case *ast.Service:
    fmt.Printf("service: %s\n", v.Name)
  }
}
```

result

```text
type: StrType
type: EnumType
type: UnionType
type: StructType
type: OutterStructType
```

## License

See [`LICENSE`](./LICENSE.md)
