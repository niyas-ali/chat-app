package model

// MessageRecipient model
type MessageRecipient struct {
	ID               int64 `xorm:"PK autoincr bigint"`
	RecipientID      int64 `xorm:"bigint"`
	RecipientGroupID int64 `xorm:"bigint"`
	MessageID        int64 `xorm:"bigint"`
	IsRead           bool  `xorm:"default false"`
}
