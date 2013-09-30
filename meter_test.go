package gangliamr_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/daaku/go.ganglia/gmon"
	"github.com/daaku/go.ganglia/gmondtest"
	"github.com/daaku/go.gangliamr"
	"github.com/daaku/go.metrics"
)

func TestMeterSimple(t *testing.T) {
	t.Parallel()
	h := gmondtest.NewHarness(t)
	defer h.Stop()

	const name = "meter_simple_metric"
	var meter metrics.Meter
	meter = &gangliamr.Meter{
		Name: name,
	}

	registry := gangliamr.Registry{
		Client:            h.Client,
		WriteTickDuration: 5 * time.Millisecond,
	}
	registry.Register(meter)

	const v1 = 43
	meter.Mark(v1)
	meter.Tick()
	if meter.Count() != v1 {
		t.Fatalf("was expecting %d got %d", v1, meter.Count())
	}

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
	meter.Tick()
	if meter.Count() != v1+v2 {
		t.Fatalf("was expecting %d got %d", v1+v2, meter.Count())
	}

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
