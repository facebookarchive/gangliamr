// Package gangliamr provides metrics backed by Ganglia.
//
// The underlying in-memory metrics are used from:
// http://godoc.org/github.com/daaku/go.metrics. Application code should use
// the interfaces defined in that package in order to not be Ganglia specific.
//
// The underlying Ganglia library is:
// http://godoc.org/github.com/daaku/go.ganglia/gmetric.
package gangliamr
