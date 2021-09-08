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
	ErrDeserializingData = errors.New("error deserializing data")
)

type DataStore struct {
	fileLocation string
	mx *sync.Mutex
}

func (ds *DataStore) Counter(now time.Time) (int, error){
	ds.mx.Lock()
	defer ds.mx.Unlock()

	bs, err := ioutil.ReadFile(ds.fileLocation)
	if err != nil {
		return 0, ErrFileNotFound
	}

	var arr []string

	err = json.Unmarshal(bs, &arr)
	if err != nil {
		0, ErrDeserializingData
	}

	for _, timeStr := range arr {
		time timeStr
	}
}