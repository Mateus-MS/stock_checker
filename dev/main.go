package main

import (
	"flag"
	"net/http"
	_ "placeholder/dev/backend/routes/pages"
	"placeholder/dev/features/app"
	"placeholder/dev/features/middlewares"
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
