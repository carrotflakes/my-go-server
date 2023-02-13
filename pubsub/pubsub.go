package pubsub

import "my-arch/broker"

type Pubsub struct {
	Note *broker.SimpleBroker
}

func NewPubsub() *Pubsub {
	return &Pubsub{
		Note: broker.New(),
	}
}
