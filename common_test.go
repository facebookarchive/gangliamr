package gangliamr

import (
	"time"

	"github.com/facebookgo/ganglia/gmetric"
)

func testRegistry(client *gmetric.Client) *Registry {
	r := &Registry{
		Client:            client,
		WriteTickDuration: 5 * time.Millisecond,
	}

	// consume the start once to disable the background goroutine
	r.startOnce.Do(func() {})

	return r
}
