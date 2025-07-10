package service

import (
	"context"
	"database/sql"
	"isustrategisService/helper"
	"isustrategisService/model/domain"
	"isustrategisService/model/web"
	"isustrategisService/repository"

	"github.com/go-playground/validator/v10"
)

type CsfServiceImpl struct {
	Repository repository.CsfRepository
	DB         *sql.DB
	Validator  *validator.Validate
}

func NewCsfServiceImpl(repository repository.CsfRepository, db *sql.DB, validator *validator.Validate) *CsfServiceImpl {
	return &CsfServiceImpl{
		Repository: repository,
		DB:         db,
		Validator:  validator,
	}
}

func (service *CsfServiceImpl) Create(ctx context.Context, request web.CsfCreateRequest) (*web.CsfResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	// Create CSF
	csf := &domain.Csf{
		PohonId:                    request.PohonId,
		PernyataanKondisiStrategis: request.PernyataanKondisiStrategis,
	}

	csf, err = service.Repository.Create(ctx, tx, csf)
	if err != nil {
		return nil, err
	}

	// Create Alasan Kondisi dan Data Terukur
	var alasanResponses []web.AlasanKondisiResponse
	for _, alasanRequest := range request.AlasanKondisi {
		// Create Alasan Kondisi
		alasan := &domain.AlasanKondisi{
			CSFid:                  csf.Id,
			AlasanKondisiStrategis: alasanRequest.AlasanKondisiStrategis,
		}

		alasan, err = service.Repository.CreateAlasanKondisi(ctx, tx, alasan)
		if err != nil {
			return nil, err
		}

		// Create Data Terukur untuk setiap Alasan
		var dataTerukurResponses []web.DataTerukurResponse
		for _, dataTerukurRequest := range alasanRequest.DataTerukurPendukungPernyataan {
			dataTerukur := &domain.DataTerukurPendukungPernyataan{
				AlasanKondisiId: alasan.Id,
				DataTerukur:     dataTerukurRequest.DataTerukur,
			}

			dataTerukur, err = service.Repository.CreateDataTerukur(ctx, tx, dataTerukur)
			if err != nil {
				return nil, err
			}

			dataTerukurResponses = append(dataTerukurResponses, web.DataTerukurResponse{
				Id:              dataTerukur.Id,
				AlasanKondisiId: dataTerukur.AlasanKondisiId,
				DataTerukur:     dataTerukur.DataTerukur,
				CreatedAt:       dataTerukur.CreatedAt,
				UpdatedAt:       dataTerukur.UpdatedAt,
			})
		}

		alasanResponses = append(alasanResponses, web.AlasanKondisiResponse{
			Id:                             alasan.Id,
			CSFid:                          alasan.CSFid,
			AlasanKondisiStrategis:         alasan.AlasanKondisiStrategis,
			DataTerukurPendukungPernyataan: dataTerukurResponses,
			CreatedAt:                      alasan.CreatedAt,
			UpdatedAt:                      alasan.UpdatedAt,
		})
	}

	return &web.CsfResponse{
		Id:                         csf.Id,
		PohonId:                    csf.PohonId,
		PernyataanKondisiStrategis: csf.PernyataanKondisiStrategis,
		AlasanKondisi:              alasanResponses,
		CreatedAt:                  csf.CreatedAt,
		UpdatedAt:                  csf.UpdatedAt,
	}, nil
}

func (service *CsfServiceImpl) Update(ctx context.Context, request web.CsfUpdateRequest) (*web.CsfResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	// Validasi CSF exists
	_, err = service.Repository.FindById(ctx, tx, request.Id)
	if err != nil {
		return nil, err
	}

	// Update CSF
	csf := &domain.Csf{
		Id:                         request.Id,
		PohonId:                    request.PohonId,
		PernyataanKondisiStrategis: request.PernyataanKondisiStrategis,
		AlasanKondisi:              make([]domain.AlasanKondisi, 0),
	}

	// Convert alasan kondisi
	for _, alasanRequest := range request.AlasanKondisi {
		alasan := domain.AlasanKondisi{
			Id:                             alasanRequest.Id,
			CSFid:                          request.Id,
			AlasanKondisiStrategis:         alasanRequest.AlasanKondisiStrategis,
			DataTerukurPendukungPernyataan: make([]domain.DataTerukurPendukungPernyataan, 0),
		}

		// Convert data terukur
		for _, dtRequest := range alasanRequest.DataTerukurPendukungPernyataan {
			dataTerukur := domain.DataTerukurPendukungPernyataan{
				Id:              dtRequest.Id,
				AlasanKondisiId: alasanRequest.Id,
				DataTerukur:     dtRequest.DataTerukur,
			}
			alasan.DataTerukurPendukungPernyataan = append(alasan.DataTerukurPendukungPernyataan, dataTerukur)
		}

		csf.AlasanKondisi = append(csf.AlasanKondisi, alasan)
	}

	// Lakukan update melalui repository
	updatedCsf, err := service.Repository.Update(ctx, tx, csf)
	if err != nil {
		return nil, err
	}

	// Convert ke response
	var alasanResponses []web.AlasanKondisiResponse
	for _, alasan := range updatedCsf.AlasanKondisi {
		var dataTerukurResponses []web.DataTerukurResponse
		for _, dt := range alasan.DataTerukurPendukungPernyataan {
			dataTerukurResponses = append(dataTerukurResponses, web.DataTerukurResponse{
				Id:              dt.Id,
				AlasanKondisiId: dt.AlasanKondisiId,
				DataTerukur:     dt.DataTerukur,
				CreatedAt:       dt.CreatedAt,
				UpdatedAt:       dt.UpdatedAt,
			})
		}

		alasanResponses = append(alasanResponses, web.AlasanKondisiResponse{
			Id:                             alasan.Id,
			CSFid:                          alasan.CSFid,
			AlasanKondisiStrategis:         alasan.AlasanKondisiStrategis,
			DataTerukurPendukungPernyataan: dataTerukurResponses,
			CreatedAt:                      alasan.CreatedAt,
			UpdatedAt:                      alasan.UpdatedAt,
		})
	}

	return &web.CsfResponse{
		Id:                         updatedCsf.Id,
		PohonId:                    updatedCsf.PohonId,
		PernyataanKondisiStrategis: updatedCsf.PernyataanKondisiStrategis,
		AlasanKondisi:              alasanResponses,
		CreatedAt:                  updatedCsf.CreatedAt,
		UpdatedAt:                  updatedCsf.UpdatedAt,
	}, nil
}

func (service *CsfServiceImpl) Delete(ctx context.Context, csfId int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	// Ambil semua alasan kondisi
	alasanList, err := service.Repository.FindAlasanByCsfId(ctx, tx, csfId)
	if err != nil {
		return err
	}

	// Hapus data terukur untuk setiap alasan
	for _, alasan := range alasanList {
		dataTerukurList, err := service.Repository.FindDataTerukurByAlasanId(ctx, tx, alasan.Id)
		if err != nil {
			return err
		}

		for _, dataTerukur := range dataTerukurList {
			err = service.Repository.DeleteDataTerukur(ctx, tx, dataTerukur.Id)
			if err != nil {
				return err
			}
		}

		// Hapus alasan kondisi
		err = service.Repository.DeleteAlasanKondisi(ctx, tx, alasan.Id)
		if err != nil {
			return err
		}
	}

	// Hapus CSF
	return service.Repository.Delete(ctx, tx, csfId)
}

func (service *CsfServiceImpl) FindById(ctx context.Context, csfId int) (*web.CsfResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	csf, err := service.Repository.FindById(ctx, tx, csfId)
	if err != nil {
		return nil, err
	}

	// Ambil alasan kondisi
	alasanList, err := service.Repository.FindAlasanByCsfId(ctx, tx, csfId)
	if err != nil {
		return nil, err
	}

	var alasanResponses []web.AlasanKondisiResponse
	for _, alasan := range alasanList {
		// Ambil data terukur untuk setiap alasan
		dataTerukurList, err := service.Repository.FindDataTerukurByAlasanId(ctx, tx, alasan.Id)
		if err != nil {
			return nil, err
		}

		var dataTerukurResponses []web.DataTerukurResponse
		for _, dataTerukur := range dataTerukurList {
			dataTerukurResponses = append(dataTerukurResponses, web.DataTerukurResponse{
				Id:              dataTerukur.Id,
				AlasanKondisiId: dataTerukur.AlasanKondisiId,
				DataTerukur:     dataTerukur.DataTerukur,
				CreatedAt:       dataTerukur.CreatedAt,
				UpdatedAt:       dataTerukur.UpdatedAt,
			})
		}

		alasanResponses = append(alasanResponses, web.AlasanKondisiResponse{
			Id:                             alasan.Id,
			CSFid:                          alasan.CSFid,
			AlasanKondisiStrategis:         alasan.AlasanKondisiStrategis,
			DataTerukurPendukungPernyataan: dataTerukurResponses,
			CreatedAt:                      alasan.CreatedAt,
			UpdatedAt:                      alasan.UpdatedAt,
		})
	}

	return &web.CsfResponse{
		Id:                         csf.Id,
		PohonId:                    csf.PohonId,
		PernyataanKondisiStrategis: csf.PernyataanKondisiStrategis,
		AlasanKondisi:              alasanResponses,
		CreatedAt:                  csf.CreatedAt,
		UpdatedAt:                  csf.UpdatedAt,
	}, nil
}

func (service *CsfServiceImpl) FindAll(ctx context.Context, pohonId int) ([]web.CsfResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	csfs, err := service.Repository.FindAll(ctx, tx, pohonId)
	if err != nil {
		return nil, err
	}

	var csfResponses []web.CsfResponse
	for _, csf := range csfs {
		// Ambil alasan kondisi untuk setiap CSF
		alasanList, err := service.Repository.FindAlasanByCsfId(ctx, tx, csf.Id)
		if err != nil {
			return nil, err
		}

		var alasanResponses []web.AlasanKondisiResponse
		for _, alasan := range alasanList {
			// Ambil data terukur untuk setiap alasan
			dataTerukurList, err := service.Repository.FindDataTerukurByAlasanId(ctx, tx, alasan.Id)
			if err != nil {
				return nil, err
			}

			var dataTerukurResponses []web.DataTerukurResponse
			for _, dataTerukur := range dataTerukurList {
				dataTerukurResponses = append(dataTerukurResponses, web.DataTerukurResponse{
					Id:              dataTerukur.Id,
					AlasanKondisiId: dataTerukur.AlasanKondisiId,
					DataTerukur:     dataTerukur.DataTerukur,
					CreatedAt:       dataTerukur.CreatedAt,
					UpdatedAt:       dataTerukur.UpdatedAt,
				})
			}

			alasanResponses = append(alasanResponses, web.AlasanKondisiResponse{
				Id:                             alasan.Id,
				CSFid:                          alasan.CSFid,
				AlasanKondisiStrategis:         alasan.AlasanKondisiStrategis,
				DataTerukurPendukungPernyataan: dataTerukurResponses,
				CreatedAt:                      alasan.CreatedAt,
				UpdatedAt:                      alasan.UpdatedAt,
			})
		}

		csfResponses = append(csfResponses, web.CsfResponse{
			Id:                         csf.Id,
			PohonId:                    csf.PohonId,
			PernyataanKondisiStrategis: csf.PernyataanKondisiStrategis,
			AlasanKondisi:              alasanResponses,
			CreatedAt:                  csf.CreatedAt,
			UpdatedAt:                  csf.UpdatedAt,
		})
	}

	return csfResponses, nil
}
