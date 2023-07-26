package server

import (
	gas "gas-detect/api/gas"
	"gas-detect/internal/conf"
	"gas-detect/internal/service"
	middlewares "gas-detect/pkg/middleware"
	aimpResponse "gas-detect/pkg/response"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/validate"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger,
	gd *service.GasDetectService,
) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			//tracing.Server(),
			logging.Server(logger),
			validate.Validator(),
			//middlewares.RespMiddleware,
			middlewares.HeaderMiddleware,
		),
		http.ResponseEncoder(aimpResponse.ResponseEncoder),
		http.ErrorEncoder(aimpResponse.ErrorEncoder),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	gas.RegisterGasDetectServiceHTTPServer(srv, gd)

	h := openapiv2.NewHandler()
	srv.HandlePrefix("", h)
	return srv
}
