package main

func main() {
	var logObj Logger
	logObj.ToFile = false
	logObj.LogToConsole("Starting Server")

	var utility Utility
	
	var appDriver AppDriver
	appDriver.LogObj = logObj
	appDriver.UtilityObj = utility
	appDriver.AppInit()
	appDriver.AppDriverRun()
}
