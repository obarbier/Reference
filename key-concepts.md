## Terminology

This is used as a top level document to review some key software engineering concept

TODO: find a way to make this an index page

* Hermetic: A Hermetic Build is a release engineering best practice for increasing the reliability and consistency of software builds. They are self-contained, and do not depend on anything outside of the build environment. This means they do not have network access, and cannot fetch dependencies at runtime. [blog](https://testing.googleblog.com/2012/10/hermetic-servers.html)

* Radix tree: A radix tree (also radix trie or compact prefix tree or compressed trie) is a data structure that represents a space-optimized trie (prefix tree) in which each node that is the only child is merged with its parent. The result is that the number of children of every internal node is at most the radix r of the radix tree, where r is a positive integer and a power x of 2, having x ≥ 1. Unlike regular trees, edges can be labeled with sequences of elements as well as single elements. This makes radix trees much more efficient for small sets (especially if the strings are long) and for sets of strings that share long prefixes.  [wiki](https://en.wikipedia.org/wiki/Radix_tree)