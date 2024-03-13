package ports

type QueuePort interface {
	GenerateConsumer(string)
	RetriveLastMessage(string) *[]byte
}
