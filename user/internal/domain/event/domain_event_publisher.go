package event

// DomainEventPublisher định nghĩa contract cho việc publish event.
type DomainEventPublisher interface {
	Publish(event DomainEvent) error
}
