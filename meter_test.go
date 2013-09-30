package gangliamr

import (
	"fmt"
	"testing"

	"github.com/daaku/go.ganglia/gmon"
	"github.com/daaku/go.ganglia/gmondtest"
	"github.com/daaku/go.metrics"
)

func TestMeterSimple(t *testing.T) {
	t.Parallel()
	h := gmondtest.NewHarness(t)
	defer h.Stop()

	const name = "meter_simple_metric"
	var meter metrics.Meter
	meter = &Meter{
		Name: name,
	}

	registry := testRegistry(h.Client)
	registry.Register(meter)

	const v1 = 43
	meter.Mark(v1)
	registry.tick()
	if meter.Count() != v1 {
		t.Fatalf("was expecting %d got %d", v1, meter.Count())
	}

	registry.write()
	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".count",
		Value: fmt.Sprint(meter.Count()),
		Unit:  "count",
		Slope: "both",
	})

	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".one-minute",
		Value: "8.6",
		Unit:  "count",
		Slope: "both",
	})

	const v2 = 42
	meter.Mark(v2)
	registry.tick()
	if meter.Count() != v1+v2 {
		t.Fatalf("was expecting %d got %d", v1+v2, meter.Count())
	}

	registry.write()
	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".count",
		Value: fmt.Sprint(meter.Count()),
		Unit:  "count",
		Slope: "both",
	})

	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".one-minute",
		Value: "8.584008882925865",
		Unit:  "count",
		Slope: "both",
	})
}
