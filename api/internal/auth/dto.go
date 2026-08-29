package auth

type UserRegister struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u UserRegister) Validate() error {
	if len(u.Username) < 6 {
		return ErrUsernameTooShort
	}

	if len(u.Password) < 8 {
		return ErrPasswordTooShort
	}

	if u.Password != u.RepeatPassword {
		return ErrPasswordsDontMatch
	}

	return nil
}
