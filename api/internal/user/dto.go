package user

type RegisterUser struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type UserCredentials struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (u RegisterUser) Validate() error {
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
