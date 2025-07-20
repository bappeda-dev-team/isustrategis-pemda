package service

import (
	"context"
	"isustrategisService/model/web"
)

type CsfService interface {
	Create(ctx context.Context, request web.CsfCreateRequest) (*web.CsfResponse, error)
	Update(ctx context.Context, request web.CsfUpdateRequest) (*web.CsfResponse, error)
	Delete(ctx context.Context, idPohon int) error
	FindById(ctx context.Context, csfId int) (*web.CsfResponse, error)
	FindAll(ctx context.Context, tahun string) ([]web.CsfResponse, error)
}
