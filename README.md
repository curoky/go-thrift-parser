# Go thrift parser

[![License](http://img.shields.io/:license-mit-blue.svg)](https://github.com/curoky/go-thrift-parser/blob/master/LICENSE.md)

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
