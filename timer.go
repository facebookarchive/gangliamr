package gangliamr

import (
	"time"

	"github.com/daaku/go.ganglia/gmetric"
	"github.com/daaku/go.metrics"
)

var timerTrackedPercentiles = []float64{0.5, 0.75, 0.95, 0.98, 0.99, 0.999}

// Timers capture the duration and rate of events.
type Timer struct {
	// Unless explicitly specified, this will be a timer with a standard
	// histogram and meter. The histogram will use an exponentially-decaying
	// sample with the same reservoir size and alpha as UNIX load averages.
	metrics.Timer

	// Defaults to time.Second, supports time.Microsecond, time.Millisecond
	// or time.Second. Other values will trigger falling back to time.Second.
	Resolution time.Duration

	Name        string // Required
	Title       string
	Description string
	Groups      []string

	max    gmetric.Metric
	mean   gmetric.Metric
	min    gmetric.Metric
	stddev gmetric.Metric
	p50    gmetric.Metric
	p75    gmetric.Metric
	p95    gmetric.Metric
	p98    gmetric.Metric
	p99    gmetric.Metric
	p999   gmetric.Metric
}

func (t *Timer) writeValue(c *gmetric.Client) {
	ps := t.Percentiles(timerTrackedPercentiles)
	c.WriteValue(&t.max, t.normalizeInt64(t.Max()))
	c.WriteValue(&t.mean, t.normalizeFloat64(t.Mean()))
	c.WriteValue(&t.min, t.normalizeInt64(t.Min()))
	c.WriteValue(&t.stddev, t.normalizeFloat64(t.StdDev()))
	c.WriteValue(&t.p50, t.normalizeFloat64(ps[0]))
	c.WriteValue(&t.p75, t.normalizeFloat64(ps[1]))
	c.WriteValue(&t.p95, t.normalizeFloat64(ps[2]))
	c.WriteValue(&t.p98, t.normalizeFloat64(ps[3]))
	c.WriteValue(&t.p99, t.normalizeFloat64(ps[4]))
	c.WriteValue(&t.p999, t.normalizeFloat64(ps[5]))
}

func (t *Timer) writeMeta(c *gmetric.Client) {
	c.WriteMeta(&t.max)
	c.WriteMeta(&t.mean)
	c.WriteMeta(&t.min)
	c.WriteMeta(&t.stddev)
	c.WriteMeta(&t.p50)
	c.WriteMeta(&t.p75)
	c.WriteMeta(&t.p95)
	c.WriteMeta(&t.p98)
	c.WriteMeta(&t.p99)
	c.WriteMeta(&t.p999)
}

func (t *Timer) register(r *Registry) {
	if t.Timer == nil {
		t.Timer = metrics.NewTimer()
	}
	t.max = gmetric.Metric{
		Name:        r.makeName(t.Name, "max"),
		Title:       makeOptional(t.Title, "maximum"),
		Description: makeOptional(t.Description, "maximum"),
		Units:       t.units(),
		Groups:      t.Groups,
		ValueType:   gmetric.ValueUint32,
		Slope:       gmetric.SlopeBoth,
	}
	t.mean = gmetric.Metric{
		Name:        r.makeName(t.Name, "mean"),
		Title:       makeOptional(t.Title, "mean"),
		Description: makeOptional(t.Description, "mean"),
		Units:       t.units(),
		Groups:      t.Groups,
		ValueType:   gmetric.ValueUint32,
		Slope:       gmetric.SlopeBoth,
	}
	t.min = gmetric.Metric{
		Name:        r.makeName(t.Name, "min"),
		Title:       makeOptional(t.Title, "minimum"),
		Description: makeOptional(t.Description, "minimum"),
		Units:       t.units(),
		Groups:      t.Groups,
		ValueType:   gmetric.ValueUint32,
		Slope:       gmetric.SlopeBoth,
	}
	t.stddev = gmetric.Metric{
		Name:        r.makeName(t.Name, "stddev"),
		Title:       makeOptional(t.Title, "standard deviation"),
		Description: makeOptional(t.Description, "standard deviation"),
		Units:       t.units(),
		Groups:      t.Groups,
		ValueType:   gmetric.ValueUint32,
		Slope:       gmetric.SlopeBoth,
	}
	t.p50 = gmetric.Metric{
		Name:        r.makeName(t.Name, "p50"),
		Title:       makeOptional(t.Title, "50th percentile"),
		Description: makeOptional(t.Description, "50th percentile"),
		Units:       t.units(),
		Groups:      t.Groups,
		ValueType:   gmetric.ValueFloat64,
		Slope:       gmetric.SlopeBoth,
	}
	t.p75 = gmetric.Metric{
		Name:        r.makeName(t.Name, "p75"),
		Title:       makeOptional(t.Title, "75th percentile"),
		Description: makeOptional(t.Description, "75th percentile"),
		Units:       t.units(),
		Groups:      t.Groups,
		ValueType:   gmetric.ValueFloat64,
		Slope:       gmetric.SlopeBoth,
	}
	t.p95 = gmetric.Metric{
		Name:        r.makeName(t.Name, "p95"),
		Title:       makeOptional(t.Title, "95th percentile"),
		Description: makeOptional(t.Description, "95th percentile"),
		Units:       t.units(),
		Groups:      t.Groups,
		ValueType:   gmetric.ValueFloat64,
		Slope:       gmetric.SlopeBoth,
	}
	t.p98 = gmetric.Metric{
		Name:        r.makeName(t.Name, "p98"),
		Title:       makeOptional(t.Title, "98th percentile"),
		Description: makeOptional(t.Description, "98th percentile"),
		Units:       t.units(),
		Groups:      t.Groups,
		ValueType:   gmetric.ValueFloat64,
		Slope:       gmetric.SlopeBoth,
	}
	t.p99 = gmetric.Metric{
		Name:        r.makeName(t.Name, "p99"),
		Title:       makeOptional(t.Title, "99th percentile"),
		Description: makeOptional(t.Description, "99th percentile"),
		Units:       t.units(),
		Groups:      t.Groups,
		ValueType:   gmetric.ValueFloat64,
		Slope:       gmetric.SlopeBoth,
	}
	t.p999 = gmetric.Metric{
		Name:      r.makeName(t.Name, "p999"),
		Title:     makeOptional(t.Title, "99.9th percentile"),
		Units:     t.units(),
		Groups:    t.Groups,
		ValueType: gmetric.ValueFloat64,
		Slope:     gmetric.SlopeBoth,
	}
}

func (t *Timer) normalizeInt64(v int64) int64 {
	switch t.Resolution {
	default:
		return v / int64(time.Second)
	case time.Microsecond:
		return v / int64(time.Microsecond)
	case time.Millisecond:
		return v / int64(time.Millisecond)
	case time.Nanosecond:
		return v
	}
}

func (t *Timer) normalizeFloat64(v float64) float64 {
	switch t.Resolution {
	default:
		return v / float64(time.Second)
	case time.Microsecond:
		return v / float64(time.Microsecond)
	case time.Millisecond:
		return v / float64(time.Millisecond)
	case time.Nanosecond:
		return v
	}
}

func (t *Timer) units() string {
	switch t.Resolution {
	default:
		return "seconds"
	case time.Microsecond:
		return "microseconds"
	case time.Millisecond:
		return "milliseconds"
	case time.Nanosecond:
		return "nanoseconds"
	}
}
