package errors

type DuplicationError struct{}

func (err DuplicationError) Error() string {
	return "instance already exists"
}

type LoginFailed struct{}

func (err LoginFailed) Error() string {
	return "username or password is wrong"
}

type TokenInvalid struct{}

func (err TokenInvalid) Error() string {
	return "token invalid"
}
