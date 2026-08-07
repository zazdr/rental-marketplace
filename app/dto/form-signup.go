package dto

type FormSignup struct {
	Value struct {
		Mail           string `form:"mail"`
		MailCode       string `form:"mail-code"`
		Password       string `form:"password"`
		RepeatPassword string `form:"repeat-password"`
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
	fs.Name.MailCode = "mail-code"
	fs.Name.Password = "password"
	fs.Name.RepeatPassword = "repeat-password"

	return fs
}
