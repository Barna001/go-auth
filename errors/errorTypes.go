package errors

// NoUserError when user not found
type NoUserError struct {
	s string
}

func (e NoUserError) Error() string {
	return "No user found"
}

// AlreadyAddedError when trying to add
type AlreadyAddedError struct {
	s string
}

func (e AlreadyAddedError) Error() string {
	return "User already added"
}
