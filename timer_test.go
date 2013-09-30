package gangliamr

import (
	"fmt"
	"testing"
	"time"

	"github.com/daaku/go.ganglia/gmon"
	"github.com/daaku/go.ganglia/gmondtest"
	"github.com/daaku/go.metrics"
)

func TestTimerSimple(t *testing.T) {
	t.Parallel()
	h := gmondtest.NewHarness(t)
	defer h.Stop()

	const name = "timer_simple_metric"
	var timer metrics.Timer
	timer = &Timer{
		Name: name,
	}

	registry := testRegistry(h.Client)
	registry.Register(timer)

	const v1 = time.Second
	timer.Update(v1)
	registry.tick()
	if timer.Count() != 1 {
		t.Fatalf("was expecting %d got %d", 1, timer.Count())
	}

	registry.write()
	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".calls.count",
		Value: "1",
		Unit:  "count",
		Slope: "both",
	})

	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".calls.one-minute",
		Value: "0.2",
		Unit:  "nanoseconds",
		Slope: "both",
	})

	const v2 = 2 * time.Second
	timer.Update(v2)
	registry.tick()
	if timer.Count() != 2 {
		t.Fatalf("was expecting %d got %d", 2, timer.Count())
	}

	registry.write()
	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".calls.count",
		Value: fmt.Sprint(timer.Count()),
		Unit:  "count",
		Slope: "both",
	})

	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".calls.one-minute",
		Value: "0.2",
		Unit:  "nanoseconds",
		Slope: "both",
	})
}
