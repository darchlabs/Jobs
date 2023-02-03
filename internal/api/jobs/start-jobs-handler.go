package jobsapi

import (
	"fmt"
	"time"

	"github.com/darchlabs/jobs/internal/api"
	"github.com/darchlabs/jobs/internal/provider"
	"github.com/darchlabs/jobs/internal/storage"
)

type StartJobHandler struct {
	storage *storage.Job
}

func NewStartJobHandler(js *storage.Job) *StartJobHandler {
	return &StartJobHandler{
		storage: js,
	}
}

func (StartJobHandler) Invoke(ctx Context) *api.HandlerRes {
	// Get and check id
	id := ctx.c.Params("id")
	if id == "" {
		err := fmt.Errorf("%s", "id param in route is empty")
		return &api.HandlerRes{Payload: err.Error(), HttpStatus: 500, Err: err}
	}

	// Get job using id
	job, err := ctx.JobStorage.GetById(id)
	if err != nil {
		return &api.HandlerRes{Payload: err.Error(), HttpStatus: 500, Err: err}
	}

	// Validate the job to start is not already running
	if job.Status == provider.StatusRunning {
		err := fmt.Errorf("%s", "job is already running")
		return &api.HandlerRes{Payload: err.Error(), HttpStatus: 400, Err: err}
	}
	// Start job using the job manager
	ctx.Manager.Start(job.ID)

	// Update job values
	job.Status = provider.StatusRunning
	job.UpdatedAt = time.Now()
	job, err = ctx.JobStorage.Update(job)
	if err != nil {
		return &api.HandlerRes{Payload: err.Error(), HttpStatus: 500, Err: err}
	}

	res := map[string]interface{}{"id": job.ID, "status": job.Status}
	return &api.HandlerRes{Payload: res, HttpStatus: 200, Err: nil}
}
