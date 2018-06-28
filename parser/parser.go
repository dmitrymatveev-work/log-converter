package parser

import (
	"fmt"
	"log-converter/model"
	"strings"
	"time"
)

const longForm = "Jan 2, 2006 at 3:04:05pm (UTC)"

// Parser is dedicated to parse log entries
type Parser struct {
	parseDate func(s string) (time.Time, error)
	logFormat string
}

// New creates new instance of Parser
func New(parseDate func(s string) (time.Time, error), logFormat string) *Parser {
	return &Parser{
		parseDate: parseDate,
		logFormat: logFormat,
	}
}

// Parse parses log entry
func (p *Parser) Parse(str string, fileName string) (model.Entry, error) {
	strs := strings.Split(str, "|")
	if len(strs) != 2 {
		return model.Entry{}, fmt.Errorf("Incorrect log string format: %s", str)
	}
	t, err := p.parseDate(strings.Trim(strs[0], " \t\n\r"))
	if err != nil {
		return model.Entry{}, err
	}

	return model.Entry{
		FileName:  fileName,
		LogFormat: p.logFormat,
		LogMsg:    strings.Trim(strs[1], " \t\n\r"),
		LogTime:   t,
	}, nil
}

// GetParseDate returns date parsing function
func GetParseDate(logFormat string) func(s string) (time.Time, error) {
	switch logFormat {
	case model.FirstFormat:
		return parseFirstDate
	case model.SecondFormat:
		return parseSecondDate
	}
	panic(fmt.Sprintf("Unsupported log format: %s.\n", logFormat))
}

func parseFirstDate(s string) (time.Time, error) {
	return time.Parse(longForm, s)
}

func parseSecondDate(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}
