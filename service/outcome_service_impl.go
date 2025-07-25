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

type OutcomeServiceImpl struct {
	outcomeRepository repository.OutcomeRepository
	db                *sql.DB
	validator         *validator.Validate
}

func NewOutcomeServiceImpl(outcomeRepository repository.OutcomeRepository, db *sql.DB, validator *validator.Validate) *OutcomeServiceImpl {
	return &OutcomeServiceImpl{
		outcomeRepository: outcomeRepository,
		db:                db,
		validator:         validator,
	}
}

func (service *OutcomeServiceImpl) Create(ctx context.Context, request web.OutcomeCreateRequest) (*web.OutcomeResponse, error) {
	err := service.validator.Struct(request)
	if err != nil {
		return nil, err
	}

	tx, err := service.db.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	outcome := &domain.Outcome{
		PohonId:       request.PohonId,
		FaktorOutcome: helper.EmptyStringIfNull(request.FaktorOutcome),
		DataTerukur:   helper.EmptyStringIfNull(request.DataTerukur),
		Tahun:         request.Tahun,
		ParentId:      request.ParentId,
	}

	outcome, err = service.outcomeRepository.Create(ctx, tx, outcome)
	if err != nil {
		return nil, err
	}

	return &web.OutcomeResponse{
		PohonId:       outcome.PohonId,
		Tahun:         outcome.Tahun,
		FaktorOutcome: outcome.FaktorOutcome,
		DataTerukur:   outcome.DataTerukur,
		CreatedAt:     outcome.CreatedAt.Format(time.DateTime),
		UpdatedAt:     outcome.UpdatedAt.Format(time.DateTime),
		ParentId:      outcome.ParentId,
	}, nil
}

func (service *OutcomeServiceImpl) Update(ctx context.Context, request web.OutcomeUpdateRequest) (*web.OutcomeResponse, error) {
	tx, err := service.db.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	outcome, err := service.outcomeRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return nil, err
	}

	outcome.PohonId = request.PohonId
	outcome.Tahun = request.Tahun
	outcome.FaktorOutcome = helper.EmptyStringIfNull(request.FaktorOutcome)
	outcome.DataTerukur = helper.EmptyStringIfNull(request.DataTerukur)

	outcome, err = service.outcomeRepository.Update(ctx, tx, outcome)
	if err != nil {
		return nil, err
	}

	return &web.OutcomeResponse{
		PohonId:       outcome.PohonId,
		Tahun:         outcome.Tahun,
		FaktorOutcome: outcome.FaktorOutcome,
		DataTerukur:   outcome.DataTerukur,
		CreatedAt:     outcome.CreatedAt.Format(time.DateTime),
		UpdatedAt:     outcome.UpdatedAt.Format(time.DateTime),
	}, nil
}

func (service *OutcomeServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.db.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	return service.outcomeRepository.Delete(ctx, tx, id)
}

func (service *OutcomeServiceImpl) FindById(ctx context.Context, id int) (*web.OutcomeResponse, error) {
	tx, err := service.db.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	outcome, err := service.outcomeRepository.FindById(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	return &web.OutcomeResponse{
		Id:            outcome.ID,
		PohonId:       outcome.PohonId,
		Tahun:         outcome.Tahun,
		FaktorOutcome: outcome.FaktorOutcome,
		DataTerukur:   outcome.DataTerukur,
		CreatedAt:     outcome.CreatedAt.Format(time.DateTime),
		UpdatedAt:     outcome.UpdatedAt.Format(time.DateTime),
		ParentId:      outcome.ParentId,
	}, nil
}

func (service *OutcomeServiceImpl) FindAll(ctx context.Context, tahun string) ([]web.OutcomeResponse, error) {
	tx, err := service.db.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	outcomes, err := service.outcomeRepository.FindAll(ctx, tx, tahun)
	if err != nil {
		return nil, err
	}

	var outcomeResponses []web.OutcomeResponse
	for _, outcome := range outcomes {
		outcomeResponses = append(outcomeResponses, web.OutcomeResponse{
			Id:            outcome.ID,
			PohonId:       outcome.PohonId,
			ParentId:      outcome.ParentId,
			Tahun:         outcome.Tahun,
			FaktorOutcome: outcome.FaktorOutcome,
			DataTerukur:   outcome.DataTerukur,
			CreatedAt:     outcome.CreatedAt.Format(time.DateTime),
			UpdatedAt:     outcome.UpdatedAt.Format(time.DateTime),
		})
	}

	return outcomeResponses, nil
}
