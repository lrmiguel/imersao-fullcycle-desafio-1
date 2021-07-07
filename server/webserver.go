package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type WebServer struct {
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()
	e.Static("/images", "images")
	e.GET("/home", w.getAll)
	e.Logger.Fatal(e.Start(":8000"))
}

func (w WebServer) getAll(c echo.Context) error {
	return c.HTML(http.StatusOK,
		"<html>"+
			"\n    <head>\n        "+
			"<title>Desafio 1 Webserver com Docker e Golang</title>\n    "+
			"<style>"+
			"body {"+
			"margin: 0;"+
			"}"+
			"div {"+
			"width: 100%;"+
			"height: 100%;"+
			"display: block;"+
			"position: relative;"+
			"border: 0;"+
			"}"+
			"div:after {"+
			"content: \"\";"+
			"background: url(\"images/marvel.jpg\");"+
			"opacity: 0.5;"+
			"top: 0;"+
			"left: 0;"+
			"bottom: 0;"+
			"right: 0;"+
			"position: absolute;"+
			"z-index: -1;"+
			"}"+
			"h1 {"+
			"font-family: sans-serif;"+
			"width: 100%;"+
			"height: 100%;"+
			"color: #DB2C52;"+
			"display: flex;"+
			"align-items: center;"+
			"justify-content: center;"+
			"font-size: 10em;"+
			"}"+
			"</style>"+
			"</head>\n    "+
			"<body>\n        "+
			"<div><h1>Imersão Full Cycle</h1></div>\n    "+
			"</body>\n"+
			"</html>")
}
