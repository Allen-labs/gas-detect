package biz

import (
	"gas-detect/pkg/request"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type GasDetect struct {
	// ID of the ent.
	ID          int32     `json:"id"`
	CreatedTime time.Time `json:"created_time"`
	IsDeleted   bool      `json:"is_deleted"`
	// Metrics holds the value of the "metrics" field.
	Metrics []map[string]string `json:"metrics"`
}

type GasDetectRepo interface {
	// db
	ListGasDetect(ctx context.Context) ([]*GasDetect, error)
	PageGasDetect(ctx context.Context, pageParams *request.PageParams) ([]*GasDetect, error)
	GetGasDetect(ctx context.Context, id int32) (*GasDetect, error)
	CreateGasDetect(ctx context.Context, GasDetect *GasDetect) (*GasDetect, error)
	DeleteGasDetect(ctx context.Context, id int32) error
	SDeleteGasDetect(ctx context.Context, id int32) error
}

type GasDetectUsecase struct {
	repo GasDetectRepo
	log  *log.Helper
}

func NewGasDetectUsecase(repo GasDetectRepo, logger log.Logger) *GasDetectUsecase {
	return &GasDetectUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *GasDetectUsecase) List(ctx context.Context) (ps []*GasDetect, err error) {
	ps, err = uc.repo.ListGasDetect(ctx)
	if err != nil {
		return
	}
	return
}

func (uc *GasDetectUsecase) Page(ctx context.Context, pageParams *request.PageParams) (ps []*GasDetect, total int32, err error) {

	ps, err = uc.repo.PageGasDetect(ctx, pageParams)
	if err != nil {
		return nil, 0, errors.InternalServer("Gas Detect 获取失败", err.Error())
	}
	if err != nil {
		return
	}
	//total, err = uc.repo.TotalGasDetect(ctx)
	total = int32(len(ps))
	if err != nil {
		return
	}
	return
}

func (uc *GasDetectUsecase) Get(ctx context.Context, id int32) (p *GasDetect, err error) {
	p, err = uc.repo.GetGasDetect(ctx, id)
	if err != nil {
		return
	}

	return
}

func (uc *GasDetectUsecase) Create(ctx context.Context, GasDetect *GasDetect) (p *GasDetect, err error) {
	return uc.repo.CreateGasDetect(ctx, GasDetect)
}

func (uc *GasDetectUsecase) SDelete(ctx context.Context, id int32) error {
	return uc.repo.SDeleteGasDetect(ctx, id)
}

func (uc *GasDetectUsecase) Delete(ctx context.Context, id int32) error {
	return uc.repo.DeleteGasDetect(ctx, id)
}
