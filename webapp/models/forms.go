package models

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type RegistrationForm struct {
	Username  string `form:"username"  binding:"required"`
	Email     string `form:"email"     binding:"required"`
	PasswordA string `form:"passworda" binding:"required"`
	PasswordB string `form:"passwordb" binding:"required"`
}

func (r *RegistrationForm) CheckPasswords() bool {
	return r.PasswordA == r.PasswordB
}

type EditProfile struct {
	Username string `form:"username"  binding:"required"`
	AboutMe  string `form:"about_me"  binding:"required"`
}
