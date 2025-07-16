package repository

import (
	"context"
	"database/sql"
	"isustrategisService/model/domain"
)

type OutcomeRepositoryImpl struct{}

func NewOutcomeRepositoryImpl() *OutcomeRepositoryImpl {
	return &OutcomeRepositoryImpl{}
}

func (repository *OutcomeRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, outcome *domain.Outcome) (*domain.Outcome, error) {
	SQL := `INSERT INTO tb_outcome(pohon_id, faktor_outcome, data_terukur, created_at, updated_at) VALUES (?, ?, ?, NOW(), NOW())`
	_, err := tx.ExecContext(ctx, SQL, outcome.PohonId, outcome.FaktorOutcome, outcome.DataTerukur)
	if err != nil {
		return nil, err
	}

	return outcome, nil
}

func (repository *OutcomeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, outcome *domain.Outcome) (*domain.Outcome, error) {
	SQL := `UPDATE tb_outcome SET pohon_id = ?, faktor_outcome = ?, data_terukur = ?, updated_at = NOW() WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, outcome.PohonId, outcome.FaktorOutcome, outcome.DataTerukur, outcome.ID)
	if err != nil {
		return nil, err
	}

	return outcome, nil
}

func (repository *OutcomeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, outcomeId int) error {
	SQL := `DELETE FROM tb_outcome WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, outcomeId)
	if err != nil {
		return err
	}

	return nil
}
func (repository *OutcomeRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, outcomeId int) (*domain.Outcome, error) {
	SQL := `SELECT id, pohon_id, faktor_outcome, data_terukur, created_at, updated_at FROM tb_outcome WHERE id = ?`
	var outcome domain.Outcome
	err := tx.QueryRowContext(ctx, SQL, outcomeId).Scan(
		&outcome.ID,
		&outcome.PohonId,
		&outcome.FaktorOutcome,
		&outcome.DataTerukur,
		&outcome.CreatedAt,
		&outcome.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &outcome, nil
}

func (repository *OutcomeRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, pohonId int) ([]domain.Outcome, error) {
	SQL := `SELECT id, pohon_id, faktor_outcome, data_terukur, created_at, updated_at FROM tb_outcome WHERE pohon_id = ?`
	rows, err := tx.QueryContext(ctx, SQL, pohonId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var outcomes []domain.Outcome
	for rows.Next() {
		var outcome domain.Outcome
		err = rows.Scan(
			&outcome.ID,
			&outcome.PohonId,
			&outcome.FaktorOutcome,
			&outcome.DataTerukur,
			&outcome.CreatedAt,
			&outcome.UpdatedAt)
		if err != nil {
			return nil, err
		}
		outcomes = append(outcomes, outcome)
	}

	return outcomes, nil
}
