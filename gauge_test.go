package gangliamr

import (
	"fmt"
	"testing"

	"github.com/facebookgo/ganglia/gmon"
	"github.com/facebookgo/ganglia/gmondtest"
	"github.com/facebookgo/metrics"
)

func TestGaugeSimple(t *testing.T) {
	t.Parallel()
	h := gmondtest.NewHarness(t)
	defer h.Stop()

	const name = "gauge_simple_metric"
	var gauge metrics.Gauge
	gauge = &Gauge{
		Name: name,
	}

	registry := testRegistry(h.Client)
	registry.Register(gauge)

	const v1 = 42
	gauge.Update(v1)
	if gauge.Value() != v1 {
		t.Fatalf("was expecting %d got %d", v1, gauge.Value())
	}

	registry.write()
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

	registry.write()
	h.ContainsMetric(&gmon.Metric{
		Name:  name,
		Value: fmt.Sprint(v2),
		Unit:  "value",
		Slope: "both",
	})
}
