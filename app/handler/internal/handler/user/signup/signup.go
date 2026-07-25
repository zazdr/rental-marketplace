package signup

import (
	"app/dto"
	"app/handler/internal/dep"
	"app/shared"
	"net/http"
	"strings"

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

func (s *Signup) Post(c *echo.Context) error {
	form := dto.FormUserSignupNew()

	if valid, err := s.valid(c, form); err != nil {
		return err
	} else if !valid {
		return s.d.Util.Render(c, s.d.UI.User.Signup(s.d.Util.State(c), form))
	}

	return c.Redirect(http.StatusSeeOther, shared.RouterUserSignup)
}

func (s *Signup) valid(
	c *echo.Context,
	form *dto.FormUserSignup,
) (bool, error) {
	v := &form.Value
	f := &form.Fail

	if err := c.Bind(v); err != nil {
		return false, dto.FailClientNew(http.StatusBadRequest)
	}

	if strings.TrimSpace(v.Mail) == "" {
		f.Mail = "field cannot be blank"
	}

	if strings.TrimSpace(v.MailCode) == "" {
		f.MailCode = "field cannot be blank"
	}

	if strings.TrimSpace(v.Password) == "" {
		f.Password = "field cannot be blank"
	}

	if strings.TrimSpace(v.RepeatPassword) == "" {
		f.RepeatPassword = "field cannot be blank"
	}

	return form.Valid(), nil
}
