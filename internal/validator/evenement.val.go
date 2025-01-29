package validator

import (
	"net/http"
	"time"

	"github.com/Cyber-cicco/jardin-pc/internal/dto"
)

func ValidateEvenement(evt *dto.EvenementDto) *Diagnostics {
    diags := GetDiagnostics(http.StatusBadRequest)
    t, err := evt.ParseDate()
    if err != nil {
        diags.AppendError("date", "Invalid date format")
        return diags
    }
    diags.PushIfBlank(evt.Title, "title",  "Le titre ne peut pas être vide")
    diags.PushIfLenAbove(60, &evt.Title, "title", "Le titre ne peut faire plus de 60 caractères")
    diags.PushIfLenAbove(512, evt.Description, "description", "La description ne peut faire plus de 512 caractères")
    diags.PushIfConditionIsTrue(time.Now().After(t), "date", "La date doit nécessairement être dans le futur")
    return diags
}
