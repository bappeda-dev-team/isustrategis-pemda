package repository

import (
	"context"
	"database/sql"
	"isustrategisService/model/domain"
)

type IntermediateRepositoryImpl struct{}

func NewIntermediateRepositoryImpl() *IntermediateRepositoryImpl {
	return &IntermediateRepositoryImpl{}
}

func (repository *IntermediateRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, intermediate *domain.Intermediate) (*domain.Intermediate, error) {
	SQL := `INSERT INTO tb_intermediate (pohon_id, faktor_outcome, data_terukur, tahun, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())`
	_, err := tx.ExecContext(ctx, SQL, intermediate.PohonId, intermediate.FaktorOutcome, intermediate.DataTerukur, intermediate.Tahun)
	if err != nil {
		return nil, err
	}

	return intermediate, nil
}

func (repository *IntermediateRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, intermediate *domain.Intermediate) (*domain.Intermediate, error) {
	SQL := `UPDATE tb_intermediate SET pohon_id = ?, faktor_outcome = ?, data_terukur = ?, tahun = ?, updated_at = NOW() WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, intermediate.PohonId, intermediate.FaktorOutcome, intermediate.DataTerukur, intermediate.Tahun, intermediate.ID)
	if err != nil {
		return nil, err
	}

	return intermediate, nil
}

func (repository *IntermediateRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, pohonId int) error {
	SQL := `DELETE FROM tb_intermediate WHERE pohon_id = ?`
	_, err := tx.ExecContext(ctx, SQL, pohonId)
	if err != nil {
		return err
	}

	return nil
}

func (repository *IntermediateRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, intermediateId int) (*domain.Intermediate, error) {
	SQL := `SELECT id, pohon_id, faktor_outcome, data_terukur, tahun, created_at, updated_at FROM tb_intermediate WHERE id = ?`
	var intermediate domain.Intermediate
	err := tx.QueryRowContext(ctx, SQL, intermediateId).Scan(
		&intermediate.ID,
		&intermediate.PohonId,
		&intermediate.FaktorOutcome,
		&intermediate.DataTerukur,
		&intermediate.Tahun,
		&intermediate.CreatedAt,
		&intermediate.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &intermediate, nil
}

func (repository *IntermediateRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, tahun string) ([]domain.Intermediate, error) {
	SQL := `SELECT id, pohon_id, faktor_outcome, data_terukur, tahun, created_at, updated_at FROM tb_intermediate WHERE tahun = ?`
	rows, err := tx.QueryContext(ctx, SQL, tahun)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var intermediates []domain.Intermediate
	for rows.Next() {
		var intermediate domain.Intermediate
		err = rows.Scan(
			&intermediate.ID,
			&intermediate.PohonId,
			&intermediate.FaktorOutcome,
			&intermediate.DataTerukur,
			&intermediate.Tahun,
			&intermediate.CreatedAt,
			&intermediate.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		intermediates = append(intermediates, intermediate)
	}

	return intermediates, nil
}
