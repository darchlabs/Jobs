package client

import "github.com/darchlabs/jobs/internal/job"

type ListJobsResponse struct {
	Data []*job.Job `json:"jobs"`
}
