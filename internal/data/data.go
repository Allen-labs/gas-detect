package data

import (
	"gas-detect/internal/conf"
	"gas-detect/internal/data/model/ent"
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData,
	NewGasDetectRepo,
)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	drv, err := sql.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	db := drv.DB()
	db.SetMaxIdleConns(int(conf.Database.MaxIdleConnections))
	db.SetMaxOpenConns(int(conf.Database.MaxOpenConnections))
	db.SetConnMaxLifetime(conf.Database.ConnMaxLifeTime.AsDuration())
	client := ent.NewClient(ent.Driver(drv))
	if conf.Database.Debug {
		sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
			log.WithContext(ctx).Info(i...)
			tracer := otel.Tracer("ent.")
			kind := trace.SpanKindServer
			_, span := tracer.Start(ctx,
				"Query",
				trace.WithAttributes(
					attribute.String("sql", fmt.Sprint(i...)),
				),
				trace.WithSpanKind(kind),
			)
			span.End()
		})
		client = ent.NewClient(ent.Driver(sqlDrv))
	}

	if err != nil {
		log.Errorf("failed opening connection to sql: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	d := &Data{
		db: client,
	}
	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
