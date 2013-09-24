package gangliamr

import (
	"github.com/daaku/go.ganglia/gmetric"
	"github.com/daaku/go.metrics"
)

type Counter struct {
	// The underlying in-memory metric.
	metrics.Counter

	// The name is used as the file name, and also the title unless one is
	// explicitly provided.
	Name string

	// The title is for human consumption and is shown atop the graph.
	Title string

	// The units are shown in the graph to provide context to the numbers.
	// Default is "count".
	Units string

	// Descriptions serve as documentation.
	Description string

	// The groups ensure your metric is kept alongside sibling metrics.
	Groups []string

	gmetric gmetric.Metric
}

func (c *Counter) writeMeta(client *gmetric.Client) {
	client.WriteMeta(&c.gmetric)
}

func (c *Counter) writeValue(client *gmetric.Client) {
	client.WriteValue(&c.gmetric, c.Count())
}

func (c *Counter) register(r *Registry) {
	c.Counter = metrics.NewCounter()
	c.gmetric = gmetric.Metric{
		Name:        r.makeName(c.Name),
		Title:       c.Title,
		Units:       nonEmpty(c.Units, "count"),
		Description: c.Description,
		Groups:      c.Groups,
		ValueType:   gmetric.ValueUint32,
		Slope:       gmetric.SlopeBoth,
	}
}
