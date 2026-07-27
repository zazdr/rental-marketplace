package fail

import (
	"app/handler/internal/dep"
	"app/handler/internal/u"
	"net/http"

	"github.com/labstack/echo/v5"
)

type Fail struct{}

var d *dep.Dep

func New(dep *dep.Dep) Fail {
	d = dep
	return Fail{}
}

func (f *Fail) Valid(c *echo.Context, err error) {
	code := http.StatusInternalServerError
	message := http.StatusText(http.StatusInternalServerError)

	if sc := echo.StatusCode(err); 0 < sc && sc < 500 {
		code = sc
		message = http.StatusText(sc)
	}

	if code > 499 {
		c.Logger().Error(err.Error())
	}

	u.Render(c, d.UI.Layout(d.UI.Fail(code, message)))
}
