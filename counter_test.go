package gangliamr

import (
	"testing"

	"github.com/facebookgo/ganglia/gmon"
	"github.com/facebookgo/ganglia/gmondtest"
	"github.com/facebookgo/metrics"
)

func TestCounterSimple(t *testing.T) {
	t.Parallel()
	h := gmondtest.NewHarness(t)
	defer h.Stop()

	const name = "counter_simple_metric"
	var counter metrics.Counter
	counter = &Counter{
		Name: name,
	}

	registry := testRegistry(h.Client)
	registry.Register(counter)

	counter.Inc(1)
	if counter.Count() != 1 {
		t.Fatalf("was expecting 1 got %d", counter.Count())
	}

	registry.write()
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

	registry.write()
	h.ContainsMetric(&gmon.Metric{
		Name:  name,
		Value: "11",
		Unit:  "count",
		Slope: "both",
	})
}
