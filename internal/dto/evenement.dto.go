package dto

import (
	"time"

	"github.com/Cyber-cicco/jardin-pc/internal/utils"
)


type EvenementDto struct {
	Title       string    `form:"title"`
	Description *string   `form:"description"`
	Date        string `form:"date"`
}

func (e *EvenementDto) ParseDate() (time.Time, error) {
    return time.Parse(utils.DATE_TIME_LAYOUT, e.Date)
}
