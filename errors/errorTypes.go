package errors

// NoUserError when user not found
type NoUserError struct{}

func (e NoUserError) Error() string {
	return "No user found"
}

// AlreadyAddedError when trying to add
type AlreadyAddedError struct{}

func (e AlreadyAddedError) Error() string {
	return "User already added"
}

type UnparsableTokenError struct {
	Message string
}

func (e UnparsableTokenError) Error() string {
	return "Could not handle this token:" + e.Message
}

type NotActivatedOrExpiredTokenError struct {
	Message string
}

func (e NotActivatedOrExpiredTokenError) Error() string {
	return "Token is not yet activated or already expired:" + e.Message
}

type WrongTypeOfClaimsTokenError struct {
	Message string
}

func (e WrongTypeOfClaimsTokenError) Error() string {
	return "Token has wrong type of claims:" + e.Message
}
