package notification

import (
	"context"
	"mailgo/lib/log"
	mailer "mailgo/lib/sender"
	notificationtype "mailgo/modules/notification_type"
	"mailgo/modules/template"
	"mailgo/modules/user"
	"time"
)

func getNotificationsByUserService(userID string, ctx context.Context) ([]Notification, error) {
	notifications, err := getNotificationsByUser(userID, ctx)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func CreateNotificationService(notificationDto *EventNotificationDto) error {
	ctx := context.Background()

	// Obtener el tipo de notificación basado en el evento recibido
	notificationType, err := notificationtype.GetNotificationTypeByEventKeyService(
		notificationDto.EventKey, ctx)
	if err != nil {
		log.Get(ctx).Error("Notification type for event key '%s' not found: %v",
			notificationDto.EventKey, err)
		return err
	}
	// Cargar template asociado al tipo de notificación
	templateMail, err := template.FindTemplateByIDService(notificationType.TemplateId, ctx)
	if err != nil {
		log.Get(ctx).Error("Template for notification type '%s' not found: %v", notificationType.ID, err)
		return err
	}

	// Obtener detalles del usuario desde el microservicio auth
	recipientEmail, err := user.GetUserData(notificationDto.UserId, ctx)
	if err != nil {
		return err
	}

	// Construir el contenido del mail usando los detalles del evento
	mailSubject := templateMail.Subject
	mailBodyHTML := replacePlaceholders(templateMail.BodyHTML, notificationDto.EventDetails)
	mailBodyText := replacePlaceholders(templateMail.BodyText, notificationDto.EventDetails)

	// Crear la notificación en la base de datos
	notification := &CreateNotificationDto{
		TypeId:       notificationType.ID.Hex(),
		UserId:       notificationDto.UserId,
		Recipient:    recipientEmail.Email,
		RelatedId:    notificationDto.RelatedId,
		CreatedAt:    time.Now(),
		EventDetails: notificationDto.EventDetails,
		Mail: template.MailNotificationTemplate{
			Subject:  mailSubject,
			BodyHTML: mailBodyHTML,
			BodyText: mailBodyText,
		},
	}

	// Persistir la notificación
	_, err = createNotification(notification, ctx)
	if err != nil {
		return err
	}

	// Opcional: Disparar el envío del correo
	err = mailer.SendEmail(notification.Mail, notification.Recipient)
	if err != nil {
		return err
	}

	return nil
}
