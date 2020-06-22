package factory

import (
	"scaffold/shared/interfaces"
	"scaffold/src/module/delivery"
)

type deliveries struct {
	streams *interfaces.Streams
	rest    *interfaces.Rest
	rpc     *interfaces.RPC
}

// Deliveries factory
func Deliveries(app interfaces.Application) interfaces.Deliveries {
	deliveries := new(deliveries)
	deliveries.streams = &interfaces.Streams{}
	deliveries.rpc = &interfaces.RPC{}

	deliveries.rest = &interfaces.Rest{
		HTTP: []interfaces.DeliveryHTTP{
			delivery.Setup(app),
		},
	}

	return deliveries
}

/*
	all getter func
*/
// GetStreams
func (d *deliveries) GetStreams() *interfaces.Streams {
	return d.streams
}

// GetKafkaStreams func
func (d *deliveries) GetKafkaStreams() map[string]interfaces.DeliveryKafka {
	return d.streams.Kafka
}

// GetKafkaStreams func
func (d *deliveries) GetRest() []interfaces.DeliveryHTTP {
	return d.rest.HTTP
}
