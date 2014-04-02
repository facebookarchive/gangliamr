package gangliamr_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/facebookgo/ganglia/gmon"
	"github.com/facebookgo/ganglia/gmondtest"
	"github.com/facebookgo/gangliamr"
	"github.com/facebookgo/metrics"
)

func TestRegistryPrefix(t *testing.T) {
	t.Parallel()
	h := gmondtest.NewHarness(t)
	defer h.Stop()

	const name = "gauge_simple_metric"
	var gauge metrics.Gauge
	gauge = &gangliamr.Gauge{
		Name: name,
	}

	const prefix = "prefix"
	registry := gangliamr.Registry{
		Prefix:            prefix,
		Client:            h.Client,
		WriteTickDuration: 5 * time.Millisecond,
	}
	registry.Register(gauge)

	const v1 = 42
	gauge.Update(v1)
	h.ContainsMetric(&gmon.Metric{
		Name:  prefix + "." + name,
		Value: fmt.Sprint(v1),
		Unit:  "value",
		Slope: "both",
	})
}

func TestRegistryInvalidRegister(t *testing.T) {
	t.Parallel()
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("was expecting a panic")
		}
	}()
	registry := gangliamr.Registry{
		WriteTickDuration: 5 * time.Millisecond,
	}
	registry.Register(1)
}

func TestTestRegistryAndGet(t *testing.T) {
	t.Parallel()
	const name = "gauge_simple_metric"
	var gauge metrics.Gauge
	gauge = &gangliamr.Gauge{
		Name: name,
	}

	registry := gangliamr.NewTestRegistry()
	registry.Register(gauge)

	if actual := registry.Get(name); actual != gauge {
		t.Fatal("did not find expected metric")
	}
}
