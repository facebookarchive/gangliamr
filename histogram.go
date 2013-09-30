package gangliamr

import (
	"github.com/daaku/go.ganglia/gmetric"
	"github.com/daaku/go.metrics"
)

// Histograms calculate distribution statistics from an int64 value.
type Histogram struct {
	// The underlying in-memory metric. Unless explicitly specified, this will be
	// histogram with an exponentially-decaying sample with the same reservoir
	// size and alpha as UNIX load averages.
	metrics.Histogram

	Name        string // Required.
	Title       string
	Units       string // Default is "value".
	Description string
	Groups      []string
	count       gmetric.Metric
	histime     *histime
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
		h.Histogram = metrics.NewHistogram(metrics.NewExpDecaySample(1028, 0.015))
	}
	h.histime = &histime{
		histimeMetric: h,
		Name:          h.Name,
		Title:         h.Title,
		Units:         nonEmpty(h.Units, "value"),
		Description:   h.Description,
		Groups:        h.Groups,
	}
	h.histime.register(r)
	h.count = gmetric.Metric{
		Name:        r.makeName(h.Name, "count"),
		Title:       makeOptional(h.Title, "count"),
		Description: makeOptional(h.Description, "count"),
		Groups:      h.Groups,
		Units:       "count",
		ValueType:   gmetric.ValueUint32,
		Slope:       gmetric.SlopeBoth,
	}
}
