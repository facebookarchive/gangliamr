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

func TestGaugeSimple(t *testing.T) {
	t.Parallel()
	h := gmondtest.NewHarness(t)
	defer h.Stop()

	const name = "gauge_simple_metric"
	var gauge metrics.Gauge
	gauge = &gangliamr.Gauge{
		Name: name,
	}

	registry := gangliamr.Registry{
		Client:            h.Client,
		WriteTickDuration: 5 * time.Millisecond,
	}
	registry.Register(gauge)

	const v1 = 42
	gauge.Update(v1)
	if gauge.Value() != v1 {
		t.Fatalf("was expecting %d got %d", v1, gauge.Value())
	}

	h.ContainsMetric(&gmon.Metric{
		Name:  name,
		Value: fmt.Sprint(v1),
		Unit:  "value",
		Slope: "both",
	})

	const v2 = 43
	gauge.Update(v2)
	if gauge.Value() != v2 {
		t.Fatalf("was expecting %d got %d", v2, gauge.Value())
	}

	h.ContainsMetric(&gmon.Metric{
		Name:  name,
		Value: fmt.Sprint(v2),
		Unit:  "value",
		Slope: "both",
	})
}
