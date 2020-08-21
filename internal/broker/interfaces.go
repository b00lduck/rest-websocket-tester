package broker

//go:generate mockgen -source=interfaces.go -destination=mocks/broker.go -package=mock_broker

type Broadcaster interface {
	Broadcast(message []byte)
}

type Broker interface {
	Run()
	Message(message []byte)
}
