package ports

type MQTTPort interface {
	CreateClient(*string)

	Publish(*string, byte, bool, interface{})
}
