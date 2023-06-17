package logs

var logs *Logger

func init() {
	logs = &Logger{}
	logs.createLogger()
}

func SetupLogger() {
	logs.createLogger()
}
