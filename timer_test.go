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

func TestTimerSimple(t *testing.T) {
	t.Parallel()
	h := gmondtest.NewHarness(t)
	defer h.Stop()

	const name = "timer_simple_metric"
	var timer metrics.Timer
	timer = &gangliamr.Timer{
		Name: name,
	}

	registry := gangliamr.Registry{
		Client:            h.Client,
		WriteTickDuration: 5 * time.Millisecond,
	}
	registry.Register(timer)

	const v1 = 43
	timer.Update(v1)
	timer.Tick()
	if timer.Count() != 1 {
		t.Fatalf("was expecting %d got %d", 1, timer.Count())
	}

	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".calls.count",
		Value: "1",
		Unit:  "count",
		Slope: "both",
	})

	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".calls.one-minute",
		Value: "0.18400888292586468",
		Unit:  "nanoseconds",
		Slope: "both",
	})

	const v2 = 42
	timer.Update(v2)
	timer.Tick()
	if timer.Count() != 2 {
		t.Fatalf("was expecting %d got %d", 2, timer.Count())
	}

	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".calls.count",
		Value: fmt.Sprint(timer.Count()),
		Unit:  "count",
		Slope: "both",
	})

	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".calls.one-minute",
		Value: "0.17047269456202283",
		Unit:  "nanoseconds",
		Slope: "both",
	})
}
