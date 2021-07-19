package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/julienschmidt/httprouter"
	"github.com/meulengracht/golang-rest/controllers"
)

func main() {
	fmt.Println("Starting the REST API server up!")
	if runtime.GOOS == "windows" {
		fmt.Println("Server relies on linux for systemd-analyze and thus is not working properly for windows, sorry!")
		os.Exit(-1)
	}

	// instantiate the router, we use the httprouter package
	// as our base for this REST api server
	r := httprouter.New()

	// register all our callbacks
	r.GET("/version", controllers.GetServerVersion)
	r.GET("/duration", controllers.GetStartupTimingInfo)

	// start the server
	fmt.Println("Server ready, endpoints: /version, /duration")
	http.ListenAndServe("localhost:8082", r)
}
