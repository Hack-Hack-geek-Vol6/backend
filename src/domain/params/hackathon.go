package params

import "github.com/hackhack-Geek-vol6/backend/src/repository"

type CreateHackathon struct {
	Hackathon  repository.CreateHackathonsParams
	StatusTags []int32 `json:"status_tags"`
}