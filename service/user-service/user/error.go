package user

type NotFoundError struct {
	message string
}

func (e *NotFoundError) Error() string { return e.message }

type InternalError struct {
	message string
}

func (e *InternalError) Error() string { return e.message }

type InvalidArgument struct {
	message string
}

func (e *InvalidArgument) Error() string { return e.message }

type PreconditionAlreadyExists struct {
	message string
}

func (e *PreconditionAlreadyExists) Error() string { return e.message }
