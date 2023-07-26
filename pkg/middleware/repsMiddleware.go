package middlewares

import (
	"gas-detect/pkg/request"
	"context"
	"fmt"
	_ "github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"log"
)

func RespMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		log.Println("resp middleware in", req)
		reply, err = handler(ctx, req)
		if err != nil {
			log.Println(err)
			err = errors.FromError(err)
			log.Println("convert to err resp", req)
			log.Println(err)
			return
		}
		fmt.Println("resp middleware out", reply)
		return
	}
}

func HeaderMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		//log.Println("header fields check", req)

		err = request.ValidateAimpHeader(ctx)
		if err != nil {
			log.Println("header fields 校验失败", err.Error())
			return nil, err
		}
		return handler(ctx, req)
	}
}
