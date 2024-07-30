package entity

type User struct {
    ID       uint   `gorm:"primaryKey" json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Rating   Rating `json:"rating"`
    IsBanned bool   `json:"is_banned"`
}
