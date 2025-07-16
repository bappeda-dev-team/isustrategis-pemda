package service

import (
	"context"
	"isustrategisService/model/web"
)

type OutcomeService interface {
	Create(ctx context.Context, request web.OutcomeCreateRequest) (*web.OutcomeResponse, error)
	Update(ctx context.Context, request web.OutcomeUpdateRequest) (*web.OutcomeResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (*web.OutcomeResponse, error)
	FindAll(ctx context.Context, pohonId int) ([]web.OutcomeResponse, error)
}
