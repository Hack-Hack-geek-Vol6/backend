package request

import "time"

type CreateHackathon struct {
	Name      string    `form:"name" validate:"required"`
	Link      string    `form:"link" validate:"required"`
	Expired   time.Time `form:"expired" validate:"required"`
	StartDate time.Time `form:"start_date" validate:"required"`
	Term      int       `form:"term" validate:"required"`
	Statuses  []int64   `form:"statuses[]"`
}