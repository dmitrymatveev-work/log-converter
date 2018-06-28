package model

import "time"

// Entry is the type containing information about a log entry
type Entry struct {
	LogTime   time.Time `bson:"log_time"`
	LogMsg    string    `bson:"log_msg"`
	FileName  string    `bson:"file_name"`
	LogFormat string    `bson:"log_format"`
}
