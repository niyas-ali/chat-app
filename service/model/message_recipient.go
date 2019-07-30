package model

// MessageRecipient model
type MessageRecipient struct {
	ID               int
	RecipientID      int
	RecipientGroupID int
	MessageID        int
	IsRead           bool
}
