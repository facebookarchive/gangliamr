package gangliamr

import (
	"fmt"

	"github.com/daaku/go.ganglia/gmetric"
	"github.com/daaku/go.metrics"
)

type Histogram struct {
	// The underlying in-memory metric. This must be specified.
	metrics.Histogram

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

	count   gmetric.Metric
	histime *histime
}

func (h *Histogram) writeValue(c *gmetric.Client) {
	h.histime.writeValue(c)
	c.WriteValue(&h.count, h.Count())
}

func (h *Histogram) writeMeta(c *gmetric.Client) {
	h.histime.writeMeta(c)
	c.WriteMeta(&h.count)
}

func (h *Histogram) register(r *Registry) {
	if h.Histogram == nil {
		panic(fmt.Sprintf("histogram misconfigured: %+v", h))
	}
	h.histime = &histime{
		histimeMetric: h,
		Name:          h.Name,
		Title:         h.Title,
		Units:         h.Units,
		Description:   h.Description,
		Groups:        h.Groups,
	}
	h.histime.register(r)
	h.count = gmetric.Metric{
		Name:      r.makeName(h.Name + "count"),
		ValueType: gmetric.ValueUint32,
		Slope:     gmetric.SlopeBoth,
	}
}
