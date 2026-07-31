package signup

import (
	"app/dto"
	"app/handler/internal/dep"
	"app/handler/internal/u"
	"app/shared"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
)

type Signup struct{}

var d *dep.Dep

func New(dep *dep.Dep) Signup {
	d = dep
	return Signup{}
}

func (s *Signup) Get(c *echo.Context) error {
	return u.RenderFull(d, c, d.UI.Signup(dto.FormSignupNew()))
}

func (s *Signup) Post(c *echo.Context) error {
	form := dto.FormSignupNew()

	if valid, err := valid(c, form); err != nil {
		return err
	} else if !valid {
		return u.RenderFull(d, c, d.UI.Signup(form))
	}

	return c.Redirect(http.StatusSeeOther, shared.RouterSignup)
}

func valid(
	c *echo.Context,
	form *dto.FormSignup,
) (bool, error) {
	v := &form.Value
	f := &form.Fail

	if err := c.Bind(v); err != nil {
		return false, u.FailClientNew(http.StatusBadRequest)
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
