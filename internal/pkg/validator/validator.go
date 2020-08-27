package validator

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/onepaas/onepaas/internal/pkg/db"
	"github.com/rs/zerolog/log"
)

var uniquenessValidator validator.Func = func(fl validator.FieldLevel) bool {
	fieldValue, ok := fl.Field().Interface().(string)
	if ok {
		// table name: paramas[0], column name: params[1]
		params := strings.Split(fl.Param(), ";")

		query := fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? LIMIT 1", params[1], params[0], params[1])
		res, err := db.GetDB().Exec(query, fieldValue)
		if err != nil {
			log.Error().Err(err)
			return false
		}

		if res.RowsReturned() == 1 {
			return false
		}
	}

	return true
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("uniqueness", uniquenessValidator)
	}
}
