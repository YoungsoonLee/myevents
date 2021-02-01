package main

import (
	"flag"
	"fmt"
	"log"

	msgqueue_amqp "github.com/YoungsoonLee/myevent/src/lib/msgqueue/amqp"
	"github.com/YoungsoonLee/myevents/src/eventsservice/rest"
	"github.com/YoungsoonLee/myevents/src/lib/configuration"
	"github.com/YoungsoonLee/myevents/src/lib/persistence/dblayer"
	"github.com/streadway/amqp"
)

func main() {
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()

	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	fmt.Println("Connecting to rabitmq")
	conn, err := amqp.Dial(config.AMQPMessageBroker)
	if err != nil {
		panic(err)
	}

	emitter, err := msgqueue_amqp.NewAMQPEventEmitter(conn)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)

	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler))
}
