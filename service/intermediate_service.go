package service

import (
	"context"
	"isustrategisService/model/web"
)

type IntermediateService interface {
	Create(ctx context.Context, request web.IntermediateCreateRequest) (*web.IntermediateResponse, error)
	Update(ctx context.Context, request web.IntermediateUpdateRequest) (*web.IntermediateResponse, error)
	Delete(ctx context.Context, pohonId int) error
	FindById(ctx context.Context, id int) (*web.IntermediateResponse, error)
	FindAll(ctx context.Context, pohonId int) ([]web.IntermediateResponse, error)
}
