package entity

type Chat struct {
	SenderID   string `json:"senderId" `
	ReceiverID string `json:"receiverId"`
	MessageID  string `json:"messageId"`
	Message    string `json:"message" `
}
