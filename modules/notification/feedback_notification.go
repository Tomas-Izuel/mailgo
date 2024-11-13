package notification

type FeedbackNotification struct {
	Notification
}

func (fn *FeedbackNotification) Process() error {
	// Implementa la lógica específica para notificaciones de feedback
	return nil
}

type FeedbackNotificationMessage struct {
	ArticleId    string `json:"user_id"`
	UserId       string `json:"article_id"`
	FeedbackInfo string `json:"feedback_info"`
	Rating       int    `json:"rating"`
	Reason       string `json:"reason"`
}
