package main

// The purpose of any package system is to make the design and maintenance of large
// programs practical by grouping related features together into units that can be easily
// understood and changed, independent of the other packages of the program. This
// modularity allows packages to be shared and reused by different projects, distributed
// within an organization, or made available to the wider world.

// Packages also provide encapsulation by controlling which names are visible or exported outside the package

//? Go compilation is notably faster than most other compiled languages, even when building from scratch.
// There are three main reasons for the compiler’s speed.
//
//! First, all imports must be explicitly listed at the beginning of each source file,
//! so the compiler does not have to read and process an entire file to determine its dependencies.
//
//! Second, the dependencies of a package form a directed acyclic graph, and because there are no cycles,
//! packages can be compiled separately and perhaps in parallel.
//
//! Finally, the object file for a compiled Go package
//! records export information not just for the package itself, but for its dependencies too
// When compiling a package, the compiler must read one object file for each import
// but need not look beyond these files
