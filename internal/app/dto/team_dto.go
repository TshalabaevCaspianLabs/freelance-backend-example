package dto

type CreateTeamDTO struct {
    Name    string `json:"name"`
    Members []uint `json:"members"` // IDs of the members
}

type UpdateTeamDTO struct {
    Name    string `json:"name"`
    Members []uint `json:"members"` // IDs of the members
}
