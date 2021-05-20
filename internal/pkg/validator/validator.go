package validator

import (
	"database/sql"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/onepaas/onepaas/internal/pkg/database"
	"github.com/rs/zerolog/log"
)

var uniquenessValidator validator.Func = func(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()

	// table name: paramas[0], column name: params[1]
	params := strings.Split(fl.Param(), ";")

	row := database.GetDB().Table(params[0]).Where(params[1] + " = ?", fieldValue).Select(params[1]).Row()
	err := row.Scan()
	if err != nil {
		if err == sql.ErrNoRows {
			return true
		}

		log.Error().Err(err)
	}

	return false
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("uniqueness", uniquenessValidator)
	}
}
