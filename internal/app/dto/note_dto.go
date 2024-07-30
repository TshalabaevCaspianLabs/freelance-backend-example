package dto

type CreateNoteDTO struct {
    Title   string `json:"title"`
    Content string `json:"content"`
    UserID  uint   `json:"user_id"`
}

type UpdateNoteDTO struct {
    Title   string `json:"title"`
    Content string `json:"content"`
}
