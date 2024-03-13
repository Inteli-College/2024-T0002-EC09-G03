package ports

type DBPort interface {
	GenerateConsumer(string)
	RetriveLastMessage(interface{}, int) error
}
