package notification

import (
	"context"
	"fmt"
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

func CreateNotificationService(notificationDto *CreateNotificationDto, ctx context.Context) (*Notification, error) {
	// Obtener el tipo de notificación basado en el evento recibido
	currentType, err := notificationtype.GetNotificationTypeByEventKeyService(
		notificationDto.Event, ctx)
	if err != nil {
		return nil, fmt.Errorf("error fetching notification type: %w", err)
	}

	// Cargar el template asociado al tipo de notificación
	templateMail, err := template.FindTemplateByIDService(currentType.
		TemplateId, ctx)
	if err != nil {
		return nil, fmt.Errorf("error fetching template: %w", err)
	}

	// Obtener detalles del usuario desde el microservicio auth
	recipientEmail, err := user.GetUserData(notificationDto.UserId, ctx)
	if err != nil {
		return nil, fmt.Errorf("error fetching user email: %w", err)
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
		Mail:         template.MailNotificationTemplate{
			mailSubject,
			mailBodyHTML,
			mailBodyText,
		},
	}

	// Persistir la notificación
	id, err := createNotification(notification, ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating notification: %w", err)
	}
	notification.ID = id

	// Opcional: Disparar el envío del correo
	err = mailer.SendMail(notification.Mail.Subject, notification.Mail.BodyHTML, notification.Mail.BodyText, recipientEmail)
	if err != nil {
		return nil, fmt.Errorf("error sending mail: %w", err)
	}

	return notification, nil
}
