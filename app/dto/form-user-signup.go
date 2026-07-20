package dto

type FormUserSignup struct {
	Form struct {
		Mail           string `form:"mail"`
		MailCode       string `form:"mail_code"`
		Password       string `form:"password"`
		RepeatPassword string `form:"repeat_password"`
	}
	Name struct {
		Mail           string
		MailCode       string
		Password       string
		RepeatPassword string
	}
	Fail struct {
		Mail           string
		MailCode       string
		Password       string
		RepeatPassword string
	}
}

func (fus *FormUserSignup) Valid() bool {
	return fus.Fail.Mail == "" &&
		fus.Fail.MailCode == "" &&
		fus.Fail.Password == "" &&
		fus.Fail.RepeatPassword == ""
}

func FormUserSignupNew() *FormUserSignup {
	fus := &FormUserSignup{}

	fus.Name.Mail = "mail"
	fus.Name.MailCode = "mail_code"
	fus.Name.Password = "password"
	fus.Name.RepeatPassword = "repeat_password"

	return fus
}
