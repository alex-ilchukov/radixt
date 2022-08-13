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

The goals of project can be found in [corresponding document][./GOALS.md]. In
short, those would be the following:
*   [ ] definition of tree interface;
*   [ ] null tree implementation;
*   [ ] generic tree implementation;
*   [ ] implementation of lookup process;
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

TODO: fill the section, when lookup process will be implemented
