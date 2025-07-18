package middleware

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"

	ec "github.com/gyq14/kratosx/config"
	"github.com/gyq14/kratosx/library/logger"
	lg "github.com/gyq14/kratosx/library/logging"
)

func Logging(conf *ec.Logging) middleware.Middleware {
	if conf == nil || !conf.Enable {
		return nil
	}

	return selector.Server(logging.Server(logger.Instance(logger.AddCallerSkip(-1)))).Match(func(ctx context.Context, operation string) bool {
		path := ""
		if h, is := http.RequestFromServerContext(ctx); is {
			path = h.Method + ":" + h.URL.Path
		}
		lgIns := lg.Instance()
		return !(lgIns.IsWhitelist(operation) || lgIns.IsWhitelist(path))
	}).Build()
}
