# Goals

The project has some background and, due to it, a goals to achieve through
steps. This document describes the steps and presents their status.

Content:
*   [Introduction](#introduction);
*   [Basic description and limitations](#basic-description-and-limitations);
*   [The interface](#the-interface);
*   [Null implementation](#null-implementation);
*   [Generic implementation](#generic-implementation);
*   [The lookup process](#the-lookup-process);
*   [Compactified implementations](#compactified-implementations);
*   [Fossilization](#fossilization).


## Introduction

I arrived to idea of the project long ago. Being a web services developer
usually means using some very high-leveled tools, abstracted layers over
abstracted layers. Working for years with mastodons like Ruby on Rails 
framework, one can totally forget about Rack layer, server layer, HTTP,
sockets, TCP etc. But sometimes the cruel fate smiles at you and gives you an 
awesome problem to solve, and that expands your vision. In my case it was a 
task to create highly efficient and stupidly simple (as in KISS principle) file 
service, which would take a file via POST request and store it, or give a file 
via GET request. Highly efficient meant that the service should have had low 
memory footprint, low CPU consumption, ability to work with tens of thousands 
of connections simultaneously, and be patient with slow clients. No wonder the 
service could barely be implemented in Ruby. I chose Golang for the task.

The language provided some nice abilities. Compiled binaries were blazing fast, 
goroutines were neat, resource consumption was low, and standard libary was
just awesome. It was a fresh air after a language, where method invocation is
basically a lookup in hash table, green thread could use up to 1 megabytes of
memory just after creation, and standard HTTP server is no more than a brick.
Nevertheless, hunger for optimization, if it has appeared, can rarely be sated.
The main question was the following: Do I really need something as 
[complex][go-net-http], as [HTTP-compliant][go-net-http-server-990], with tens
of fields in [request][go-net-http-request-103] just for _stupidly simple_
service? I think, the last straw which broke the back of camel, was the
processing of HTTP headers.

Oh, the HTTP headers! I'm terrified every time I look at _what_ my browser
sends to servers! Trillions of some mindless `key: value` thingies! So, just
imagine, even _stupidly simple_ service should docily accept any blabbering
sent by mad clients. I had had a growing suspicion that in some cases (say,
small pics for user avatars), the incoming requests have more bytes in headers
than in bodies. It is a disaster. And how the disaster is processed? The
processing just happily [puts everything][go-net-textproto-reader-483] into a
hash table. Into a hash table of string slices, no less. Its cousins do some
weird voodoo, too. Look, they even [roll][puma-http-79] over some three dozens 
of so-called standard headers for _every_ incoming key. Of course, just to put 
the result into hash table, again. And the processing can _not_ be changed, 
tweaked, modified.

What would be proper headers processing for the _stupdily simple_ service, in 
my opinion? Start with feeding bytes of key into some lookup structure til we
find colon. If the lookup tells us it's something useful (for example,
`Content-Length`), get bytes of its value and parse them. If it's not useful,
discard the bytes of its value til CRLF. Actually, as we feed the bytes, the 
lookup can tell us immediately if the key is unsupported, so we can discard the 
rest til the  colon. Of course, the lookup structure should keep, what and how 
many bytes we've already fed it. Basically, the bytes are just a prefix of the 
supported keys, so the underlying data could be represented just as a compact 
prefix tree or a [radix tree][wiki-radix-tree]:

```
       "content-"
     /     |     \
"length" "type" "disposition"

```

The lookup process can also be represented as feeding bytes to some finite
state automaton, but the tree allows to greatly compactify the representation.
Of course, the tree can have may representations too. That's why the project is
aimed to support more than one implementation (at least three), which are
united by common interface.

[go-net-http]: https://github.com/golang/go/tree/go1.19/src/net/http
[go-net-http-server-990]: https://github.com/golang/go/blob/go1.19/src/net/http/server.go#L990
[go-net-http-request-103]: https://github.com/golang/go/blob/go1.19/src/net/http/request.go#L103-L324
[go-net-textproto-reader-483]: https://github.com/golang/go/blob/go1.19/src/net/textproto/reader.go#L483
[puma-http-79]: https://github.com/puma/puma/blob/v5.6.4/ext/puma_http11/puma_http11.c#L79-L161
[wiki-radix-tree]: https://en.wikipedia.org/wiki/Radix_tree


## Basic description and limitations

Before the description of the interface of radix tree, some limitations should
be discussed.

1.  The trees would be static, read-only. That means, that they would be
    constructed by a kind of factory method from a strings list, but after
    that, no new strings would be added, no added strings would be deleted.
    Indeed, the HTTP headers supported by a _stupidly simple_ service from
    introduction's example are for sure predefined. That also would allow to 
    reach the goal of possible tree _fossilization_ in Golang code, as no 
    dynamic tree could be fossilized.

2.  There would be no direct access to node structures of a tree, so the tree
    interface would be the only interface to access any tree or node data. That 
    means, that a tree is not just a storage for nodes, but it is the whole and 
    the only context.

3.  Nevertheless, a tree's node still should have a name. As any tree can have
    only limited amount of nodes, an integer "name" would be enough. As those
    indices could have some inner meaning for a tree implementation, the tree
    interface should provide methods to extract the amount of nodes, extract
    index of root node (if the tree is not empty), check if it has a node with 
    the specific index. Integer numbers, which are not indices of nodes of the
    tree in question, would be called non-node indices.

4.  Tree's nodes would need the following data. First, there should be an
    integer mark with the following interpretation: if the mark is negative,
    then the node doesn't point to a string in the original list, and if the
    mark is non-negative, then it is the index of a string in the list. Second,
    there should be read-only access to the string (if the mark is negative, 
    the string returned would be empty string). Third, the same access would be
    granted to the node's prefix. Fourth, there should be read-only access to
    node's children, which would be presented via a kind of iterator.

5.  There would be no limitation on amount of nodes with empty prefix.
    Actually, the whole tree could consist only of nodes with empty prefix. 

6.  First bytes of non-empty children's prefixes should be unique for every 
    parent node. (That would allow correct lookup process.)

7.  To simplify lookup process, a method of _transition_ would be required. The
    method would accept a node index, a cursor over the node's prefix, and a
    byte, and return the node index if the prefix's byte over cursor coincides
    with the provided byte, otherwise index of a child (of grandchild, 
    grandgrandchild etc in case of their empty prefixes), if first byte of the 
    child's prefix coincides with the provided byte, otherwise a non-node
    index. (The lookup process could actually be implemented without the
    transition method, as there would be prefix extraction and children
    extraction methods, but absence of limitations on amount of nodes with
    empty prefix would made the process much more complex. As the tree has more 
    information of its inner structure, it can implement the method in an
    efficient way.)


## The interface

The following Golang interface represents a radix tree according to the section
above:
```golang
type Tree interface {
	Size() int
	Has(n int) bool
	Root() int
	NodeMark(n int) int
	NodeString(n int) string
	NodePref(n int) string
	NodeEachChild(n int, e func(int) bool) 
	NodeTransit(n, npos int, b byte) int
}
```

Status: realized in source file


## Null implementation

The implementation would be required, for example, in case of providing null 
value of Tree interface above. It would have no nodes at all and would be just 
empty struct. Obviously, only one empty instance of null implementation would 
be required, so there would be no factory.

Status: realized in source file


## Generic implementation

The implementation would cover all sane cases. The struct's definitions would
be roughly as the following:
```go tree
type tree struct {
  strings []string
  nodes   []node
}
```
```go node
type node struct {
  pref     string
  mark     int
  children []int
}
```

Valid node indices in non-empty tree would cover all integer numbers from zero
(root index) to (Size() - 1) number. A tree would be created via sequential
insertion of original strings. There could be the following cases on every
insertion.

1.  The tree is empty. In this case, the insertion is just creation of root:
    ```
               nil
                |
                V
     insert("authorization")
                |
                V
    "authorization", mark: 0
    ```

2.  The inserted string coincides with a string of a node, which is already 
    present in the tree. The insertion then does nothing:
    ```
    "authorization", mark: 0
                |
                V
     insert("authorization")
                |
                V
    "authorization", mark: 0
    ```

3.  The inserted string coincides with a string of a node, which is absent in 
    the tree. The insertion then sets the node's mark:
    ```
                   "auth", mark: -1
                  /               \
    "orization", mark: 0   "entication", mark: 1
                          |
                          V
                    insert("auth")
                          |
                          V
                   "auth", mark: 2
                  /               \
    "orization", mark: 0   "entication", mark: 1
    ```

4.  The inserted string doesn't coincide with any string, but it is a prefix
    for string of a node. There can be many nodes like that, but the insertion 
    should find the one closest to the root and split it, creating a child node 
    which inherits children and mark, setting the corresponding mark to the 
    node after:
    ```
         "authori", mark: -1
        /                   \
    "zation, mark: 0      "ty", mark: 1
                   |
                   V
             insert("auth")
                   |
                   V
           "auth", mark: 2
                   |
           "ori", mark: -1
          /               \
    "zation, mark: 0      "ty", mark: 1
    ```

5.  The inserted string doesn't coincide with any string, and it is not a
    prefix for a node, but there is string of a node which is prefix for the
    string. Again, there can be many nodes like that, but the insertion should
    find the most far one and add a new leaf for it with corresponding mark and
    the suffix:
    ```
                   "auth", mark: 0
                   /
    "orization", mark: 1
                          |
                          V
               insert("authentication")
                          |
                          V
                   "auth", mark: 0
                   /               \
    "orization", mark: 1   "entication", mark: 2
    ```

6.  The inserted string doesn't coincide with any string, and it is not a
    prefix for a node, and there is no node with prefix for the string. The
    insertion then should find the longest common prefix (which can be empty
    and, that's why, always exists), split the node accordingly as in 4th case
    and add new leaf as in 5th case:
    ```
                             "auth", mark: 0
                            /               \
              "orization", mark: 1   "entication", mark: 2
                                    |
                                    V
                           insert("authority")
                                    |
                                    V
                             "auth", mark: 0
                            /               \
                   "ori", mark: -1   "entication", mark: 2
                   /            \
    "zation", mark: 1        "ty", mark: 3
    ```

It's pretty obvious that insertion process uses heavily the lookup process.

Status: realized in source file


## The lookup process

The process would start with supplied tree, take byte by byte and return, if it 
has found corresponding string in the tree. Its algorithm is as following.
1.  Start with root node index as current node index and zero position.
2.  Take a byte, try to transit from the current node.
3.  If the node index returned by transition method is non-node index for the
    tree, the process should stop, reporting that it has found nothing.
4.  Otherwise, if the node index is the same, then increment position.
5.  Otherwise, set position to 1, set the returned child node index as current.
6.  Repeat from step 2.
7.  User of the process can check any time, if the string is found, looking at
    the current node, its mark, and if position is more or equal to length of
    node prefix.

Status: yet to appear in source file


## Compactified implementations

The generic implementation has pretty high memory consumption for the cases of
short length strings or trees with high amount of leaves, as for 64-bit 
environments every string has 16 bytes in header and every slice has 24 bytes 
in one. For example, the following tree
```
                               ""
          /                  /  |     \       \      \
       "P"          "OPTIONS" "DELETE" "TRACE" "HEAD" "GET"
     /  |  \
"OST"  "UT" "ATCH"
```
could represent a list of HTTP methods. The generic implementation would use as
much as
```
24 + 24 +                               // tree struct
9 * 16 +                                // 9 string headers
11 * (16 + 8 + 24) +                    // 11 node structs
4 + 3 + 5 + 7 + 6 + 5 + 4 + 3 +         // lengths of original strings
0 + 1 + 3 + 2 + 4 + 7 + 6 + 5 + 4 + 3 + // lengths of prefixes
9 * 8 =                                 // 9 leaves in all children slices
864                                     // bytes altogether
```
and more than 83% of the memory consumption is the headers. The following 
techniques could be of use to compactify representation.

1.  The list of strings could be discarded, the node strings could be recovered
    from their prefixes.
2.  All prefixes could be jammed into one large string, and data on prefix
    could be put into start/length pair of bytes.
3.  Mark field could be of int8 type.
4.  The leaves could have their own struct without children slices. Leaves and
    non-leaves could be distincted, for example, by having negative and
    non-negative indices.
5.  Non-leaves could use arrays instead of slices for children indices.

Status: needs research, yet to appear in source file


## Fossilization of a radix tree

As the lookup process of a string in a radix tree can be formulated in terms of
finite state automaton, it can also be hard-coded into code base. Golang 
provides some tools and means to _generate_ source files, and all those marks,
prefixes, and going over children can be hard-coded directly into functions.
The lookup process would still need a state, which is just a current processing 
function, and a counter of processed bytes for the state.

Status: needs research, study of the Golang generation, yet to appear in source
file
