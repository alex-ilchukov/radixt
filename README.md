# Radixt

Implementation of Radix Tree and lookups with use of the structure, written in 
Go (Golang).


## Short description of Radix Tree and examples

Radix Tree is a special kind of tree which stores string parts (chunks) in its
nodes. One can read about them in [Wiki][wiki-radix-tree], here goes an example
with HTTP methods:
```
                               ""
          /                  /  |     \       \      \
       "P"          "OPTIONS" "DELETE" "TRACE" "HEAD" "GET"
     /  |  \
"OST"  "UT" "ATCH"
```

The trees can be used to do efficient lookup of a string in predefined string
list. One can mark its nodes, and find the marks, going over the nodes and their chunks by incoming bytes.

[wiki-radix-tree]: https://en.wikipedia.org/wiki/Radix_tree


## Goals of the project

The goals of project can be found in [corresponding document](./GOALS.md). In
short, those would be the following:
*   [X] definition of tree interface;
*   [X] null tree implementation;
*   [X] generic tree implementation;
*   [X] implementation of lookup process;
*   [ ] compactified implementations;
*   [ ] fossilization of a tree into Golang code.


## Installation

To install Radixt package, you need to install Go (of version 1.19 or higher), 
get the package and import it.

### Getting the package

```sh
go get -u github.com/alex-ilchukov/radixt
```

### Import of the package

```go
import "github.com/alex-ilchukov/radixt"
```


## Usage

One would need a radix tree, which implements [`radixt.Tree`](./tree.go) 
interface, to invoke lookups of strings. Such a tree can be created, for 
example, from strings with use of [`generic`](./generic) subpackage:
```go
import "github.com/alex-ilchukov/radixt/generic"
…
tree := generic.New("authorization", "content-length", "content-type")
```

If one has a proper tree, it can be used for lookups in the following way:
```go
import "github.com/alex-ilchukov/radixt/lookup"
…
l := lookup.New(tree) // Here tree is an implmentation of radixt.Tree
```

To lookup a string:
```go
s := "content-length"
l.Reset()
for i := 0; i < len(s); i++ {
  // Can be changed to just l.Feed(s[i]) if early break is unrequired
  if !l.Feed(s[i]) {
    break
  }
}
```

To lookup a consequence of bytes:
```
import "io"
…
var r io.ByteReader = …
l.Reset()
for {
  b, err := r.ReadByte()
  if err != nil {
    … // break, return or whatelse
  }
  if !l.Feed(b) {
    … // same
  }
}
```

To get boolean true or false as status of the lookup:
```go
l.Found()
```

To get a "mark" (just an integer) of the node if it is found, or -1 if it is 
not:
```go
l.Tree().NodeMark(l.Node())
```

More examples of usage can be found in [./examples] directory.
