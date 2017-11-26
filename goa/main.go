//go:generate goagen bootstrap -d github.com/imyslx/go-practice/goa/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/imyslx/go-practice/goa/app"
	"github.com/imyslx/go-practice/goa/controller"
)

func main() {
	// Create service
	service := goa.New("goa Practice")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "hosts" controller
	c := controller.NewHostsController(service)
	app.MountHostsController(service, c)
	// Mount "swagger" controller
	c2 := controller.NewSwaggerController(service)
	app.MountSwaggerController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError("startup", "err", err)
	}

}
