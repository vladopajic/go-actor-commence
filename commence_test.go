package commence_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vladopajic/go-actor/actor"

	commence "github.com/vladopajic/go-actor-commence"
)

func TestCommencer(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		return
	}

	w := &worker{}
	c := commence.New()
	a := actor.New(w, c.OptOnStart(nil), c.OptOnStop(nil))

	for i := 0; i < 5; i++ {
		a.Start()

		<-c.WaitC()

		a.Stop()
	}
}

func TestCommencerWithOpts(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		return
	}

	onStartC, onStopC := make(chan any, 1), make(chan any, 1)
	w := &worker{}
	c := commence.New()
	a := actor.New(w,
		c.OptOnStart(func(ctx actor.Context) { onStartC <- `🌞` }),
		c.OptOnStop(func() { onStopC <- `🌚` }),
	)

	for i := 0; i < 5; i++ {
		a.Start()

		<-c.WaitC()
		assert.Equal(t, `🌞`, <-onStartC)

		a.Stop()
		assert.Equal(t, `🌚`, <-onStopC)
	}
}

type worker struct{}

func (w *worker) OnStart(ctx actor.Context) {
	// this worker has heavy onStart processes
	time.Sleep(time.Millisecond * 100)
}

func (w *worker) DoWork(ctx actor.Context) actor.WorkerStatus {
	select { //nolint:gosimple // relax
	case <-ctx.Done():
		return actor.WorkerEnd
	}
}
