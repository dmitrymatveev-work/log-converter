package model

import "time"

// Entry is the type containing information about a log entry
type Entry struct {
	LogTime   time.Time `json:"log_time"`
	LogMsg    string    `json:"log_msg"`
	FileName  string    `json:"file_name"`
	LogFormat string    `json:"log_format"`
}
