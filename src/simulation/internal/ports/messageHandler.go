package ports

type MessageHandlerPort interface {
	Consume(QueuePort)
}
