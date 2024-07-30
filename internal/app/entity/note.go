package entity

type Note struct {
    ID      uint   `gorm:"primaryKey" json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
    UserID  uint   `json:"user_id"`
    User    User   `gorm:"foreignKey:UserID" json:"user"`
}
