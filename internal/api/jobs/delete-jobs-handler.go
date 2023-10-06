package jobsapi

import (
	"fmt"

	"github.com/darchlabs/jobs/internal/api"
	"github.com/darchlabs/jobs/internal/provider"
	"github.com/darchlabs/jobs/internal/storage"
)

type DeleteJobHandler struct {
	storage *storage.Job
}

func NewDeleteJobHandler(js *storage.Job) *DeleteJobHandler {
	return &DeleteJobHandler{
		storage: js,
	}
}

func (DeleteJobHandler) Invoke(ctx Context) *api.HandlerRes {
	// Get id param and assert is not empty
	id := ctx.c.Params("id")
	if id == "" {
		return &api.HandlerRes{Payload: nil, HttpStatus: 500, Err: fmt.Errorf("%s", "id param in route is empty")}
	}

	// Get user id
	userID, err := api.GetUserIDFromRequestCtx(ctx.c)
	if err != nil {
		return &api.HandlerRes{Payload: nil, HttpStatus: 500, Err: err}
	}

	// Get and check that the job exists in DB
	job, err := ctx.JobStorage.GetById(id, userID)
	if err != nil {
		return &api.HandlerRes{Payload: nil, HttpStatus: 500, Err: err}
	}

	// Stop job if is running
	if job.Status == provider.StatusRunning {
		ctx.Manager.Stop(id)
	}

	// Delete job from jobstorage DB
	err = ctx.JobStorage.Delete(id, userID)
	if err != nil {
		return &api.HandlerRes{Payload: nil, HttpStatus: 500, Err: err}
	}

	res := map[string]interface{}{"id": job.ID, "status": "Deleted"}
	return &api.HandlerRes{Payload: res, HttpStatus: 200, Err: nil}
}
