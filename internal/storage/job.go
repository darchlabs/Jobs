package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/darchlabs/jobs/internal/job"
	"github.com/syndtr/goleveldb/leveldb/util"
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
func (j *Job) List(userID string) ([]*job.Job, error) {
	fmt.Printf("List: user %s", userID)

	data := make([]*job.Job, 0)

	var iter = j.storage.DB.NewIterator(nil, nil)
	if userID != "" {
		prefix := []byte(userID + ":")
		iter = j.storage.DB.NewIterator(util.BytesPrefix(prefix), nil)
	}

	for iter.Next() {
		key := iter.Key()
		keyParts := strings.Split(string(key), ":")
		if len(keyParts) == 2 {
			if userID != "" && keyParts[0] != userID {
				continue
			}

			var job *job.Job
			err := json.Unmarshal(iter.Value(), &job)
			if err != nil {
				return nil, err
			}

			data = append(data, job)
		}
	}
	iter.Release()

	err := iter.Error()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Method for getting job by its id
func (j *Job) GetById(id string, userID string) (*job.Job, error) {
	fmt.Println("ZOZOZOZOZO")
	fmt.Println("ZOZOZOZOZO", id)
	var iter = j.storage.DB.NewIterator(nil, nil)
	if userID != "" {
		prefix := []byte(userID + ":" + id)
		iter = j.storage.DB.NewIterator(util.BytesPrefix(prefix), nil)
	}

	for iter.Next() {
		key := iter.Key()
		fmt.Println("------------------")
		fmt.Println("iter.Next()", string(key))
		fmt.Println("------------------")
		keyParts := strings.Split(string(key), ":")
		fmt.Printf("%+v\n", keyParts)

		if len(keyParts) == 2 && keyParts[1] == id {
			var job *job.Job
			err := json.Unmarshal(iter.Value(), &job)
			if err != nil {
				return nil, err
			}

			if userID != "" && job.UserID != userID {
				return nil, errors.New("unauthorized: the provided userID does not have access to get this job")
			}

			return job, nil
		}
	}
	iter.Release()

	err := iter.Error()
	if err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("Job with id %s not found", id)
}

// Method for inserting job in the DB
func (j *Job) Insert(job *job.Job) (*job.Job, error) {
	b, err := json.Marshal(job)
	if err != nil {
		return nil, err
	}

	// save in database
	key := []byte(job.UserID + ":" + job.ID)
	err = j.storage.DB.Put(key, b, nil)
	if err != nil {
		return nil, err
	}

	return job, nil
}

// Method for updating job in the DB
func (j *Job) Update(newJob *job.Job, userID string) (*job.Job, error) {
	// check if userID is owner of job.id in db
	if userID != "" {
		job, err := j.GetById(newJob.ID, userID)
		if err != nil {
			return nil, err
		}

		if job.UserID != userID {
			return nil, errors.New("unauthorized: the provided userID does not have access to update this job")
		}
	}

	// update timestamp of job
	newJob.UpdatedAt = time.Now()

	b, err := json.Marshal(newJob)
	if err != nil {
		return nil, err
	}

	// save in database
	key := []byte(newJob.UserID + ":" + newJob.ID)
	err = j.storage.DB.Put(key, b, nil)
	if err != nil {
		return nil, err
	}

	return newJob, nil
}

// Method for deleting job from the DB
func (j *Job) Delete(id string, userID string) error {
	// check if userID is owner of job.id in db
	job, err := j.GetById(id, userID)
	if err != nil {
		return err
	}

	// check if userID is owner of job.id in db
	if userID != "" && job.UserID != userID {
		return errors.New("unauthorized: the provided userID does not have access to delete this job")
	}

	key := []byte(job.UserID + ":" + job.ID)
	err = j.storage.DB.Delete(key, nil)
	if err != nil {
		return err
	}

	return nil
}

// Method for stopping job
func (j *Job) Stop() error {
	return j.storage.DB.Close()
}
