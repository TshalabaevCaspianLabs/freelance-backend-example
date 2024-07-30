package entity

type Notification struct {
    ID        uint   `gorm:"primaryKey" json:"id"`
    UserID    uint   `json:"user_id"`
    Message   string `json:"message"`
    IsSent    bool   `json:"is_sent"`
}
