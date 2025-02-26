package initialize

import (
	"github.com/bhtoan2204/user/internal/application/event"
	consumer "github.com/bhtoan2204/user/internal/infrastructure/event_consumer"
)

type Debezium struct {
	eventBus event.EventBus
}

func NewDebezium(eventBus event.EventBus) *Debezium {
	return &Debezium{eventBus: eventBus}
}

func (d *Debezium) Start() {
	debeziumConsumer := consumer.NewDebezium(&d.eventBus)
	go debeziumConsumer.Consume()
}
