package dto

type FormSignup struct {
	Value struct {
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

func (fs *FormSignup) Valid() bool {
	return fs.Fail.Mail == "" &&
		fs.Fail.MailCode == "" &&
		fs.Fail.Password == "" &&
		fs.Fail.RepeatPassword == ""
}

func FormSignupNew() *FormSignup {
	fs := &FormSignup{}

	fs.Name.Mail = "mail"
	fs.Name.MailCode = "mail_code"
	fs.Name.Password = "password"
	fs.Name.RepeatPassword = "repeat_password"

	return fs
}
