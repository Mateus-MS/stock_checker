package main

import (
	"flag"
	"net/http"

	_ "github.com/Mateus-MS/stock_checker/dev/backend/routes/api/spreadsheet"
	"github.com/Mateus-MS/stock_checker/dev/features/app"
	"github.com/Mateus-MS/stock_checker/dev/features/middlewares"
)

func main() {
	enviroment := flag.String("env", "dev", "The enviroment to run")
	flag.Parse()

	app := app.GetInstance()

	app.Router.Use(middlewares.CorsMiddleware(app.Router.Routes))

	startServer(app.Router, *enviroment)
}

func startServer(router *app.Router, env string) {
	if env == "dev" {
		println("Starting SERVER in DEV mode")
		go func() {
			if err := http.ListenAndServe("localhost:3000", router.Handle()); err != nil {
				println("HTTP server error:", err)
			}
		}()
		select {}
	}
}
