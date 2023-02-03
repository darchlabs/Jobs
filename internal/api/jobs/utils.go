package jobsapi

import (
	"fmt"

	"github.com/darchlabs/jobs/internal/api"
	"github.com/darchlabs/jobs/internal/job"
	"github.com/robfig/cron"
)

// Method for validating request params are ok and not completely empty
func ValidateInputs(job *job.Job, reqBody *UpdateBody) (*job.Job, error) {
	var errorLog error

	// Check that at least there is one param to be modified in the request
	empty := true
	// Update jobs params that were sent in the request
	if reqBody.Name != "" && reqBody.Name != job.Name {
		job.Name = reqBody.Name
		empty = false
	}

	if reqBody.Address != "" && reqBody.Address != job.Address {
		job.Address = reqBody.Address
		empty = false
	}

	if reqBody.Abi != "" && reqBody.Abi != job.Abi {
		job.Abi = reqBody.Abi
		empty = false
	}

	if reqBody.CheckMethod != "" && reqBody.CheckMethod != *job.CheckMethod {
		job.CheckMethod = &reqBody.CheckMethod
		empty = false
	}

	if reqBody.ActionMethod != "" && reqBody.ActionMethod != job.ActionMethod {
		job.ActionMethod = reqBody.ActionMethod
		empty = false
	}

	if reqBody.Network != "" && reqBody.Network != job.Network {
		job.Network = reqBody.Network
		empty = false
	}

	if reqBody.NodeURL != "" && reqBody.NodeURL != job.NodeURL {
		job.NodeURL = reqBody.NodeURL
		empty = false
	}

	if reqBody.Privatekey != "" && reqBody.Privatekey != job.Privatekey {
		job.Privatekey = reqBody.Privatekey
		empty = false
	}

	if reqBody.Cronjob != "" && reqBody.Cronjob != job.Cronjob {
		// Validate the cronjob received is correct
		specParser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
		_, err := specParser.Parse(reqBody.Cronjob)
		if err != nil {
			errorLog = fmt.Errorf("%s", api.ErrorInvalidCron)
			return nil, errorLog
		}

		// If it is ok, update the value
		job.Cronjob = reqBody.Cronjob
		empty = false

	}

	// If empty is true, it means that no params will be updated so an error is returned
	if empty {
		errorLog = fmt.Errorf("%s", "All params are empty or the same than the actual job")
		return nil, errorLog
	}

	// Return the updated values of the job based on the update request received
	return job, nil
}
