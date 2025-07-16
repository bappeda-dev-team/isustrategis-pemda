package repository

import (
	"context"
	"database/sql"
	"isustrategisService/model/domain"
)

type OutcomeRepository interface {
	Create(ctx context.Context, tx *sql.Tx, outcome *domain.Outcome) (*domain.Outcome, error)
	Update(ctx context.Context, tx *sql.Tx, outcome *domain.Outcome) (*domain.Outcome, error)
	Delete(ctx context.Context, tx *sql.Tx, outcomeId int) error
	FindById(ctx context.Context, tx *sql.Tx, outcomeId int) (*domain.Outcome, error)
	FindAll(ctx context.Context, tx *sql.Tx, pohonId int) ([]domain.Outcome, error)
}
