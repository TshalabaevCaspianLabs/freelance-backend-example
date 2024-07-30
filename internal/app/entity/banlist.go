package entity

import "time"

type BanList struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    UserID    uint      `json:"user_id"`
    EndDate   time.Time `json:"end_date"`
}
