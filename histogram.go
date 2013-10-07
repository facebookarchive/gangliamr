package gangliamr

import (
	"github.com/daaku/go.ganglia/gmetric"
	"github.com/daaku/go.metrics"
)

// Histograms calculate distribution statistics from an int64 value.
type Histogram struct {
	// Unless explicitly specified, this will be a histogram with an
	// exponentially-decaying sample with the same reservoir size and alpha as
	// UNIX load averages.
	metrics.Histogram

	Name        string // Required.
	Title       string
	Units       string // Default is "value".
	Description string
	Groups      []string
	histime     *histime
}

func (h *Histogram) writeValue(c *gmetric.Client) {
	h.histime.writeValue(c)
}

func (h *Histogram) writeMeta(c *gmetric.Client) {
	h.histime.writeMeta(c)
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
}
