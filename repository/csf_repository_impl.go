package repository

import (
	"context"
	"database/sql"
	"errors"
	"isustrategisService/model/domain"
)

type CsfRepositoryImpl struct{}

func NewCsfRepositoryImpl() *CsfRepositoryImpl {
	return &CsfRepositoryImpl{}
}

func (repository *CsfRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, csf *domain.Csf) (*domain.Csf, error) {
	SQL := `INSERT INTO tb_csf(pohon_id, pernyataan_kondisi_strategis, created_at, updated_at) 
			VALUES (?, ?, NOW(), NOW())`

	result, err := tx.ExecContext(ctx, SQL, csf.PohonId, csf.PernyataanKondisiStrategis)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Get the created data
	SQL = `SELECT id, created_at, updated_at FROM tb_csf WHERE id = ?`
	err = tx.QueryRowContext(ctx, SQL, id).Scan(
		&csf.Id,
		&csf.CreatedAt,
		&csf.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return csf, nil
}

func (repository *CsfRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, csf *domain.Csf) (*domain.Csf, error) {
	// 1. Update CSF utama
	SQL := `UPDATE tb_csf SET 
            pohon_id = ?, 
            pernyataan_kondisi_strategis = ?,
            updated_at = NOW()
            WHERE id = ?`

	result, err := tx.ExecContext(ctx, SQL, csf.PohonId, csf.PernyataanKondisiStrategis, csf.Id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, errors.New("csf not found")
	}

	// 2. Hapus data terukur yang tidak ada di request
	existingAlasan, err := repository.FindAlasanByCsfId(ctx, tx, csf.Id)
	if err != nil {
		return nil, err
	}

	// Buat map untuk tracking alasan yang akan dipertahankan
	keepAlasanIds := make(map[int]bool)
	for _, alasan := range csf.AlasanKondisi {
		if alasan.Id > 0 {
			keepAlasanIds[alasan.Id] = true
		}
	}

	// Hapus alasan yang tidak ada di request
	for _, existing := range existingAlasan {
		if !keepAlasanIds[existing.Id] {
			// Hapus data terukur terlebih dahulu
			err = repository.DeleteDataTerukurByAlasanId(ctx, tx, existing.Id)
			if err != nil {
				return nil, err
			}
			// Kemudian hapus alasan
			err = repository.DeleteAlasanKondisi(ctx, tx, existing.Id)
			if err != nil {
				return nil, err
			}
		}
	}

	// 3. Update atau Create Alasan Kondisi
	for _, alasan := range csf.AlasanKondisi {
		if alasan.Id > 0 {
			// Update existing alasan
			updateAlasanSQL := `UPDATE tb_alasan_kondisi SET 
                              alasan_kondisi_strategis = ?,
                              updated_at = NOW()
                              WHERE id = ? AND csf_id = ?`
			_, err = tx.ExecContext(ctx, updateAlasanSQL, alasan.AlasanKondisiStrategis, alasan.Id, csf.Id)
		} else {
			// Create new alasan
			insertAlasanSQL := `INSERT INTO tb_alasan_kondisi 
                              (csf_id, alasan_kondisi_strategis, created_at, updated_at)
                              VALUES (?, ?, NOW(), NOW())`
			result, err = tx.ExecContext(ctx, insertAlasanSQL, csf.Id, alasan.AlasanKondisiStrategis)
			if err != nil {
				return nil, err
			}
			lastId, err := result.LastInsertId()
			if err != nil {
				return nil, err
			}
			alasan.Id = int(lastId)
		}
		if err != nil {
			return nil, err
		}

		// Buat map untuk tracking data terukur yang akan dipertahankan
		keepDataIds := make(map[int]bool)
		for _, dt := range alasan.DataTerukurPendukungPernyataan {
			if dt.Id > 0 {
				keepDataIds[dt.Id] = true
			}
		}

		// Hapus data terukur yang tidak ada di request
		existingData, err := repository.FindDataTerukurByAlasanId(ctx, tx, alasan.Id)
		if err != nil {
			return nil, err
		}

		for _, existing := range existingData {
			if !keepDataIds[existing.Id] {
				err = repository.DeleteDataTerukur(ctx, tx, existing.Id)
				if err != nil {
					return nil, err
				}
			}
		}

		// Update atau Create Data Terukur
		for _, dt := range alasan.DataTerukurPendukungPernyataan {
			if dt.Id > 0 {
				// Update existing data terukur
				updateDataSQL := `UPDATE tb_data_terukur_pendukung_pernyataan SET 
                                data_terukur = ?,
                                updated_at = NOW()
                                WHERE id = ? AND alasan_kondisi_id = ?`
				_, err = tx.ExecContext(ctx, updateDataSQL, dt.DataTerukur, dt.Id, alasan.Id)
			} else {
				// Create new data terukur
				insertDataSQL := `INSERT INTO tb_data_terukur_pendukung_pernyataan 
                                (alasan_kondisi_id, data_terukur, created_at, updated_at)
                                VALUES (?, ?, NOW(), NOW())`
				result, err = tx.ExecContext(ctx, insertDataSQL, alasan.Id, dt.DataTerukur)
				if err != nil {
					return nil, err
				}
				lastId, err := result.LastInsertId()
				if err != nil {
					return nil, err
				}
				dt.Id = int(lastId)
			}
			if err != nil {
				return nil, err
			}
		}
	}

	// 4. Ambil data terbaru
	return repository.FindById(ctx, tx, csf.Id)
}

// Tambahkan fungsi helper
func (repository *CsfRepositoryImpl) DeleteDataTerukurByAlasanId(ctx context.Context, tx *sql.Tx, alasanId int) error {
	SQL := `DELETE FROM tb_data_terukur_pendukung_pernyataan WHERE alasan_kondisi_id = ?`
	_, err := tx.ExecContext(ctx, SQL, alasanId)
	return err
}

func (repository *CsfRepositoryImpl) CreateAlasanKondisi(ctx context.Context, tx *sql.Tx, alasan *domain.AlasanKondisi) (*domain.AlasanKondisi, error) {
	SQL := `INSERT INTO tb_alasan_kondisi(csf_id, alasan_kondisi_strategis, created_at, updated_at) 
			VALUES (?, ?, NOW(), NOW())`

	result, err := tx.ExecContext(ctx, SQL, alasan.CSFid, alasan.AlasanKondisiStrategis)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Get the created data
	SQL = `SELECT id, created_at, updated_at FROM tb_alasan_kondisi WHERE id = ?`
	err = tx.QueryRowContext(ctx, SQL, id).Scan(
		&alasan.Id,
		&alasan.CreatedAt,
		&alasan.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return alasan, nil
}

func (repository *CsfRepositoryImpl) CreateDataTerukur(ctx context.Context, tx *sql.Tx, data *domain.DataTerukurPendukungPernyataan) (*domain.DataTerukurPendukungPernyataan, error) {
	SQL := `INSERT INTO tb_data_terukur_pendukung_pernyataan(alasan_kondisi_id, data_terukur, created_at, updated_at) 
			VALUES (?, ?, NOW(), NOW())`

	result, err := tx.ExecContext(ctx, SQL, data.AlasanKondisiId, data.DataTerukur)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Get the created data
	SQL = `SELECT id, created_at, updated_at FROM tb_data_terukur_pendukung_pernyataan WHERE id = ?`
	err = tx.QueryRowContext(ctx, SQL, id).Scan(
		&data.Id,
		&data.CreatedAt,
		&data.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// func (repository *CsfRepositoryImpl) UpdateAlasanKondisi(ctx context.Context, tx *sql.Tx, alasan *domain.AlasanKondisi) (*domain.AlasanKondisi, error) {
// 	SQL := `UPDATE tb_alasan_kondisi SET
// 			alasan_kondisi_strategis = ?,
// 			updated_at = NOW()
// 			WHERE id = ?`

// 	result, err := tx.ExecContext(ctx, SQL, alasan.AlasanKondisiStrategis, alasan.Id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return nil, err
// 	}

// 	if rowsAffected == 0 {
// 		return nil, sql.ErrNoRows
// 	}

// 	// Get updated data
// 	SQL = `SELECT id, csf_id, alasan_kondisi_strategis, created_at, updated_at
// 		   FROM tb_alasan_kondisi WHERE id = ?`
// 	err = tx.QueryRowContext(ctx, SQL, alasan.Id).Scan(
// 		&alasan.Id,
// 		&alasan.CSFid,
// 		&alasan.AlasanKondisiStrategis,
// 		&alasan.CreatedAt,
// 		&alasan.UpdatedAt,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return alasan, nil
// }

// func (repository *CsfRepositoryImpl) UpdateDataTerukur(ctx context.Context, tx *sql.Tx, data *domain.DataTerukurPendukungPernyataan) (*domain.DataTerukurPendukungPernyataan, error) {
// 	SQL := `UPDATE tb_data_terukur_pendukung_pernyataan SET
// 			data_terukur = ?,
// 			updated_at = NOW()
// 			WHERE id = ?`

// 	result, err := tx.ExecContext(ctx, SQL, data.DataTerukur, data.Id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return nil, err
// 	}

// 	if rowsAffected == 0 {
// 		return nil, sql.ErrNoRows
// 	}

// 	// Get updated data
// 	SQL = `SELECT id, alasan_kondisi_id, data_terukur, created_at, updated_at
// 		   FROM tb_data_terukur_pendukung_pernyataan WHERE id = ?`
// 	err = tx.QueryRowContext(ctx, SQL, data.Id).Scan(
// 		&data.Id,
// 		&data.AlasanKondisiId,
// 		&data.DataTerukur,
// 		&data.CreatedAt,
// 		&data.UpdatedAt,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

//		return data, nil
//	}
func (repository *CsfRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, idPohon int) error {
	SQL := `DELETE FROM tb_csf WHERE pohon_id = ?`
	_, err := tx.ExecContext(ctx, SQL, idPohon)
	return err
}

func (repository *CsfRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, csfId int) (*domain.Csf, error) {
	SQL := `SELECT id, pohon_id, pernyataan_kondisi_strategis, created_at, updated_at 
            FROM tb_csf WHERE id = ?`

	csf := &domain.Csf{}
	err := tx.QueryRowContext(ctx, SQL, csfId).Scan(
		&csf.Id,
		&csf.PohonId,
		&csf.PernyataanKondisiStrategis,
		&csf.CreatedAt,
		&csf.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	// Ambil alasan kondisi
	alasanList, err := repository.FindAlasanByCsfId(ctx, tx, csfId)
	if err != nil {
		return nil, err
	}
	csf.AlasanKondisi = alasanList

	// Ambil data terukur untuk setiap alasan
	for i := range alasanList {
		dataTerukurList, err := repository.FindDataTerukurByAlasanId(ctx, tx, alasanList[i].Id)
		if err != nil {
			return nil, err
		}
		csf.AlasanKondisi[i].DataTerukurPendukungPernyataan = dataTerukurList
	}

	return csf, nil
}

func (repository *CsfRepositoryImpl) FindByIds(ctx context.Context, tx *sql.Tx, csfId int) (*domain.Csf, error) {
	// Get CSF data
	csfSQL := `SELECT id, pohon_id, pernyataan_kondisi_strategis, created_at, updated_at 
               FROM tb_csf WHERE id = ?`

	csf := &domain.Csf{}
	err := tx.QueryRowContext(ctx, csfSQL, csfId).Scan(
		&csf.Id,
		&csf.PohonId,
		&csf.PernyataanKondisiStrategis,
		&csf.CreatedAt,
		&csf.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	// Get Alasan Kondisi
	alasanSQL := `SELECT id, csf_id, alasan_kondisi_strategis, created_at, updated_at 
                  FROM tb_alasan_kondisi WHERE csf_id = ?`

	alasanRows, err := tx.QueryContext(ctx, alasanSQL, csfId)
	if err != nil {
		return nil, err
	}
	defer alasanRows.Close()

	var alasanList []domain.AlasanKondisi // Ubah ke slice of value bukan pointer
	for alasanRows.Next() {
		var alasan domain.AlasanKondisi // Ubah ke value bukan pointer
		err := alasanRows.Scan(
			&alasan.Id,
			&alasan.CSFid,
			&alasan.AlasanKondisiStrategis,
			&alasan.CreatedAt,
			&alasan.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Get Data Terukur for each Alasan
		dataTerukurSQL := `SELECT id, alasan_kondisi_id, data_terukur, created_at, updated_at 
                          FROM tb_data_terukur_pendukung_pernyataan 
                          WHERE alasan_kondisi_id = ?`

		dataTerukurRows, err := tx.QueryContext(ctx, dataTerukurSQL, alasan.Id)
		if err != nil {
			return nil, err
		}
		defer dataTerukurRows.Close()

		var dataTerukurList []domain.DataTerukurPendukungPernyataan // Ubah ke slice of value
		for dataTerukurRows.Next() {
			var dataTerukur domain.DataTerukurPendukungPernyataan // Ubah ke value
			err := dataTerukurRows.Scan(
				&dataTerukur.Id,
				&dataTerukur.AlasanKondisiId,
				&dataTerukur.DataTerukur,
				&dataTerukur.CreatedAt,
				&dataTerukur.UpdatedAt,
			)
			if err != nil {
				return nil, err
			}
			dataTerukurList = append(dataTerukurList, dataTerukur)
		}
		alasan.DataTerukurPendukungPernyataan = dataTerukurList
		alasanList = append(alasanList, alasan)
	}
	csf.AlasanKondisi = alasanList

	return csf, nil
}

func (repository *CsfRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, pohonId int) ([]domain.Csf, error) {
	SQL := `SELECT id, pohon_id, pernyataan_kondisi_strategis, created_at, updated_at 
			FROM tb_csf WHERE pohon_id = ?`

	rows, err := tx.QueryContext(ctx, SQL, pohonId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var csfs []domain.Csf
	for rows.Next() {
		csf := domain.Csf{}
		err := rows.Scan(
			&csf.Id,
			&csf.PohonId,
			&csf.PernyataanKondisiStrategis,
			&csf.CreatedAt,
			&csf.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		csfs = append(csfs, csf)
	}

	return csfs, nil
}

func (repository *CsfRepositoryImpl) DeleteAlasanKondisi(ctx context.Context, tx *sql.Tx, alasanId int) error {
	SQL := `DELETE FROM tb_alasan_kondisi WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, alasanId)
	return err
}

func (repository *CsfRepositoryImpl) FindAlasanByCsfId(ctx context.Context, tx *sql.Tx, csfId int) ([]domain.AlasanKondisi, error) {
	SQL := `SELECT id, csf_id, alasan_kondisi_strategis, created_at, updated_at 
			FROM tb_alasan_kondisi WHERE csf_id = ?`

	rows, err := tx.QueryContext(ctx, SQL, csfId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alasanList []domain.AlasanKondisi
	for rows.Next() {
		alasan := domain.AlasanKondisi{}
		err := rows.Scan(
			&alasan.Id,
			&alasan.CSFid,
			&alasan.AlasanKondisiStrategis,
			&alasan.CreatedAt,
			&alasan.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		alasanList = append(alasanList, alasan)
	}

	return alasanList, nil
}

func (repository *CsfRepositoryImpl) DeleteDataTerukur(ctx context.Context, tx *sql.Tx, dataId int) error {
	SQL := `DELETE FROM tb_data_terukur_pendukung_pernyataan WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, dataId)
	return err
}

func (repository *CsfRepositoryImpl) FindDataTerukurByAlasanId(ctx context.Context, tx *sql.Tx, alasanId int) ([]domain.DataTerukurPendukungPernyataan, error) {
	SQL := `SELECT id, alasan_kondisi_id, data_terukur, created_at, updated_at 
			FROM tb_data_terukur_pendukung_pernyataan WHERE alasan_kondisi_id = ?`

	rows, err := tx.QueryContext(ctx, SQL, alasanId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []domain.DataTerukurPendukungPernyataan
	for rows.Next() {
		data := domain.DataTerukurPendukungPernyataan{}
		err := rows.Scan(
			&data.Id,
			&data.AlasanKondisiId,
			&data.DataTerukur,
			&data.CreatedAt,
			&data.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

func (repository *CsfRepositoryImpl) UpdateAlasanKondisi(ctx context.Context, tx *sql.Tx, alasan *domain.AlasanKondisi) (*domain.AlasanKondisi, error) {
	SQL := `UPDATE tb_alasan_kondisi SET 
            alasan_kondisi_strategis = ?,
            updated_at = NOW()
            WHERE id = ? AND csf_id = ?`

	_, err := tx.ExecContext(ctx, SQL, alasan.AlasanKondisiStrategis, alasan.Id, alasan.CSFid)
	if err != nil {
		return nil, err
	}

	// Get updated data
	return repository.FindAlasanById(ctx, tx, alasan.Id)
}

func (repository *CsfRepositoryImpl) UpdateDataTerukur(ctx context.Context, tx *sql.Tx, dataTerukur *domain.DataTerukurPendukungPernyataan) (*domain.DataTerukurPendukungPernyataan, error) {
	SQL := `UPDATE tb_data_terukur_pendukung_pernyataan SET 
            data_terukur = ?,
            updated_at = NOW()
            WHERE id = ? AND alasan_kondisi_id = ?`

	_, err := tx.ExecContext(ctx, SQL, dataTerukur.DataTerukur, dataTerukur.Id, dataTerukur.AlasanKondisiId)
	if err != nil {
		return nil, err
	}

	// Get updated data
	return repository.FindDataTerukurById(ctx, tx, dataTerukur.Id)
}

func (repository *CsfRepositoryImpl) FindAlasanById(ctx context.Context, tx *sql.Tx, alasanId int) (*domain.AlasanKondisi, error) {
	SQL := `SELECT id, csf_id, alasan_kondisi_strategis, created_at, updated_at 
            FROM tb_alasan_kondisi WHERE id = ?`

	alasan := &domain.AlasanKondisi{}
	err := tx.QueryRowContext(ctx, SQL, alasanId).Scan(
		&alasan.Id,
		&alasan.CSFid,
		&alasan.AlasanKondisiStrategis,
		&alasan.CreatedAt,
		&alasan.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return alasan, nil
}

func (repository *CsfRepositoryImpl) FindDataTerukurById(ctx context.Context, tx *sql.Tx, dataTerukurId int) (*domain.DataTerukurPendukungPernyataan, error) {
	SQL := `SELECT id, alasan_kondisi_id, data_terukur, created_at, updated_at 
            FROM tb_data_terukur_pendukung_pernyataan WHERE id = ?`

	dataTerukur := &domain.DataTerukurPendukungPernyataan{}
	err := tx.QueryRowContext(ctx, SQL, dataTerukurId).Scan(
		&dataTerukur.Id,
		&dataTerukur.AlasanKondisiId,
		&dataTerukur.DataTerukur,
		&dataTerukur.CreatedAt,
		&dataTerukur.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return dataTerukur, nil
}
