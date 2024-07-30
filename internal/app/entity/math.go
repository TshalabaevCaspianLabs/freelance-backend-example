package entity

import "time"

type Match struct {
    ID           uint      `gorm:"primaryKey"`
    TournamentID uint      `gorm:"not null"`
    TeamAID      uint      `gorm:"not null"`
    TeamBID      uint      `gorm:"not null"`
    Schedule     time.Time `gorm:"not null"`
    ScoreA       int       `gorm:"default:0"`
    ScoreB       int       `gorm:"default:0"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
}
