package dto

type CreateNotificationDTO struct {
    UserID  uint   `json:"user_id"`
    Message string `json:"message"`
}

type UpdateNotificationDTO struct {
    Message string `json:"message"`
}
