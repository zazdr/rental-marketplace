package signup

import (
	"app/dto"
	"app/handler/internal/dep"

	"github.com/labstack/echo/v5"
)

type Signup struct {
	d *dep.Dep
}

func New(dep *dep.Dep) Signup {
	return Signup{
		d: dep,
	}
}

func (s *Signup) Get(c *echo.Context) error {
	return s.d.Util.Render(
		c,
		s.d.UI.User.Signup(s.d.Util.State(c), dto.FormUserSignupNew()),
	)
}
