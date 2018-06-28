package data

import (
	"log-converter/model"

	"gopkg.in/mgo.v2"
)

// CreateLogEntry stores log entry into a storage
func CreateLogEntry(e model.Entry) (model.Entry, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return model.Entry{}, err
	}
	defer session.Close()

	c := session.DB("log").C("logs")
	err = c.Insert(&e)
	if err != nil {
		return model.Entry{}, err
	}

	return e, nil
}
