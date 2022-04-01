package broker_test

import (
	"my-arch/broker"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Broker(t *testing.T) {
	broker := broker.New()

	go func() {
		ch, _ := broker.Subscribe()
		assert.Equal(t, <-ch, "hoge")
	}()

	go func() {
		ch, _ := broker.Subscribe()
		assert.Equal(t, <-ch, "hoge")
	}()

	time.Sleep(time.Second * 1)

	broker.Publish("hoge")

	time.Sleep(time.Second * 1)
}
