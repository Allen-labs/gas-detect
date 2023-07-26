package data

import (
	"gas-detect/internal/data/model/ent"
	"gas-detect/internal/data/model/ent/gasdetect"
	"gas-detect/internal/data/model/ent/predicate"
	"gas-detect/pkg/request"
	utils "gas-detect/pkg/utils"
	"context"
	"time"

	"gas-detect/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type gasDetectRepo struct {
	data *Data
	log  *log.Helper
}

func NewGasDetectRepo(data *Data, logger log.Logger) biz.GasDetectRepo {
	return &gasDetectRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (gd *gasDetectRepo) PageGasDetect(ctx context.Context, pageParams *request.PageParams) ([]*biz.GasDetect, error) {
	var ps []*ent.GasDetect
	var err error
	querys := make([]predicate.GasDetect, 0)
	querys = append(querys, gasdetect.IsDeleted(false))

	psQuery := gd.data.db.GasDetect.Query().Where(querys...).Offset(int(pageParams.Offset)).Limit(int(pageParams.PageSize))

	switch pageParams.Sorter.Order {
	case request.ORDER_ASC:
		ps, err = psQuery.Order(ent.Asc(pageParams.Sorter.Field)).All(ctx)
	case request.ORDER_DESC:
		ps, err = psQuery.Order(ent.Desc(pageParams.Sorter.Field)).All(ctx)
	}
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.GasDetect, 0)
	for _, p := range ps {
		gas := new(biz.GasDetect)
		utils.StructCopy(gas, p)
		rv = append(rv, gas)
	}
	return rv, nil
}

func (gd *gasDetectRepo) ListGasDetect(ctx context.Context) ([]*biz.GasDetect, error) {
	ps, err := gd.data.db.GasDetect.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.GasDetect, 0)
	for _, p := range ps {
		gasDetect := new(biz.GasDetect)
		utils.StructCopy(gasDetect, p)
		rv = append(rv, gasDetect)
	}
	return rv, nil
}

func (gd *gasDetectRepo) GetGasDetect(ctx context.Context, id int32) (*biz.GasDetect, error) {
	p, err := gd.getIsNotDeletedGasDetect(ctx, id)
	if err != nil {
		return nil, err
	}
	gasDetect := new(biz.GasDetect)
	if p != nil {
		utils.StructCopy(gasDetect, p)
	}

	return gasDetect, nil
}

func (gd *gasDetectRepo) getIsNotDeletedGasDetect(ctx context.Context, id int32) (*ent.GasDetect, error) {
	p, err := gd.data.db.GasDetect.Query().Where(gasdetect.IsDeleted(false),
		gasdetect.ID(id),
	).First(ctx)

	return p, err
}

func (gd *gasDetectRepo) CreateGasDetect(ctx context.Context, gasDetect *biz.GasDetect) (*biz.GasDetect, error) {
	p, err := gd.data.db.GasDetect.
		Create().
		SetCreatedTime(time.Now()).
		SetMetrics(gasDetect.Metrics).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	//tx, err := gd.data.db.Tx(ctx)
	//tx.GasDetect.Create()
	rgd := new(biz.GasDetect)
	if p != nil {
		utils.StructCopy(rgd, p)
	}

	return rgd, err
}

func (gd *gasDetectRepo) DeleteGasDetect(ctx context.Context, id int32) error {
	return gd.data.db.GasDetect.DeleteOneID(id).Exec(ctx)
}

func (gd *gasDetectRepo) SDeleteGasDetect(ctx context.Context, id int32) error {
	p, err := gd.getIsNotDeletedGasDetect(ctx, id)
	if err != nil {
		return err
	}
	_, err = p.Update().
		SetIsDeleted(true).
		Save(ctx)
	return err
}

func (gd *gasDetectRepo) TotalGasDetect(ctx context.Context) (int32, error) {
	total, err := gd.data.db.GasDetect.Query().Count(ctx)
	if err != nil {
		return 0, err
	}
	return int32(total), nil
}
