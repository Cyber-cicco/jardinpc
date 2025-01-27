package validator

import (
	"strings"
)

type ErrorWrapper struct {
	Message string `json:"message"`
}

type Diagnostics struct {
	StatusCode int
    Errors map[string]string
}

type Validator interface {
    Validate()
}

//Returns a diagnostic object which aims to contain messages that 
//will be read by the client
func GetDiagnostics(statusCode int) *Diagnostics {
    return &Diagnostics{
        statusCode,
        map[string]string{},
    }
}
func (d *Diagnostics) PushIfBlank(val string, field string, message string) bool {
    if len(strings.TrimSpace(val)) == 0 {
        d.AppendError(field, message)
        return true
    }
    return false
}

func (d *Diagnostics) PushIfNullOrBlank(val *string, field string, message string) bool {
    if val == nil || len(strings.TrimSpace(*val)) == 0 {
        d.AppendError(field, message)
        return true
    }
    return false
}

func (d *Diagnostics) PushIfConditionIsTrue(condition bool, field string, message string) {
    if condition {
        d.AppendError(field, message)
    }
}

func (d *Diagnostics) PushIfLenAbove(length int, param *string, field string, message string) {
    if (param != nil && len(*param) > length) {
        d.AppendError(field, message)
    }
}

func (d *Diagnostics) AppendError(field string, msg string) {
    d.Errors[field] = msg
}

func (d *Diagnostics) IsNotEmpty() bool {
    return len(d.Errors) > 0
}

