package jobsapi

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/darchlabs/jobs/internal/api"
	"github.com/darchlabs/jobs/internal/job"
	"github.com/darchlabs/jobs/internal/provider"
	"github.com/darchlabs/jobs/internal/storage"
	"github.com/go-playground/validator"
	"github.com/robfig/cron"
	"github.com/teris-io/shortid"
)

type CreateJobsHandler struct {
	storage *storage.Job
}

func NewCreateJobsHandler(js *storage.Job) *CreateJobsHandler {
	return &CreateJobsHandler{
		storage: js,
	}
}

func (CreateJobsHandler) Invoke(ctx Context) *api.HandlerRes {

	fmt.Println("JEJE")
	fmt.Println("JEJE")
	fmt.Println("JEJE")
	fmt.Println("JEJE")
	fmt.Println("JEJE")
	// Prepare body request struct for parsing and validating
	body := struct {
		Job *job.Job `json:"job"`
	}{}

	// Parse body to Job struct
	err := json.Unmarshal(ctx.c.Body(), &body)
	if err != nil {
		return &api.HandlerRes{Payload: err.Error(), HttpStatus: 400, Err: err}
	}

	validate := validator.New()
	err = validate.Struct(ctx.JobStorage)
	if err != nil {
		return &api.HandlerRes{Payload: err.Error(), HttpStatus: 400, Err: err}
	}

	// Get user id
	userID, err := api.GetUserIDFromRequestCtx(ctx.c)
	if err != nil {
		return &api.HandlerRes{Payload: err.Error(), HttpStatus: 400, Err: err}
	}
	body.Job.UserID = userID

	// generate id for database
	id, err := shortid.Generate()
	if err != nil {
		return &api.HandlerRes{Payload: err.Error(), HttpStatus: 400, Err: err}
	}
	body.Job.ID = id

	if body.Job.Type == "cronjob" {
		// Validate the cronjob received is correct
		specParser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
		_, err = specParser.Parse(body.Job.Cronjob)

		if err != nil {
			fmt.Println(err)
			return &api.HandlerRes{Payload: api.ErrorInvalidCron, HttpStatus: 400, Err: err}
		}

		// Execute manager in order to execute the job
		err = ctx.Manager.Setup(body.Job)
		if err != nil {
			return &api.HandlerRes{Payload: err.Error(), HttpStatus: 400, Err: err}
		}
		ctx.Manager.Start(id)

	} else {
		err = fmt.Errorf("only 'cronjob' type is supported now, but received '%s' type", err)
		return &api.HandlerRes{Payload: err.Error(), HttpStatus: 400, Err: err}
	}

	body.Job.CreatedAt = time.Now()
	body.Job.Status = provider.StatusRunning

	// Insert job in jobstorage DB
	j, err := ctx.JobStorage.Insert(body.Job)
	if err != nil {
		fmt.Println(err)
		return &api.HandlerRes{Payload: err.Error(), HttpStatus: 400, Err: err}
	}

	return &api.HandlerRes{Payload: j, HttpStatus: 200, Err: nil}
}
