package entity

type Team struct {
    ID      uint   `gorm:"primaryKey" json:"id"`
    Name    string `json:"name"`
    Members []User `gorm:"many2many:team_users;" json:"members"`
}
