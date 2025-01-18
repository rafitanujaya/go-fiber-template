package loggerZap

import functionCallerInfo "github.com/rafitanujaya/go-fiber-template/src/logger/helper"

type LoggerInterface interface {
	Info(msg string, function functionCallerInfo.FunctionCaller, data ...interface{})
	Error(msg string, function functionCallerInfo.FunctionCaller, data ...interface{})
	Debug(msg string, function functionCallerInfo.FunctionCaller, data ...interface{})
	Warn(msg string, function functionCallerInfo.FunctionCaller, data ...interface{})
}
