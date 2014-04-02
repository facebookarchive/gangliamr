package gangliamr

import (
	"testing"

	"github.com/facebookgo/ganglia/gmon"
	"github.com/facebookgo/ganglia/gmondtest"
	"github.com/facebookgo/metrics"
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

	const units = "count/sec"
	const v1 = 43
	meter.Mark(v1)
	registry.tick()
	registry.write()
	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".one-minute",
		Value: "8.6",
		Unit:  units,
		Slope: "both",
	})

	const v2 = 42
	meter.Mark(v2)
	registry.tick()
	registry.write()
	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".one-minute",
		Value: "8.584008882925865",
		Unit:  units,
		Slope: "both",
	})
}
