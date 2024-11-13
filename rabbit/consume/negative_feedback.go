package consume

import (
	"encoding/json"
	"mailgo/lib"
	"mailgo/lib/log"
	"mailgo/modules/notification"

	"github.com/streadway/amqp"
)

func ConsumeNegativeFeedback() error {
	logger := log.Get().
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "feedback_negative_feedback").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "negative_feedback").
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
		"negative_feedback", // name
		"direct",            // type
		false,               // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	queue, err := chn.QueueDeclare(
		"feedback_negative_feedback", // name
		false,                        // durable
		false,                        // delete when unused
		false,                        // exclusive
		false,                        // no-wait
		nil,                          // arguments
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name,          // queue name
		"",                  // routing key
		"negative_feedback", // exchange
		false,
		nil,
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	mgs, err := chn.Consume(
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
		for msg := range mgs {
			var feedbackMsg notification.FeedbackNotificationMessage
			if err := json.Unmarshal(msg.Body, &feedbackMsg); err != nil {
				logger.Error("Failed to parse message: ", err)
				msg.Nack(false, false)
				continue
			}

			logger.Info("Received negative feedback from user: ", feedbackMsg.UserId)

			if err := processNegativeFeedback(feedbackMsg); err != nil {
				logger.Error("Failed to process negative feedback: ", err)
				msg.Nack(false, false)
			} else {
				msg.Ack(false)
			}
		}
	}()

	return nil
}

func processNegativeFeedback(feedbackMsg notification.FeedbackNotificationMessage) error {
	var createFeedbackDto notification.CreateNotificationDto = notification.CreateNotificationDto{
		Type:         "feedback",
		RelatedId:    feedbackMsg.ArticleId,
		EventDetails: map[string]interface{}{},
	}

	if err := notification.CreateNotification(createFeedbackDto); err != nil {
		return err
	}
	return nil
}
