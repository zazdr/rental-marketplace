package fail

import (
	"app/dto"
	"app/handler/internal/dep"
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
)

type Fail struct {
	d *dep.Dep
}

func New(dep *dep.Dep) Fail {
	return Fail{
		d: dep,
	}
}

func (f *Fail) Valid(c *echo.Context, err error) {
	var code int
	var message string

	if fail, ok := errors.AsType[*dto.Fail](err); ok {
		code = fail.Code
		if code < 500 {
			message = fail.Message
		} else {
			message = http.StatusText(fail.Code)
		}
	} else {
		code = http.StatusInternalServerError
		message = http.StatusText(http.StatusInternalServerError)
	}

	if code > 499 {
		c.Logger().Error(err.Error())
	}

	f.d.Util.Render(c, f.d.UI.Fail(f.d.Util.State(c), code, message))
}
