package repository

import (
	"context"
	"database/sql"
	"isustrategisService/model/domain"
)

type CsfRepository interface {
	Create(ctx context.Context, tx *sql.Tx, csf *domain.Csf) (*domain.Csf, error)
	Update(ctx context.Context, tx *sql.Tx, csf *domain.Csf) (*domain.Csf, error)
	Delete(ctx context.Context, tx *sql.Tx, idPohon int) error
	FindById(ctx context.Context, tx *sql.Tx, csfId int) (*domain.Csf, error)
	FindByIds(ctx context.Context, tx *sql.Tx, csfId int) (*domain.Csf, error)
	FindAll(ctx context.Context, tx *sql.Tx, tahun string) ([]domain.Csf, error)

	// Tambahkan method untuk alasan dan data terukur
	CreateAlasanKondisi(ctx context.Context, tx *sql.Tx, alasan *domain.AlasanKondisi) (*domain.AlasanKondisi, error)
	UpdateAlasanKondisi(ctx context.Context, tx *sql.Tx, alasan *domain.AlasanKondisi) (*domain.AlasanKondisi, error)
	DeleteAlasanKondisi(ctx context.Context, tx *sql.Tx, alasanId int) error
	FindAlasanByCsfId(ctx context.Context, tx *sql.Tx, csfId int) ([]domain.AlasanKondisi, error)

	CreateDataTerukur(ctx context.Context, tx *sql.Tx, data *domain.DataTerukurPendukungPernyataan) (*domain.DataTerukurPendukungPernyataan, error)
	UpdateDataTerukur(ctx context.Context, tx *sql.Tx, data *domain.DataTerukurPendukungPernyataan) (*domain.DataTerukurPendukungPernyataan, error)
	DeleteDataTerukur(ctx context.Context, tx *sql.Tx, dataId int) error
	FindDataTerukurByAlasanId(ctx context.Context, tx *sql.Tx, alasanId int) ([]domain.DataTerukurPendukungPernyataan, error)
}
