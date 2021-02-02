package amqp

import (
	"github.com/YoungsoonLee/myevents/src/lib/msgqueue"
	"github.com/streadway/amqp"
)

type amqpEventListener struct {
	connection *amqp.Connection
	queue      string
}

func (a *amqpEventListener) setup() error {
	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	_, err = channel.QueueDeclare(a.queue, true, false, false, false, nil)
	return err

}

func NewAMQPEventListener(conn *amqp.Connection, queue string) (msgqueue.EventListener, error) {
	listener := &amqpEventListener{
		connection: conn,
		queue:      queue,
	}

	err := listener.setup()
	if err != nil {
		return nil, err
	}

	return listener, nil
}

func (a *amqpEventListener) Listen(eventNames ...string) (<-chan masgqueue.Event, <-chan error, error) {
	channel, err := a.connection.Channel()
	if err != nil {
		return nil, nil, err
	}

	defer channel.Close()

	for _, eventName := range eventNames {
		if err := channel.QueueBind(a.queue, eventName, "events", false, nil); err != nil {
			return nil, nil, err
		}
	}
}

func (a *ammqpEventListener) Listen(eventNames ...string (<-chan msgqueue.Event, <-chan error, error) {
	
}