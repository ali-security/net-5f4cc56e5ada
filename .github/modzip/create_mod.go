// Command create_mod packs a module source zip byte-compatible with the one
// proxy.golang.org serves, using Go's own golang.org/x/mod/zip implementation so
// the inclusion rules (drops .git/, vendor/, nested modules) match by
// construction.
//
// It lives under .github/ for two reasons: the go tool ignores directories whose
// name begins with a dot, so `go build ./...` and `go test ./...` at the module
// root never see it; and the build workflow strips .github/ from the tree it
// packs, so neither this helper nor the workflow itself can leak into the
// published artifact.
//
// Usage: create_mod <module-path> <version> <source-dir> <output-zip>
package main

import (
	"log"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatal("usage: create_mod <module-path> <version> <source-dir> <output-zip>")
	}
	m := module.Version{Path: os.Args[1], Version: os.Args[2]}
	f, err := os.Create(os.Args[4])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := zip.CreateFromDir(f, m, os.Args[3]); err != nil {
		log.Fatal(err)
	}
	log.Printf("created module zip: %s", os.Args[4])
}
