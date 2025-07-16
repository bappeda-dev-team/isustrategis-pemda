package domain

import "time"

type Outcome struct {
	ID            int
	PohonId       int
	FaktorOutcome string
	DataTerukur   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
