package main

import (
	"net/http"
	"fmt"
	
	"github.com/labstack/echo"
)

type DownloaderServer struct {
	_echo *echo.Echo
	LogObj Logger
	UtilityObj Utility
	_downloader Downloader
}

func (server DownloaderServer) StartServer(port int) {
	server.LogObj.LogToConsole("Starting Download Server")
	server._downloader.LogObj = server.LogObj
	server._downloader.UtilityObj = server.UtilityObj
	server._echo = echo.New()
	server.initRoutes()
	server.startListening(port)
}

func (server DownloaderServer) initRoutes() {
	//e.POST("/normal", normal)
	//e.GET("/pathparam/:id", pathparam)
	//e.GET("/normal", normal)
	server._echo.GET("/check", server._checkMedia)
	server._echo.GET("/download", server._downloadMedia)
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)
}

func (server DownloaderServer) startListening(port int) {
	listenPort := fmt.Sprintf("%s%d", ":", port)
	server._echo.Logger.Fatal(server._echo.Start(listenPort))
}

func (server DownloaderServer) _checkMedia(c echo.Context) error {
	targetUrl := c.QueryParam("targetUrl")
	availableFormatsJson := server._downloader.CheckAvailableFormats(targetUrl)
	server.LogObj.LogToConsole("_checkMedia: JSON Response Received")
	return c.JSON(http.StatusOK, availableFormatsJson)
}

func (server DownloaderServer) _downloadMedia(c echo.Context) error {
	targetUrl := c.QueryParam("targetUrl")
	targetFormatCode := c.QueryParam("targetFormatCode")
	availableFormatsJson := server._downloader.BeginDownload(targetUrl, targetFormatCode)
	server.LogObj.LogToConsole("_downloadMedia: JSON Response Received")
	return c.JSON(http.StatusOK, availableFormatsJson)
}

/*
func normal(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World Nomral")
}

func pathparam(c echo.Context) error {
	id := c.Param("id")
	fmt.Println(id)
	return c.String(http.StatusOK, "Hello, World! Path Param")
}

func queryparam(c echo.Context) error {
	id := c.QueryParam("id")
	fmt.Println(id)
	return c.String(http.StatusOK, "Hello, World! Path Param")
}
*/