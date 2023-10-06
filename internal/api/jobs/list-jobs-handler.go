package jobsapi

import (
	"github.com/darchlabs/jobs/internal/api"
	"github.com/darchlabs/jobs/internal/storage"
)

type ListJobsHandler struct {
	storage *storage.Job
}

func NewListJobsHandler(js *storage.Job) *ListJobsHandler {
	return &ListJobsHandler{
		storage: js,
	}
}

func (ListJobsHandler) Invoke(ctx Context) *api.HandlerRes {
	// get user id
	userID, err := api.GetUserIDFromRequestCtx(ctx.c)
	if err != nil {
		return &api.HandlerRes{Payload: nil, HttpStatus: 500, Err: err}
	}

	// Get elements from db
	data, err := ctx.JobStorage.List(userID)
	if err != nil {
		return &api.HandlerRes{Payload: nil, HttpStatus: 500, Err: err}
	}

	// Prepare response
	return &api.HandlerRes{Payload: data, HttpStatus: 200, Err: nil}
}
