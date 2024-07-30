package dto

import "time"

type CreateBanListDTO struct {
    UserID  uint      `json:"user_id"`
    EndDate time.Time `json:"end_date"`
}

type UpdateBanListDTO struct {
    EndDate time.Time `json:"end_date"`
}
