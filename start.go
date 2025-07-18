package kratosx

import (
	"context"
	"os"
	"path/filepath"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gyq14/kratosx/config"
	"github.com/gyq14/kratosx/library"
	"github.com/gyq14/kratosx/library/env"
	"github.com/gyq14/kratosx/library/logger"
	"github.com/gyq14/kratosx/library/pprof"
	"github.com/gyq14/kratosx/library/registry"
	"github.com/gyq14/kratosx/library/stop"
)

var (
	envName    string
	envVersion string
	id         string
)

func init() {
	env.Load()

	envName = env.GetAppName()
	envVersion = env.GetAppVersion()
	id, _ = os.Hostname()
}

func Init(opts ...Option) func() *options {
	o := &options{
		config: config.New(file.NewSource(filepath.Join(env.RootDir(), "internal/conf/conf.yaml"))),
	}

	for _, opt := range opts {
		opt(o)
	}

	// 加载配置
	if err := o.config.Load(); err != nil {
		panic(err)
	}

	// 初始化服务信息
	if o.config.App().Name == "" {
		o.config.SetAppInfo(id, envName, envVersion)
	}

	// 重置应用名称
	if envName == "" {
		env.SetAppName(o.config.App().Name)
		envName = o.config.App().Name
	}
	// 重置应用版本
	if envVersion == "" {
		env.SetAppVersion(o.config.App().Version)
		envVersion = o.config.App().Version
	}

	// 插件初始化
	if o.loggerFields == nil {
		o.loggerFields = logger.LogField{
			"id":      id,
			"name":    envName,
			"version": envVersion,
			"trace":   tracing.TraceID(),
			"span":    tracing.SpanID(),
		}
	}

	library.Init(o.config, o.loggerFields)

	return func() *options {
		return o
	}
}

func New(opts ...Option) *kratos.App {
	o := Init(opts...)()

	// 获取中间件
	defOpts := []kratos.Option{
		kratos.ID(o.config.App().ID),
		kratos.Name(o.config.App().Name),
		kratos.Version(o.config.App().Version),
		kratos.Metadata(map[string]string{}),
		kratos.BeforeStop(func(ctx2 context.Context) error {
			stop.Instance().WaitBefore()
			return nil
		}),
		kratos.AfterStop(func(ctx2 context.Context) error {
			stop.Instance().WaitAfter()
			return nil
		}),
	}

	// 必注册服务
	if o.regSrvFn != nil {
		gsOpts, hsOpts := serverOptions(o.config, o.midOpts)
		gsOpts = append(gsOpts, o.grpcSrvOptions...)
		hsOpts = append(hsOpts, o.httpSrvOptions...)

		srv := o.config.App().Server
		gs := grpcServer(srv.Grpc, srv.Count, gsOpts)
		hs := httpServer(srv.Http, srv.Count, hsOpts)
		o.regSrvFn(o.config, hs, gs)

		var srvList []transport.Server
		if srv.Http != nil {
			srvList = append(srvList, hs)
			// 监控
			if o.config.App().Metrics {
				hs.Handle("/metrics", promhttp.Handler())
			}
			// pprof
			if o.config.App().Server.Http.Pprof != nil {
				pprof.PprofServer(o.config.App().Server.Http.Pprof, hs)
			}
		}
		if srv.Grpc != nil {
			srvList = append(srvList, gs)
		}
		defOpts = append(defOpts, kratos.Server(srvList...))

		if srv.Registry != nil {
			reg, err := registry.Create(*srv.Registry)
			if err != nil {
				panic(err)
			}
			defOpts = append(defOpts, kratos.Registrar(reg))
		}
	}

	// 日志
	if o.config.App().Log != nil {
		defOpts = append(defOpts, kratos.Logger(logger.Instance()))
	}

	defOpts = append(defOpts, o.kOpts...)

	return kratos.New(
		defOpts...,
	)
}
