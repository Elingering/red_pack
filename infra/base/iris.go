package base

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	irisrecover "github.com/kataras/iris/middleware/recover"
	"github.com/sirupsen/logrus"
	"go-demo/infra"
	"net/http"
	"time"
)

var irisApplication *iris.Application

func Iris() *iris.Application {
	return irisApplication
}

type IrisStarter struct {
	infra.BaseStarter
}

func (i *IrisStarter) Init(ctx infra.StarterContext) {
	// 创建iris application实例
	irisApplication = initIris()
	// 日志组件配置和扩展
	logx := irisApplication.Logger()
	logx.Install(logrus.StandardLogger())
}

func (i *IrisStarter) Start(ctx infra.StarterContext) {
	// 把路由信息打印到控制台
	routes := Iris().GetRoutes()
	for _, r := range routes {
		logrus.Info(r.Trace())
	}
	// todo 统一处理 HTTP 错误码
	Iris().OnAnyErrorCode(func(context iris.Context) {
		context.WriteString("something go wrong")
	})
	Iris().OnErrorCode(http.StatusNotFound, func(context iris.Context) {
		context.WriteString("oh hoo find nothing")
	})
	// 启动 iris
	port := ctx.Props().GetDefault("app.server.port", "80")
	Iris().Run(iris.Addr(":" + port))
}

func (i *IrisStarter) StartBlocking() bool {
	return true
}

func initIris() *iris.Application {
	app := iris.New()
	app.Use(irisrecover.New())
	// 主要中间件配置：recover， 日志输出中间件的定义
	cfg := logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
		Query:  true,
		//Columns:            false,
		//MessageContextKeys: nil,
		//MessageHeaderKeys:  nil,
		LogFunc: func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
			app.Logger().Infof("| %s | %s | %s | %s | %s | %s | %s | %s |", now.Format("2006-01-02 15:04:05.000000"), latency.String(), status, ip, method, path, headerMessage, message)
		},
		//Skippers:           nil,
	}
	app.Use(logger.New(cfg))
	return app
}
