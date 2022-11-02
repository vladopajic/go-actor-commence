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

func (c *Commencer) Wait() {
	<-c.commenceSigC
}

func (c *Commencer) OptOnStart(onStart func()) actor.Option {
	return actor.OptOnStart(func() {
		onStart()
		close(c.commenceSigC)
	})
}
