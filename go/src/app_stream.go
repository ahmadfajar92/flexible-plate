package src

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"scaffold/shared/interfaces"
	"scaffold/shared/log"

	confluent "github.com/confluentinc/confluent-kafka-go/kafka"
)

// ServeStream func
func (app *application) ServeStream(wg *sync.WaitGroup) {
	app.runKafkaStreamServer(wg)
}

// RunStreamServer func
func (app *application) runKafkaStreamServer(wg *sync.WaitGroup) {
	ctx := "app-RunKafkaStreamServer"

	// start stream server
	log.Log(log.InfoLevel, "Start Stream Server ...", ctx, "")

	// run kafka consumer
	app.kafkaConsumerRunner(wg)
}

func (app *application) kafkaConsumerRunner(wg *sync.WaitGroup) {
	ctx := "app-kafkaConsumerRunner"
	// setup consumer

	if !app.Cfg().Debug() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	}

	// get kafka stream deliveries
	streams := app.deliveries.GetKafkaStreams()
	for _, st := range streams {
		// start consumer kafka
		log.Log(log.InfoLevel, fmt.Sprintf("Start Consumer ::'%s' ...", st.GetGroup()), ctx, "get_consumer_delivery")

		wg.Add(1)
		go func(stream interfaces.DeliveryKafka) {
			app.kafkaConsumer(stream)
			wg.Done()
		}(st)
	}

	// done initiate
	log.Log(log.InfoLevel, "Server is ready!", ctx, "end_initial_stream_server")

}

func (app *application) kafkaConsumer(stream interfaces.DeliveryKafka) {
	ctx := "app-kafkaConsumer"

	consumer, err := confluent.NewConsumer(&confluent.ConfigMap{
		"bootstrap.servers": stream.GetHost(),
		// Avoid connecting to IPv6 brokers:
		// This is needed for the ErrAllBrokersDown show-case below
		// when using localhost brokers on OSX, since the OSX resolver
		// will return the IPv6 addresses first.
		// You typically don't need to specify this configuration property.
		"broker.address.family": "v4",
		"group.id":              stream.GetGroup(),
		"session.timeout.ms":    6000,
		"auto.offset.reset":     stream.GetOffset(),
	})

	if err != nil {
		// TODO: set error
		log.Log(
			log.ErrorLevel,
			err.Error(), ctx,
			fmt.Sprintf("start_confluent_consumer_%s", stream.GetGroup()),
		)
		return
	}

	// close consumer
	defer consumer.Close()

	topics := strings.Split(stream.GetTopic(), ",")
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		// TODO: set error
		log.Log(
			log.ErrorLevel,
			err.Error(), ctx,
			fmt.Sprintf("init_topic_%s", stream.GetTopic()),
		)
		return
	}

	// initial
	log.Log(log.InfoLevel, fmt.Sprintf("Listening to topic ::'%s'\n", stream.GetTopic()), ctx, "")
	run := true

	for run == true {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			if app.Cfg().Debug() {
				log.Log(log.InfoLevel, "msg retrived", ctx, "on_retrive_message")
			}

			err = stream.OnReceived(msg.Value)
			if err != nil {
				log.Log(
					log.InfoLevel,
					fmt.Sprintf("Got an error: %v \n (%v)", err, string(msg.Value)), ctx,
					"error_on_retrieve_msg",
				)
			}
			time.Sleep(1 * time.Second)
		} else {
			// The client will automatically try to recover from all errors.
			log.Log(
				log.ErrorLevel,
				fmt.Sprintf("Consumer error: %v (%v)\n", err, msg), ctx,
				"error_read_msg",
			)
		}
	}
}
