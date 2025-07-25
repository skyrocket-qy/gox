package erx

var MaxCallStackSize = 10

var ErrToCode = func(err error) Code {
	return ErrUnknown
}
