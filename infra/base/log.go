package base

import (
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

func init() {
	// 定义日志的格式
	formatter := &prefixed.TextFormatter{}
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02 15:04-05.000000"
	formatter.ForceFormatting = true
	//formatter.SetColorScheme(&prefixed.ColorScheme{
	//	InfoLevelStyle:  "",
	//	WarnLevelStyle:  "",
	//	ErrorLevelStyle: "",
	//	FatalLevelStyle: "",
	//	PanicLevelStyle: "",
	//	DebugLevelStyle: "",
	//	PrefixStyle:     "",
	//	TimestampStyle:  "",
	//})
	log.SetFormatter(formatter)
	// 日志级别
	level := os.Getenv("log.debug")
	if level == "true" {
		log.SetLevel(log.DebugLevel)
	}
	// 控制台高亮显示
	formatter.ForceColors = true
	formatter.DisableColors = false
	// 日志文件和滚动配置
	// todo "github.com/lestrrat/go-file-rotatelogs"
}
