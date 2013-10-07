package gangliamr

import (
	"github.com/daaku/go.ganglia/gmetric"
	"github.com/daaku/go.metrics"
)

// Timers capture the duration and rate of events.
type Timer struct {
	// Unless explicitly specified, this will be a timer with a standard
	// histogram and meter. The histogram will use an exponentially-decaying
	// sample with the same reservoir size and alpha as UNIX load averages.
	metrics.Timer

	Name        string // Required
	Title       string
	Description string
	Groups      []string
	histime     *histime
}

func (t *Timer) writeValue(c *gmetric.Client) {
	t.histime.writeValue(c)
}

func (t *Timer) writeMeta(c *gmetric.Client) {
	t.histime.writeMeta(c)
}

func (t *Timer) register(r *Registry) {
	if t.Timer == nil {
		t.Timer = metrics.NewTimer()
	}
	t.histime = &histime{
		histimeMetric: t,
		Name:          t.Name,
		Title:         t.Title,
		Units:         "nanoseconds",
		Description:   t.Description,
		Groups:        t.Groups,
	}
	t.histime.register(r)
}
