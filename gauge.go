package gangliamr

import (
	"github.com/daaku/go.ganglia/gmetric"
	"github.com/daaku/go.metrics"
)

type Gauge struct {
	// The underlying in-memory metric.
	metrics.Gauge

	// The name is used as the file name, and also the title unless one is
	// explicitly provided.
	Name string

	// The title is for human consumption and is shown atop the graph.
	Title string

	// The units are shown in the graph to provide context to the numbers.
	// Default is "value".
	Units string

	// Descriptions serve as documentation.
	Description string

	// The groups ensure your metric is kept alongside sibling metrics.
	Groups []string

	gmetric gmetric.Metric
}

func (g *Gauge) writeMeta(c *gmetric.Client) {
	c.WriteMeta(&g.gmetric)
}

func (g *Gauge) writeValue(c *gmetric.Client) {
	c.WriteValue(&g.gmetric, g.Value())
}

func (g *Gauge) register(r *Registry) {
	g.Gauge = metrics.NewGauge()
	g.gmetric = gmetric.Metric{
		Name:        r.makeName(g.Name),
		Title:       g.Title,
		Units:       nonEmpty(g.Units, "value"),
		Description: g.Description,
		Groups:      g.Groups,
		ValueType:   gmetric.ValueUint32,
		Slope:       gmetric.SlopeBoth,
	}
}
