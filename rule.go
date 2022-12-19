package vd

import (
	"strings"
)

type Rule struct {
	Fail    bool
	Message string
	Path    []string
	Format  Formatter
}

func (r *Rule) Break(message string, path string) {
	r.Fail = true
	r.Message = message
	r.Path = append(r.Path, path)
}

func (r *Rule) Validator(v interface {
	Validator(err ...error) error
}, failMessage string, path string) {
	if r.Fail {
		return
	}
	err := v.Validator()
	if err != nil {
		r.Break(failMessage, path)
	}
}

func (r Rule) CreateMessage(message string, customMessage func() string) string {
	message = strings.TrimPrefix(message, " ")
	message = strings.TrimSuffix(message, " ")
	if len(message) == 0 {
		return customMessage()
	}
	return message
}
