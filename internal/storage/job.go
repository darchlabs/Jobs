package storage

import (
	"encoding/json"
	"time"

	"github.com/darchlabs/jobs/internal/job"
)

type Job struct {
	storage *S
}

func NewJob(s *S) *Job {
	return &Job{
		storage: s,
	}
}

// Method for listing jobs array
func (j *Job) List() ([]*job.Job, error) {
	data := make([]*job.Job, 0)

	iter := j.storage.DB.NewIterator(nil, nil)
	for iter.Next() {
		var jj *job.Job
		err := json.Unmarshal(iter.Value(), &jj)
		if err != nil {
			return nil, err
		}

		data = append(data, jj)
	}
	iter.Release()

	err := iter.Error()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Method for getting job by its id
func (j *Job) GetById(id string) (*job.Job, error) {
	data, err := j.storage.DB.Get([]byte(id), nil)
	if err != nil {
		return nil, err
	}

	var job *job.Job
	err = json.Unmarshal(data, &job)
	if err != nil {
		return nil, err
	}

	return job, nil
}

// Method for inserting job in the DB
func (j *Job) Insert(job *job.Job) (*job.Job, error) {
	b, err := json.Marshal(job)
	if err != nil {
		return nil, err
	}

	// save in database
	err = j.storage.DB.Put([]byte(job.ID), b, nil)
	if err != nil {
		return nil, err
	}

	return job, nil
}

// Method for updating job in the DB
func (j *Job) Update(job *job.Job) (*job.Job, error) {
	job.UpdatedAt = time.Now()

	b, err := json.Marshal(job)
	if err != nil {
		return nil, err
	}

	// save in database
	err = j.storage.DB.Put([]byte(job.ID), b, nil)
	if err != nil {
		return nil, err
	}

	return job, nil
}

// Method for deleting job from the DB
func (j *Job) Delete(id string) error {
	err := j.storage.DB.Delete([]byte(id), nil)
	if err != nil {
		return err
	}

	return nil
}

// Method for stopping job
func (j *Job) Stop() error {
	return j.storage.DB.Close()
}
