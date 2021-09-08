package db

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"sync"
	"time"
)

var (
	ErrFileNotFound = errors.New("error file not found")
	ErrDataBase     = errors.New("error database error")
)

type DataStore struct {
	fileLocation string
	mx           *sync.Mutex
}

func New(fileLocation string) (*DataStore, error) {
	if _, err := ioutil.ReadFile(fileLocation); err != nil {

		// if there is not file, create
		jsonBytes, err := json.Marshal([]string{})
		if err != nil {
			return nil, ErrDataBase
		}
		if err = ioutil.WriteFile(fileLocation, jsonBytes, 0644); err != nil {
			return nil, ErrDataBase
		}
	}

	datastore := DataStore{
		fileLocation: fileLocation,
		mx:           &sync.Mutex{},
	}

	return &datastore, nil
}

func (ds *DataStore) Counter(secs int) (int, error) {
	// set lock
	ds.mx.Lock()
	defer ds.mx.Unlock()

	// read file
	bytes, err := ioutil.ReadFile(ds.fileLocation)
	if err != nil {
		return 0, ErrFileNotFound
	}

	// deserialize to slice of times
	var times []time.Time
	err = json.Unmarshal(bytes, &times)
	if err != nil {
		return 0, ErrDataBase
	}

	// filter old times because they are no longer relevant. Storing will cause latency for reads
	var filteredTimes []time.Time
	for _, t := range times {
		checkPoint := time.Now().Add(time.Second * time.Duration(-secs))
		if t.After(checkPoint) {
			filteredTimes = append(filteredTimes, t)
		}
	}

	// add the time of the current request
	filteredTimes = append(filteredTimes, time.Now())
	bytes, err = json.Marshal(filteredTimes)
	if err != nil {
		return 0, ErrDataBase
	}

	// overwrite the file
	err = ioutil.WriteFile(ds.fileLocation, bytes, 0644)
	if err != nil {
		return 0, ErrDataBase
	}

	// return the count of requests for the last period
	return len(filteredTimes), nil
}
