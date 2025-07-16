package domain

import "time"

type Intermediate struct {
	ID            int
	PohonId       int
	FaktorOutcome string
	DataTerukur   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
