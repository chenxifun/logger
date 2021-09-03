package logger

import "time"

type LogContent interface {

	// GetTime 格式化之后的日志输出时间
	GetTime() string

	// GetLevel 日志级别
	GetLevel() string

	// GetPath 输出日志的代码路径
	GetPath() string

	// GetName 应用名称
	GetName() string

	// GetContent 日志正文
	GetContent() string

	// GetExtendProp 扩展值
	GetExtendProp() interface{}
}

func (l *loginfo) GetTime() string {
	return l.Time
}

func (l *loginfo) GetLevel() string {
	return l.Level
}
func (l *loginfo) GetPath() string {
	return l.Path
}
func (l *loginfo) GetName() string {
	return l.Name
}
func (l *loginfo) GetContent() string {
	return l.Content
}

func (l *loginfo) GetExtendProp() interface{} {
	return l.ExtendProp
}

type LogMsgConvert func(t time.Time, l LogContent) string

func defaultLogConvert(t time.Time, msg LogContent) string {

	msgStr := t.Format(logTimeDefaultFormat) + " [" + msg.GetLevel() + "] " + "[" + msg.GetPath() + "] " + msg.GetContent()
	return msgStr
}

func (this *LocalLogger) SetLogConvert(c LogMsgConvert) {
	this.logConvert = c
}

func (this *LocalLogger) SetAppName(name string) {
	this.appName = name
}
func (this *LocalLogger) SetTimeFormat(timeFormat string) {
	this.timeFormat = timeFormat
}
