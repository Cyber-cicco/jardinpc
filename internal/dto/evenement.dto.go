package dto

import (
	"time"
)


type EvenementDto struct {
	Title       string    `form:"title"`
	Description *string   `form:"description"`
    Date        time.Time `form:"date" time_format:"2006-01-02T15:04"`
}

