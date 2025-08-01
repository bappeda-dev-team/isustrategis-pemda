package service

import (
	"context"
	"database/sql"
	"isustrategisService/helper"
	"isustrategisService/model/domain"
	"isustrategisService/model/web"
	"isustrategisService/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type IntermediateServiceImpl struct {
	intermediateRepository repository.IntermediateRepository
	db                     *sql.DB
	validator              *validator.Validate
}

func NewIntermediateServiceImpl(intermediateRepository repository.IntermediateRepository, db *sql.DB, validator *validator.Validate) *IntermediateServiceImpl {
	return &IntermediateServiceImpl{
		intermediateRepository: intermediateRepository,
		db:                     db,
		validator:              validator,
	}
}

func (service *IntermediateServiceImpl) Create(ctx context.Context, request web.IntermediateCreateRequest) (*web.IntermediateResponse, error) {
	err := service.validator.Struct(request)
	if err != nil {
		return nil, err
	}

	tx, err := service.db.Begin()
	if err != nil {
		return nil, err
	}

	defer helper.CommitOrRollback(tx)

	intermediate := &domain.Intermediate{
		PohonId:       request.PohonId,
		Tahun:         request.Tahun,
		FaktorOutcome: request.FaktorOutcome,
		DataTerukur:   request.DataTerukur,
	}

	intermediate, err = service.intermediateRepository.Create(ctx, tx, intermediate)
	if err != nil {
		return nil, err
	}

	return &web.IntermediateResponse{
		Id:            intermediate.ID,
		PohonId:       intermediate.PohonId,
		Tahun:         intermediate.Tahun,
		FaktorOutcome: intermediate.FaktorOutcome,
		DataTerukur:   intermediate.DataTerukur,
		CreatedAt:     intermediate.CreatedAt.Format(time.DateTime),
		UpdatedAt:     intermediate.UpdatedAt.Format(time.DateTime),
	}, nil
}

func (service *IntermediateServiceImpl) Update(ctx context.Context, request web.IntermediateUpdateRequest) (*web.IntermediateResponse, error) {
	err := service.validator.Struct(request)
	if err != nil {
		return nil, err
	}

	tx, err := service.db.Begin()
	if err != nil {
		return nil, err
	}

	defer helper.CommitOrRollback(tx)

	intermediate, err := service.intermediateRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return nil, err
	}

	intermediate.PohonId = request.PohonId
	intermediate.Tahun = request.Tahun
	intermediate.FaktorOutcome = request.FaktorOutcome
	intermediate.DataTerukur = request.DataTerukur

	intermediate, err = service.intermediateRepository.Update(ctx, tx, intermediate)
	if err != nil {
		return nil, err
	}

	return &web.IntermediateResponse{
		Id:            intermediate.ID,
		PohonId:       intermediate.PohonId,
		Tahun:         intermediate.Tahun,
		FaktorOutcome: intermediate.FaktorOutcome,
		DataTerukur:   intermediate.DataTerukur,
		CreatedAt:     intermediate.CreatedAt.Format(time.DateTime),
		UpdatedAt:     intermediate.UpdatedAt.Format(time.DateTime),
	}, nil
}

func (service *IntermediateServiceImpl) Delete(ctx context.Context, pohonId int) error {
	tx, err := service.db.Begin()
	if err != nil {
		return err
	}

	defer helper.CommitOrRollback(tx)

	err = service.intermediateRepository.Delete(ctx, tx, pohonId)
	if err != nil {
		return err
	}

	return nil
}

func (service *IntermediateServiceImpl) FindById(ctx context.Context, id int) (*web.IntermediateResponse, error) {
	tx, err := service.db.Begin()
	if err != nil {
		return nil, err
	}

	defer helper.CommitOrRollback(tx)

	intermediate, err := service.intermediateRepository.FindById(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	return &web.IntermediateResponse{
		Id:            intermediate.ID,
		PohonId:       intermediate.PohonId,
		Tahun:         intermediate.Tahun,
		FaktorOutcome: intermediate.FaktorOutcome,
		DataTerukur:   intermediate.DataTerukur,
		CreatedAt:     intermediate.CreatedAt.Format(time.DateTime),
		UpdatedAt:     intermediate.UpdatedAt.Format(time.DateTime),
	}, nil
}

func (service *IntermediateServiceImpl) FindAll(ctx context.Context, tahun string) ([]web.IntermediateResponse, error) {
	tx, err := service.db.Begin()
	if err != nil {
		return nil, err
	}

	defer helper.CommitOrRollback(tx)

	intermediates, err := service.intermediateRepository.FindAll(ctx, tx, tahun)
	if err != nil {
		return nil, err
	}

	var intermediateResponses []web.IntermediateResponse
	for _, intermediate := range intermediates {
		intermediateResponses = append(intermediateResponses, web.IntermediateResponse{
			Id:            intermediate.ID,
			PohonId:       intermediate.PohonId,
			Tahun:         intermediate.Tahun,
			FaktorOutcome: intermediate.FaktorOutcome,
			DataTerukur:   intermediate.DataTerukur,
			CreatedAt:     intermediate.CreatedAt.Format(time.DateTime),
			UpdatedAt:     intermediate.UpdatedAt.Format(time.DateTime),
		})
	}

	return intermediateResponses, nil
}
