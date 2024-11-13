package notification

type FeedbackNotification struct {
	Notification
}

func (fn *FeedbackNotification) Process() error {
	// Implementa la lógica específica para notificaciones de feedback
	return nil
}
