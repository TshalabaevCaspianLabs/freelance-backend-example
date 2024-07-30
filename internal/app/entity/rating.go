package entity

type Rating struct {
    ID     uint `gorm:"primaryKey" json:"id"`
    UserID uint `json:"user_id"`
    Value  int  `json:"value"`
}
