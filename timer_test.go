package gangliamr

import (
	"testing"
	"time"

	"github.com/facebookgo/ganglia/gmon"
	"github.com/facebookgo/ganglia/gmondtest"
	"github.com/facebookgo/metrics"
)

func TestTimerSeconds(t *testing.T) {
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
	registry.write()
	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".mean",
		Value: "1",
		Unit:  "seconds",
		Slope: "both",
	})

	const v2 = 2 * time.Second
	timer.Update(v2)
	registry.tick()
	registry.write()
	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".mean",
		Value: "1.5",
		Unit:  "seconds",
		Slope: "both",
	})
}

func TestTimerSecondsPartial(t *testing.T) {
	t.Parallel()
	h := gmondtest.NewHarness(t)
	defer h.Stop()

	const name = "timer_seconds_partial_metric"
	var timer metrics.Timer
	timer = &Timer{
		Name: name,
	}

	registry := testRegistry(h.Client)
	registry.Register(timer)

	const v1 = 200 * time.Millisecond
	timer.Update(v1)
	registry.tick()
	registry.write()
	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".mean",
		Value: "0.2",
		Unit:  "seconds",
		Slope: "both",
	})

	const v2 = 500 * time.Millisecond
	timer.Update(v2)
	registry.tick()
	registry.write()
	h.ContainsMetric(&gmon.Metric{
		Name:  name + ".mean",
		Value: "0.35",
		Unit:  "seconds",
		Slope: "both",
	})
}
