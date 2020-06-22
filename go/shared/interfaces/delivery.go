package interfaces

// Rest struct
type Rest struct {
	HTTP []DeliveryHTTP
}

// RPC struct
type RPC struct{}

// Streams struct
type Streams struct {
	Kafka map[string]DeliveryKafka
}

// DeliveryKafka interface
type DeliveryKafka interface {
	OnReceived(message []byte) error
	Publish(payload interface{}) error
	// all getter func
	GetTopic() string
	GetHost() string
	GetGroup() string
	GetOffset() string
}

// DeliveryHTTP interface
type DeliveryHTTP interface {
	Mount(routes interface{})
	GetPath() string
}

// Deliveries interface
type Deliveries interface {
	GetStreams() *Streams
	GetKafkaStreams() map[string]DeliveryKafka
	GetRest() []DeliveryHTTP
}
