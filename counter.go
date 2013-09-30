package gangliamr

import (
	"github.com/daaku/go.ganglia/gmetric"
	"github.com/daaku/go.metrics"
)

// Counters hold an int64 value that can be incremented and decremented.
type Counter struct {
	metrics.Counter
	Name        string // Required.
	Title       string
	Units       string // Default is "count".
	Description string
	Groups      []string
	gmetric     gmetric.Metric
}

func (c *Counter) writeMeta(client *gmetric.Client) {
	client.WriteMeta(&c.gmetric)
}

func (c *Counter) writeValue(client *gmetric.Client) {
	client.WriteValue(&c.gmetric, c.Count())
}

func (c *Counter) register(r *Registry) {
	if c.Counter == nil {
		c.Counter = metrics.NewCounter()
	}
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
