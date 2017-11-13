package rabbit

import (
	"fmt"
	"log"
	"github.com/streadway/amqp"
)

type Rabbit struct {
	hostname string
	connection *amqp.Connection
	channel *amqp.Channel
	queue *amqp.Queue
	exchange_name string
	exchange_type string
	queue_name string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// TODO: Replace parameters with config file
// TODO: Implement initializer
// TODO: Implement destructer for defer
// TODO: Make all methods except Constructor/Destructor private (non-exported)

func (r *Rabbit) Initialize(queue_name string, exchange_name string, exchange_type string) {
	r.queue_name = queue_name
	r.exchange_name = exchange_name
	r.exchange_type = exchange_type
	
	r.establishRabbitConnection()
	r.initializeChannel()
	r.initializeQueue()
	r.initializeExchange()
	r.bindQueueToExchange()
}

func (r *Rabbit) Destroy() {

}

func (r *Rabbit) establishRabbitConnection() {
	// TODO - use actual configuration information
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to open a channel")

	r.connection = conn
}

func (r *Rabbit) initializeChannel()  {
	channel, err := r.connection.Channel()
	failOnError(err, "Failed to open a channel")

	r.channel = channel
}

func (r *Rabbit) initializeExchange() {
	err := r.channel.ExchangeDeclare(
		r.exchange_name,   // name
		r.exchange_type, // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to create an exchange")
}

func (r *Rabbit) initializeQueue() {
	
	queue, err := r.channel.QueueDeclare(
		r.queue_name, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments	
	)
	failOnError(err, "Failed to declare a queue")
	r.queue = &queue
}

func (r *Rabbit) bindQueueToExchange() {
	err := r.channel.QueueBind(
		r.queue_name,		// queue name
		"",			// routing key
		r.exchange_name,	// exchange
		false,
		nil,
	)
}

func (r *Rabbit) PublishMesssage(message string) {
	err := r.channel.Publish(
		r.exchange_name,        // exchange
		r.queue_name,		// routing key
		false,			// mandatory
		false,			// immediate
		amqp.Publishing {
			DeliveryMode: amqp.Persistent,
			ContentType: "application/json",
			Body:        []byte(message),
		})
	failOnError(err, "Failed to publish a message")
}

func (r *Rabbit) ReceiveMessages(done chan string) {
	msgs, err := r.channel.Consume(
		r.queue_name, // queue
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


