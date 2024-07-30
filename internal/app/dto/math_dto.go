package dto

import "time"

type CreateMatchDTO struct {
    TournamentID uint      `json:"tournament_id"`
    TeamAID      uint      `json:"team_a_id"`
    TeamBID      uint      `json:"team_b_id"`
    Schedule     time.Time `json:"schedule"`
}

type UpdateMatchDTO struct {
    TeamAID  uint      `json:"team_a_id"`
    TeamBID  uint      `json:"team_b_id"`
    Schedule time.Time `json:"schedule"`
    ScoreA   int       `json:"score_a"`
    ScoreB   int       `json:"score_b"`
}
