package jobsapi

import (
	"github.com/darchlabs/jobs/internal/api"
	"github.com/darchlabs/jobs/internal/storage"
	"github.com/gofiber/fiber/v2"
)

type Context struct {
	JobStorage storage.Job
	c          *fiber.Ctx
}

// Define handler method
type handler func(ctx Context) *api.HandlerRes

func Route(app *fiber.App, ctx Context) {
	listJobsHandler := NewListJobsHandler(ctx.JobStorage)
	createJobsHandler := NewCreateJobsHandler(ctx.JobStorage)

	app.Get("/api/v1/jobs", HandleFunc(listJobsHandler.Invoke, ctx))
	app.Post("/api/v1/jobs", HandleFunc(createJobsHandler.Invoke, ctx))
}

// Func that receives the returns from handlers and creates an http response
func HandleFunc(fn handler, ctx Context) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// set headers
		c.Accepts("application/json")

		// Exec handler func and get its response
		ctx.c = c
		handlerRes := fn(ctx)
		payload, statusCode, err := handlerRes.Payload, handlerRes.HttpStatus, handlerRes.Err
		if err != nil {
			return c.Status(fiber.StatusConflict).JSON(api.Response{
				Error: err,
			})
		}

		// Prepare response
		res := api.Response{
			Meta: map[string]interface{}{"statusCode": statusCode},
			Data: payload,
		}

		// Return response
		return c.Status(fiber.StatusOK).JSON(res)
	}
}
