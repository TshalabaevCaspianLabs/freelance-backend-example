package dto

import "time"

type CreateTournamentDTO struct {
    Name      string    `json:"name"`
    StartDate time.Time `json:"start_date"`
    EndDate   time.Time `json:"end_date"`
}

type UpdateTournamentDTO struct {
    Name      string    `json:"name"`
    StartDate time.Time `json:"start_date"`
    EndDate   time.Time `json:"end_date"`
}
