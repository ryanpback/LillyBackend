package helpers

// LogInfo logs info messages
func LogInfo(message interface{}) {
	AppConfig.Logger.Info(message)
}

// LogError logs error messages
func LogError(message interface{}) {
	AppConfig.Logger.Error(message)
}

// LogFatal logs fatal messages
func LogFatal(message interface{}) {
	AppConfig.Logger.Fatal(message)
}
