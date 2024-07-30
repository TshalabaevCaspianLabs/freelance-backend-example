package entity

import "time"

type Tournament struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    Name        string    `json:"name"`
    StartDate   time.Time `json:"start_date"`
    EndDate     time.Time `json:"end_date"`
    IsClosed    bool      `json:"is_closed"`
    Participants []User    `gorm:"many2many:tournament_users;" json:"participants"`
    Matches     []Match   `json:"matches"`
}
