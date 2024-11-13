package notification

type OrderNotification struct {
	Notification
}

func (on *OrderNotification) Process() error {
	// Implementa la lógica específica para notificaciones de orden completada
	return nil
}
