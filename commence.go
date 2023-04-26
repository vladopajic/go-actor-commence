package commence

import "github.com/vladopajic/go-actor/actor"

type Waiter interface {
	Wait()
}

type Commencer struct {
	commenceSigC chan any
}

func New() *Commencer {
	return &Commencer{
		commenceSigC: make(chan any),
	}
}

func (c *Commencer) WaitC() <-chan any {
	return c.commenceSigC
}

func (c *Commencer) OptOnStart(onStart func(actor.Context)) actor.Option {
	return actor.OptOnStart(func(ctx actor.Context) {
		if onStart != nil {
			onStart(ctx)
		}

		close(c.commenceSigC)
	})
}

func (c *Commencer) OptOnStop(onStop func()) actor.Option {
	return actor.OptOnStop(func() {
		if onStop != nil {
			onStop()
		}

		c.commenceSigC = make(chan any)
	})
}
