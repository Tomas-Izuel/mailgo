package notification

type CartNotification struct {
	Notification
}

func (cn *CartNotification) Process() error {
	// Implementa la lógica específica para notificaciones de carrito abandonado
	return nil
}
