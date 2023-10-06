package jobsapi

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/darchlabs/jobs/internal/api"
	"github.com/darchlabs/jobs/internal/provider"
	"github.com/darchlabs/jobs/internal/storage"
	"github.com/go-playground/validator"
)

type UpdateBody struct {
	Name         string `json:"name"`
	Network      string `json:"network"`
	NodeURL      string `json:"nodeUrl"`
	Privatekey   string `json:"privateKey"`
	Address      string `json:"address"`
	Abi          string `json:"abi"`
	CheckMethod  string `json:"checkMethod"`
	ActionMethod string `json:"cctionMethod"`
	Cronjob      string `json:"cronjob"`
}

type UpdateJobHandler struct {
	storage *storage.Job
}

func NewUpdateJobHandler(js *storage.Job) *UpdateJobHandler {
	return &UpdateJobHandler{
		storage: js,
	}
}

func (UpdateJobHandler) Invoke(ctx Context) *api.HandlerRes {
	// Get and check id
	id := ctx.c.Params("id")
	if id == "" {
		err := fmt.Errorf("%s", "id param in route is empty")
		return &api.HandlerRes{Payload: nil, HttpStatus: 500, Err: err}
	}

	// Prepare body request struct for parsing and validating
	body := struct {
		Req *UpdateBody `json:"job"`
	}{}

	// Parse body to Job struct
	err := json.Unmarshal(ctx.c.Body(), &body)
	if err != nil {
		return &api.HandlerRes{Payload: nil, HttpStatus: 500, Err: err}
	}

	// Validate the job values to update are correct
	validate := validator.New()
	err = validate.Struct(ctx.JobStorage)
	if err != nil {
		return &api.HandlerRes{Payload: nil, HttpStatus: 500, Err: err}
	}

	// Get user id
	userID, err := api.GetUserIDFromRequestCtx(ctx.c)
	if err != nil {
		return &api.HandlerRes{Payload: nil, HttpStatus: 500, Err: err}
	}

	// Get j ob using id
	job, err := ctx.JobStorage.GetById(id, userID)
	if err != nil {
		return &api.HandlerRes{Payload: nil, HttpStatus: 500, Err: err}
	}

	// Check the inputs of the job and return the job with the updated values if there is no errors
	job, err = ValidateInputs(job, body.Req)
	if err != nil {
		return &api.HandlerRes{Payload: nil, HttpStatus: 500, Err: err}
	}

	// Stop running job for updating the running cron
	if job.Status == provider.StatusRunning {
		fmt.Println("Status running so stopping...")
		ctx.Manager.Stop(job.ID)
	}

	// Setup new job with the values received
	fmt.Println("Setting up ...")
	err = ctx.Manager.Setup(job)
	if err != nil {
		// This is for keep running the previous job that was ok, in case the job setup failed
		ctx.Manager.Start(job.ID)
		return &api.HandlerRes{Payload: nil, HttpStatus: 500, Err: err}
	}
	fmt.Println("Setted up!")

	// Start job
	fmt.Println("Starting ...")
	ctx.Manager.Start(job.ID)
	fmt.Println("Started!")

	// Set job values and update them in the DB
	fmt.Println("Updating...")
	job.UpdatedAt = time.Now()
	job.Status = provider.StatusRunning
	job, err = ctx.JobStorage.Update(job, userID)
	if err != nil {
		return &api.HandlerRes{Payload: nil, HttpStatus: 500, Err: err}
	}
	fmt.Println("Updated!")

	res := map[string]interface{}{"job": job}
	return &api.HandlerRes{Payload: res, HttpStatus: 200, Err: nil}
}
