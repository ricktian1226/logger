// @APIVersion 1.0.0
// @Title logger
// @Description 服务基类定义，用于定义与业务无关的服务基础属性
// @Contact tianguimao@treehousefuture.com
package common

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"os"
	"runtime"
)

var uclogger = NewUCLogger()

type UCLogger struct {
}

func NewUCLogger() *UCLogger {
	return &UCLogger{}
}

func (this *UCLogger) init() (err error) {
	//不管日志目录有没有，尝试创建下
	path, configs := beego.AppConfig.String("log::path"), beego.AppConfig.String("log::configs")
	stdout, err := beego.AppConfig.Bool("log::stdout")
	if err != nil {
		panic(err)
	}

	os.MkdirAll(path, os.ModePerm)
	//日志文件全路径名称
	fileName := fmt.Sprintf("%s/%s.log", path, beego.BConfig.AppName)
	//设置日志文件配置信息。按照日志等级拆分
	if err = beego.SetLogger(logs.AdapterMultiFile, fmt.Sprintf(configs, fileName)); err != nil {
		panic(err)
	}

	//如果禁止控制台输出，则删除默认的stdout日志记录器
	if !stdout {
		if err = beego.BeeLogger.DelLogger(logs.AdapterConsole); err != nil {
			panic(err)
		}
	}

	//todo 增加错误日志的通知slack实现
	//if err := logs.SetLogger(logs.AdapterSlack, `{"webhookurl":"https://marcox.slack.com/messages/D8MPY8672/","level":1}`); err != nil {
	//	panic(err)
	//}
	return
}

var CUR_FUNC_NAME = func(level int) string {
	if funcName, _, _, ok := runtime.Caller(level); ok {
		return runtime.FuncForPC(funcName).Name()
	}
	return ""
}

var CUR_FUNC_LINE = func(level int) string {
	if _, file, line, ok := runtime.Caller(level); ok {
		return fmt.Sprintf("%s:%d", file, line)
	}
	return ""
}

func LOG_INIT() error {
	return uclogger.init()
}

//严重错误日志
func LOG_FUNC_CRITICAL(format string, args ...interface{}) {
	beego.Critical(formatWithFileline(format, args...))
}

//错误日志
func LOG_FUNC_ERROR(format string, args ...interface{}) {
	beego.Error(formatWithFileline(format, args...))
}

//告警日志
func LOG_FUNC_WARNING(format string, args ...interface{}) {
	beego.Warning(formatWithFileline(format, args...))
}

//跟踪日志
func LOG_FUNC_TRACE(format string, args ...interface{}) {
	beego.Trace(formatWithFileline(format, args...))
}

//信息日志
func LOG_FUNC_INFO(format string, args ...interface{}) {
	beego.Info(formatWithFileline(format, args...))
}

//调试日志
func LOG_FUNC_DEBUG(format string, args ...interface{}) {
	beego.Debug(formatWithFileline(format, args...))
}

func formatWithFileline(format string, args ...interface{}) string {
	fileline := CUR_FUNC_LINE(3)
	message := fmt.Sprintf(format, args...)
	return "[" + fileline + "] " + message
}

func log_func_slack() {

}
