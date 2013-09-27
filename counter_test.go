package gangliamr_test

import (
	"testing"
	"time"

	"github.com/daaku/go.ganglia/gmon"
	"github.com/daaku/go.ganglia/gmondtest"
	"github.com/daaku/go.gangliamr"
	"github.com/daaku/go.metrics"
)

func TestCounterSimple(t *testing.T) {
	t.Parallel()
	h := gmondtest.NewHarness(t)
	defer h.Stop()

	const name = "counter_simple_metric"
	var counter metrics.Counter
	counter = &gangliamr.Counter{
		Name: name,
	}

	registry := gangliamr.Registry{
		Client:       h.Client,
		TickDuration: 5 * time.Millisecond,
	}
	registry.Register(counter)

	counter.Inc(1)

	if counter.Count() != 1 {
		t.Fatalf("was expecting 1 got %d", counter.Count())
	}

	h.ContainsMetric(&gmon.Metric{
		Name:  name,
		Value: "1",
		Unit:  "count",
		Slope: "both",
	})

	counter.Inc(10)

	if counter.Count() != 11 {
		t.Fatalf("was expecting 11 got %d", counter.Count())
	}

	h.ContainsMetric(&gmon.Metric{
		Name:  name,
		Value: "11",
		Unit:  "count",
		Slope: "both",
	})
}
