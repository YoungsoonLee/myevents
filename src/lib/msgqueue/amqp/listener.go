package amqp

import "github.com/streadway/amqp"

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
