// Separate module so resolving golang.org/x/mod never touches
// golang.org/x/net's own go.mod / go.sum, which must ship unmodified.
module example.com/modzip

go 1.24.0

require golang.org/x/mod v0.28.0
