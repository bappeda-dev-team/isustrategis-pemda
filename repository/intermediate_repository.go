package repository

import (
	"context"
	"database/sql"
	"isustrategisService/model/domain"
)

type IntermediateRepository interface {
	Create(ctx context.Context, tx *sql.Tx, intermediate *domain.Intermediate) (*domain.Intermediate, error)
	Update(ctx context.Context, tx *sql.Tx, intermediate *domain.Intermediate) (*domain.Intermediate, error)
	Delete(ctx context.Context, tx *sql.Tx, pohonId int) error
	FindById(ctx context.Context, tx *sql.Tx, intermediateId int) (*domain.Intermediate, error)
	FindAll(ctx context.Context, tx *sql.Tx, tahun string) ([]domain.Intermediate, error)
}
