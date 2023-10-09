package errors

type DuplicationError struct{}

func (err DuplicationError) Error() string {
	return "Instance already exists"
}
