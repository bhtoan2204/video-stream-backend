package initialize

import (
	"github.com/bhtoan2204/user/internal/application/event_bus"
	consumer "github.com/bhtoan2204/user/internal/infrastructure/event_consumer"
)

type Debezium struct {
	eventBus event_bus.EventBus
}

func NewDebezium(eventBus event_bus.EventBus) *Debezium {
	return &Debezium{eventBus: eventBus}
}

func (d *Debezium) Start() *consumer.DebeziumConsumer {
	debeziumConsumer := consumer.NewDebeziumConsumer(&d.eventBus)
	go debeziumConsumer.Consume()
	return debeziumConsumer
}
