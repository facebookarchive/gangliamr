package gangliamr

import (
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
	registry.write()
	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".one-minute",
		Value: "8.6",
		Unit:  "count/sec",
		Slope: "both",
	})

	const v2 = 42
	meter.Mark(v2)
	registry.tick()
	registry.write()
	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".one-minute",
		Value: "8.584008882925865",
		Unit:  "count/sec",
		Slope: "both",
	})
}
