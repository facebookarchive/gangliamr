package gangliamr

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/daaku/go.ganglia/gmetric"
	"github.com/daaku/go.metrics"
)

// Internally we verify the registered metrics match this interface.
type metric interface {
	writeMeta(c *gmetric.Client)
	writeValue(c *gmetric.Client)
	register(r *Registry)
}

// Registry provides the process to periodically report the in-memory metrics
// to Ganglia.
type Registry struct {
	Prefix            string
	NameSeparator     string // Default is a dot "."
	Client            *gmetric.Client
	WriteTickDuration time.Duration
	startOnce         sync.Once
	metrics           []metric
	mutex             sync.Mutex
}

func (r *Registry) start() {
	go func() {
		sendTicker := time.NewTicker(r.WriteTickDuration)
		metricsTicker := time.NewTicker(metrics.TickDuration)
		for {
			select {
			case <-sendTicker.C:
				ms := r.registered()
				for _, m := range ms {
					m.writeMeta(r.Client)
					m.writeValue(r.Client)
				}
			case <-metricsTicker.C:
				ms := r.registered()
				for _, m := range ms {
					if t, ok := m.(metrics.Tickable); ok {
						t.Tick()
					}
				}
			}
		}
	}()
}

// Register a metric. The only metrics acceptable for registration are the ones
// provided in this package itself. The registration function uses an untyped
// argument to make it easier for use with fields typed as one of the metrics
// in the go.metrics library. All the metrics provided by this library embed
// one of those metrics and augment them with Ganglia specific metadata.
func (r *Registry) Register(m interface{}) {
	r.startOnce.Do(r.start)
	r.mutex.Lock()
	defer r.mutex.Unlock()
	v, ok := m.(metric)
	if !ok {
		panic(fmt.Sprintf("unknown metric type: %T", m))
	}
	v.register(r)
	r.metrics = append(r.metrics, v)
}

func (r *Registry) registered() []metric {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	metrics := make([]metric, len(r.metrics))
	copy(metrics, r.metrics)
	return metrics
}

func (r *Registry) makeName(parts ...string) string {
	var nonempty []string
	if r.Prefix != "" {
		nonempty = append(nonempty, r.Prefix)
	}
	for _, p := range parts {
		if p != "" {
			nonempty = append(nonempty, p)
		}
	}
	return strings.Join(nonempty, r.nameSeparator())
}

func (r *Registry) nameSeparator() string {
	if r.NameSeparator == "" {
		return "."
	}
	return r.NameSeparator
}
