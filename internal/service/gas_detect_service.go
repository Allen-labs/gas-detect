package service

import (
	"context"
	pb "gas-detect/api/gas"
	"gas-detect/internal/biz"
	"gas-detect/pkg/request"
	util "gas-detect/pkg/utils"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GasDetectService struct {
	pb.UnimplementedGasDetectServiceServer
	log *log.Helper
	gd  *biz.GasDetectUsecase
}

func NewGasDetectService(gd *biz.GasDetectUsecase, logger log.Logger) *GasDetectService {
	return &GasDetectService{
		gd:  gd,
		log: log.NewHelper(logger),
	}
}

func (s *GasDetectService) CreateGasDetect(ctx context.Context, req *pb.CreateGasDetectRequest) (*pb.GasDetectResponse, error) {
	if req.Body.Metrics == nil {
		return nil, errors.BadRequest("metrics 不能为空", "")
	}
	metrics := util.MetricsToMap(req.Body.Metrics)

	gasDetect, err := s.gd.Create(ctx, &biz.GasDetect{
		Metrics: metrics,
	})
	if err != nil {
		return nil, errors.InternalServer("gas detect db 创建失败", err.Error())
	}

	return &pb.GasDetectResponse{
		Id:          gasDetect.ID,
		Metrics:     req.Body.Metrics,
		CreatedTime: gasDetect.CreatedTime.String(),
	}, nil
}

func (s *GasDetectService) GetGasDetect(ctx context.Context, req *pb.GetGasDetectRequest) (*pb.GasDetectResponse, error) {
	if req.Id <= 0 {
		return nil, errors.BadRequest("id 参数有误", "")
	}

	gasDetect, err := s.gd.Get(ctx, req.Id)
	if err != nil {
		return nil, errors.InternalServer("gas detect 获取失败", err.Error())
	}
	metrics := util.MapToMetrics(gasDetect.Metrics)
	//metrics, err = s.analysisGas(metrics)
	//if err != nil {
	//	s.log.Warn("gas detect analysis 分析失败", err.Error())
	//	//return nil, errors.InternalServer("gas detect analysis 分析失败", err.Error())
	//}
	return &pb.GasDetectResponse{
		Id:          gasDetect.ID,
		Metrics:     metrics,
		CreatedTime: gasDetect.CreatedTime.String(),
	}, nil
}

func (s *GasDetectService) DeleteGasDetect(ctx context.Context, req *pb.DeleteGasDetectRequest) (*emptypb.Empty, error) {
	if req.Id <= 0 {
		return nil, errors.BadRequest("id 参数有误", "")
	}
	_, err := s.gd.Get(ctx, req.Id)
	if err != nil {
		return nil, errors.BadRequest("gas detect 不存在", err.Error())
	}

	err = s.gd.SDelete(ctx, req.Id)
	if err != nil {
		return nil, errors.InternalServer("gas detect delete 失败", err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *GasDetectService) ListGasDetect(ctx context.Context, req *pb.ListGasDetectRequest) (*pb.ListGasDetectResponse, error) {
	pageParams, err := request.PaginateValidate(req.GetPageNo(), req.GetPageSize(), req.GetSorter())
	if err != nil {
		return nil, errors.BadRequest("查询参数有误", err.Error())
	}

	gasDetects, total, err := s.gd.Page(ctx, pageParams)
	if err != nil {
		return nil, errors.InternalServer("gas detect 获取失败", err.Error())
	}

	content := make([]*pb.GasDetectResponse, 0)
	for _, gasDetect := range gasDetects {
		metrics := util.MapToMetrics(gasDetect.Metrics)
		//metrics, err = s.analysisGas(metrics)
		//if err != nil {
		//	s.log.Warn("gas detect analysis 分析失败", err.Error())
		//	//return nil, errors.InternalServer("gas detect analysis 分析失败", err.Error())
		//}
		content = append(content, &pb.GasDetectResponse{
			Id:          gasDetect.ID,
			Metrics:     metrics,
			CreatedTime: gasDetect.CreatedTime.String(),
		})
	}

	return &pb.ListGasDetectResponse{
		Total:   total,
		Content: content,
	}, err
}

func (s *GasDetectService) analysisGas(metrics []*pb.Metric) ([]*pb.Metric, error) {
	for _, metric := range metrics {
		concentration := metric.Concentration
		concentrationInt64, err := util.StrToFloat64(concentration)
		if err != nil {
			return nil, errors.InternalServer(err.Error(), "")
		}
		switch metric.TypeName {
		case "co2":
			if concentrationInt64 < 0.2 {
				metric.AnalysisResult = "normal"
			} else {
				metric.AnalysisResult = "dangerous"
			}
		}
	}
	return metrics, nil

}
