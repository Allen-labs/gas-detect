package log

import (
	"gas-detect/internal/conf"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spf13/cast"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

var FileRotateLogs = new(fileRotateLogs)

type fileRotateLogs struct{}

func (r *fileRotateLogs) GetWriteSyncer(config *conf.Zap, level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotateLogs.New(
		path.Join(config.Director, "gas-detect."+"%Y-%m-%d"+".json"),
		rotateLogs.WithClock(rotateLogs.Local),
		rotateLogs.WithMaxAge(time.Duration(cast.ToInt(config.MaxAge))*24*time.Hour), // 日志留存时间
		rotateLogs.WithRotationTime(time.Hour*24),
	)
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
}
