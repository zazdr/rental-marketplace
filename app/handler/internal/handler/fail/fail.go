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
	if fail, ok := errors.AsType[*dto.Fail](err); ok {
		if fail.Code < 500 {
			c.String(fail.Code, fail.Message)
			return
		} else {
			c.String(
				http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
			)
		}
	}

	c.Logger().Error(err.Error())
}
