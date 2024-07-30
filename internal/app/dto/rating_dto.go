package dto

type CreateRatingDTO struct {
    UserID uint `json:"user_id"`
    Value  int  `json:"value"`
}

type UpdateRatingDTO struct {
    Value int `json:"value"`
}
