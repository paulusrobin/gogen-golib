package mandatory

// User object of mandatory User.
type User struct {
	login bool
	id    uint64
	email string
	phone string
}

// ID getter function of mandatory user ID.
func (u User) ID() uint64 {
	return u.id
}

// Email getter function of mandatory user Email.
func (u User) Email() string {
	return u.email
}

// Phone getter function of mandatory user Phone.
func (u User) Phone() string {
	return u.phone
}

// IsLogin getter function of mandatory user IsLogin.
func (u User) IsLogin() bool {
	return u.login
}
