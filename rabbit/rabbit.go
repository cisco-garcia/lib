package rabbit

import (
	"fmt"
	"log"
	"github.com/streadway/amqp"
)

type Rabbit struct {
	hostname *string
	connection *amqp.Connection
	channel *amqp.Channel
	queue *amqp.Queue
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// TODO: Refactor into functions
// TODO: Replace parameters with config file
// TODO: Implement initializer
// TODO: Implement destructoer for defer

func (r *Rabbit) Initialize() {

}

func (r *Rabbit) Destroy() {

}

func (r *Rabbit) EstablishRabbitConnection() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to open a channel")

	r.connection = conn
}

func (r *Rabbit) InitializeExchange() {
	err = ch.ExchangeDeclare(
		"logs",   // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to create an exchange")
}

func (r *Rabbit) InitializeChannel()  {
	channel, err := r.connection.Channel()
	failOnError(err, "Failed to open a channel")

	r.channel = channel
}

func (r *Rabbit) InitializeQueue(queueName string) {
	
	queue, err := r.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments	
	)
	failOnError(err, "Failed to declare a queue")
	r.queue = &queue
}

func (r *Rabbit) PublishMesssage(message string) {
	err := r.channel.Publish(
		"",           // exchange
		r.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing {
			DeliveryMode: amqp.Persistent,
			ContentType: "application/json",
			Body:        []byte(message),
		})
	failOnError(err, "Failed to publish a message")
}

func (r *Rabbit) ReceiveMessages(done chan string) {
	msgs, err := r.channel.Consume(
		r.queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args

	)
	failOnError(err, "Failed to consume a message")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			done <- string(d.Body)
			d.Ack(false)
		}
	}()
	<- forever
}


