package brokers_test

import (
	"testing"

	"github.com/Guazi-inc/machinery/v1/brokers"
	"github.com/Guazi-inc/machinery/v1/config"
	"github.com/Guazi-inc/machinery/v1/tasks"
	"github.com/stretchr/testify/assert"
)

func TestIsTaskRegistered(t *testing.T) {
	broker := brokers.New(new(config.Config))
	broker.SetRegisteredTaskNames([]string{"foo", "bar"})

	assert.True(t, broker.IsTaskRegistered("foo"))
	assert.False(t, broker.IsTaskRegistered("bogus"))
}

func TestAdjustRoutingKey(t *testing.T) {
	var (
		s      *tasks.Signature
		broker brokers.Broker
	)

	// Signatures with routing key

	s = &tasks.Signature{RoutingKey: "routing_key"}
	broker = brokers.New(&config.Config{
		DefaultQueue: "queue",
		AMQP: &config.AMQPConfig{
			ExchangeType: "direct",
			BindingKey:   "binding_key",
		},
	})
	broker.AdjustRoutingKey(s)
	assert.Equal(t, "routing_key", s.RoutingKey)

	s = &tasks.Signature{RoutingKey: "routing_key"}
	broker = brokers.New(&config.Config{
		DefaultQueue: "queue",
	})
	broker.AdjustRoutingKey(s)
	assert.Equal(t, "routing_key", s.RoutingKey)

	// Signatures without routing key

	s = new(tasks.Signature)
	broker = brokers.New(&config.Config{
		DefaultQueue: "queue",
		AMQP: &config.AMQPConfig{
			ExchangeType: "direct",
			BindingKey:   "binding_key",
		},
	})
	broker.AdjustRoutingKey(s)
	assert.Equal(t, "binding_key", s.RoutingKey)

	s = new(tasks.Signature)
	broker = brokers.New(&config.Config{
		DefaultQueue: "queue",
	})
	broker.AdjustRoutingKey(s)
	assert.Equal(t, "queue", s.RoutingKey)
}
