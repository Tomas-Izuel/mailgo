package rabbit

import "mailgo/rabbit/consume"

func Init() {
	go consume.ConsumeNegativeFeedback()
}
