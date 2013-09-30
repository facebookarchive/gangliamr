package gangliamr

import (
	"github.com/daaku/go.ganglia/gmetric"
	"github.com/daaku/go.metrics"
)

// Timers capture the duration and rate of events.
type Timer struct {
	// The underlying in-memory metric. Unless explicitly specified, this will be
	// a timer with a standard histogram and meter. The histogram will use an
	// exponentially-decaying sample with the same reservoir size and alpha as
	// UNIX load averages.
	metrics.Timer

	Name        string // Required
	Title       string
	Description string
	Groups      []string
	histime     *histime
	calls       *meterShared
}

func (t *Timer) writeValue(c *gmetric.Client) {
	t.histime.writeValue(c)
	t.calls.writeValue(c)
}

func (t *Timer) writeMeta(c *gmetric.Client) {
	t.histime.writeMeta(c)
	t.calls.writeMeta(c)
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
	t.calls = &meterShared{
		meterMetric: t,
	}
	t.calls = &meterShared{
		meterMetric: t,
		Name:        t.Name + r.nameSeparator() + "calls",
		Title:       makeOptional(t.Title, "calls"),
		Units:       "nanoseconds",
		Description: makeOptional(t.Description, "calls"),
		Groups:      t.Groups,
	}
	t.calls.register(r)
}
