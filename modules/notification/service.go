package notification

import (
	"context"
	mailer "mailgo/lib/sender"
	notificationtype "mailgo/modules/notification_type"
	"mailgo/modules/template"
	"mailgo/modules/user"
	"time"
)

func getNotificationsByUserService(userID string,
	ctx context.Context) ([]Notification, error) {
	notifications, err := getNotificationsByUser(userID, ctx)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func CreateNotificationService(notificationDto *CreateNotificationDto) error {
	ctx := context.Background()

	// Obtener el tipo de notificación basado en el evento recibido
	currentType, err := notificationtype.GetNotificationTypeByEventKeyService(
		notificationDto.EventKey, ctx)
	if err != nil {
		return err
	}

	// Cargar el template asociado al tipo de notificación
	templateMail, err := template.FindTemplateByIDService(currentType.
		TemplateId, ctx)
	if err != nil {
		return nil
	}

	// Obtener detalles del usuario desde el microservicio auth
	recipientEmail, err := user.GetUserData(notificationDto.UserId, ctx)
	if err != nil {
		return err
	}

	// Construir el contenido del mail usando los detalles del evento
	mailSubject := templateMail.Subject
	mailBodyHTML := replacePlaceholders(templateMail.BodyHTML,
		notificationDto.EventDetails)
	mailBodyText := replacePlaceholders(templateMail.BodyText,
		notificationDto.EventDetails)

	// Crear la notificación en la base de datos
	notification := &Notification{
		TypeId:       currentType.ID,
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
	id, err := createNotification(notification, ctx)
	if err != nil {
		return err
	}
	notification.ID = id

	// Opcional: Disparar el envío del correo
	err = mailer.SendEmail(notification.Mail, notification.Recipient)
	if err != nil {
		return err
	}

	return nil
}
