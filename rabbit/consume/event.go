package consume

import (
	"encoding/json"
	"mailgo/lib"
	"mailgo/lib/log"
	"mailgo/modules/notification"

	"github.com/streadway/amqp"
)

func ConsumeEvent() error {
	logger := log.Get().
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "event_queue").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "notification").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Consume")

	conn, err := amqp.Dial(lib.GetEnv().RabbitURL)
	if err != nil {
		logger.Error(err)
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		logger.Error(err)
		return err
	}
	defer chn.Close()

	err = chn.ExchangeDeclare(
		"notification", // name
		"direct",       // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	queue, err := chn.QueueDeclare(
		"event_queue", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name,          // queue name
		"send-notification", // routing key
		"notification",      // exchange
		false,
		nil,
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	messages, err := chn.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	go func() {
		for msg := range messages {
			var eventMsg notification.CreateNotificationDto
			if err := json.Unmarshal(msg.Body, &eventMsg); err != nil {
				logger.Error("Failed to parse message: ", err)
				msg.Nack(false, false)
				continue
			}

			logger.Info("Processing event: %s for user: %s", eventMsg.EventKey,
				eventMsg.UserId)

			if err := processEvent(&eventMsg); err != nil {
				logger.Error("Failed to process event: ", err)
				msg.Nack(false, false)
			} else {
				msg.Ack(false)
			}
		}
	}()

	return nil
}

func processEvent(eventMsg *notification.CreateNotificationDto) error {
	if err := notification.CreateNotificationService(eventMsg); err != nil {
		return err
	}
	return nil
}
