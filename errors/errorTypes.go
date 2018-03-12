package errors

// NoUserError when user not found
type NoUserError struct {
	s string
}

func (e NoUserError) Error() string {
	return "No user found " + e.s
}
