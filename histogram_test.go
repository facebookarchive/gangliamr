package gangliamr

import (
	"testing"

	"github.com/facebookgo/ganglia/gmon"
	"github.com/facebookgo/ganglia/gmondtest"
	"github.com/facebookgo/metrics"
)

func TestHistogramSimple(t *testing.T) {
	t.Parallel()
	h := gmondtest.NewHarness(t)
	defer h.Stop()

	const name = "histogram_simple_metric"
	var hg metrics.Histogram
	hg = &Histogram{
		Name: name,
	}

	registry := testRegistry(h.Client)
	registry.Register(hg)

	const v1 = 43
	hg.Update(v1)

	registry.write()

	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".mean",
		Value: "43",
		Unit:  "value",
		Slope: "both",
	})

	const v2 = 42
	hg.Update(v2)

	registry.write()

	if hg.Max() != v1 {
		t.Fatalf("was expecting %d but got %d", v1, hg.Max())
	}

	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".mean",
		Value: "42.5",
		Unit:  "value",
		Slope: "both",
	})
}
