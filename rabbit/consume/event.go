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

	// Conexión a RabbitMQ
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

	// Declarar Exchange
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

	// Declarar Cola
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

	// Enlazar Cola con Exchange
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

	// Consumir Mensajes
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

	// Procesar Mensajes
	go func() {
		for msg := range messages {
			func() {
				// Manejar panics
				defer func() {
					if r := recover(); r != nil {
						logger.Error("Recovered from panic: ", r)
						msg.Nack(false, false) // Rechazar el mensaje sin reintento
					}
				}()

				// Deserializar mensaje
				var eventMsg notification.EventNotificationDto
				if err := json.Unmarshal(msg.Body, &eventMsg); err != nil {
					logger.Error("Failed to parse message: ", err)
					msg.Nack(false, false) // Rechazar el mensaje sin reintento
					return
				}

				logger.Info("Processing event: %s for user: %s", eventMsg.EventKey, eventMsg.UserId)

				// Procesar Evento
				if err := processEvent(&eventMsg); err != nil {
					logger.Error("Failed to process event: ", err)
					msg.Nack(false, false) // Rechazar el mensaje sin reintento
				} else {
					msg.Ack(false) // Confirmar mensaje procesado
				}
			}()
		}
	}()

	return nil
}

// Lógica de procesamiento del evento
func processEvent(eventMsg *notification.EventNotificationDto) error {
	if err := notification.CreateNotificationService(eventMsg); err != nil {
		return err
	}
	return nil
}
