package controller

import (
	"errors"

	"github.com/goadesign/goa"
	"github.com/imyslx/go-practice/goa/app"
	zlog "github.com/rs/zerolog/log"
)

// HostsController implements the hosts resource.
type HostsController struct {
	*goa.Controller
}

// NewHostsController creates a hosts controller.
func NewHostsController(service *goa.Service) *HostsController {
	return &HostsController{Controller: service.NewController("HostsController")}
}

// Add runs the add action.
func (c *HostsController) Add(ctx *app.AddHostsContext) error {
	///// Params
	// 	Param("hostname", String, "ホスト名")
	// 	Param("ipaddr", String, "IPアドレス", func() {
	// 		Default("")
	// 	})

	// HostsController_Add: start_implement
	if ctx.Payload.Hostname == "" {
		return ctx.BadRequest(errors.New("Hostname is empty.¥n"))
	}

	// Put your logic here
	zlog.Debug().Msg("Got Add Request¥n  Hostname: " + ctx.Payload.Hostname)

	// HostsController_Add: end_implement
	res := &app.Result{}
	res.Message = "Hostname: " + ctx.Payload.Hostname
	res.Status = 200

	return ctx.OK(res)
}

// Delete runs the delete action.
func (c *HostsController) Delete(ctx *app.DeleteHostsContext) error {
	///// Params
	// 	Param("hostname", String, "ホスト名")
	// 	Param("ipaddr", String, "IPアドレス", func() {
	// 		Default("")
	// 	})

	// HostsController_Add: start_implement
	if ctx.Payload.Hostname == "" {
		return ctx.BadRequest(errors.New("Hostname is empty.¥n"))
	}

	// Put your logic here
	zlog.Debug().Msg("Got Add Request¥n  Hostname: " + ctx.Payload.Hostname)

	// HostsController_Delete: end_implement
	res := &app.Result{}
	res.Message = "Hostname: " + ctx.Payload.Hostname
	res.Status = 200

	return ctx.OK(res)
}
