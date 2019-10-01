package main

type AppDriver struct {
	ConfigFilePath	string
	LogObj Logger
	UtilityObj Utility
	_server	DownloaderServer
}

func (app AppDriver) AppInit() {
	// check config
	// app.ConfigFilePath
	app.LogObj.LogToConsole("Application Initialized")
	app._server.LogObj = app.LogObj
	app._server.UtilityObj = app.UtilityObj
}

func (app AppDriver) AppDriverRun() {
	app.LogObj.LogToConsole("Application Driver Started")
	app._server.StartServer(1323)
}