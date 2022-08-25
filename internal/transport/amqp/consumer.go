package amqp

// Consumers interface controller of amqp
type Consumers interface {
	// Register amqp listeners
	Register()
}
