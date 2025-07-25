package domain

import "time"

type Outcome struct {
	ID            int
	PohonId       int
	Tahun         string
	FaktorOutcome string
	DataTerukur   string
	ParentId      *int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
