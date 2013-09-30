// Package gangliamr provides metrics backed by Ganglia.
//
// The underlying in-memory metrics are used from:
// http://godoc.org/github.com/daaku/go.metrics. Application code should use
// the interfaces defined in that package in order to not be Ganglia specific.
//
// The underlying Ganglia library is:
// http://godoc.org/github.com/daaku/go.ganglia/gmetric.
//
// A handful of metrics types are provided, and they all have a similar form.
// The "name" property is always required, all other metadata properties are
// optional. The metric instances are also automatically created upon
// registration. The only exception is that the Histogram metric must be
// explicitly provided as it required user configuration.
//
// The common set of properties for the metrics are:
//
//     // The name is used as the file name, and also the title unless one is
//     // explicitly provided. This property is required.
//     Name string
//
//     // The title is for human consumption and is shown atop the graph.
//     Title string
//
//     // The units are shown in the graph to provide context to the numbers.
//     // The default value varies based on the metric type.
//     Units string
//
//     // Descriptions serve as documentation.
//     Description string
//
//     // The groups ensure your metric is kept alongside sibling metrics.
//     Groups []string
package gangliamr
