package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key string
	Msg string
}

func (v *ValidError) Error() string {
	return v.Msg
}

type ValidErrors []*ValidError

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var validErrs ValidErrors
	err := c.ShouldBind(v)
	if err != nil {
		trans, _ := c.Value("trans").(ut.Translator)
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, validErrs
		}
		for key, value := range errs.Translate(trans) {
			validErrs = append(validErrs, &ValidError{
				Key: key,
				Msg: value,
			})
		}
		return false, validErrs
	}

	return true, nil
}
