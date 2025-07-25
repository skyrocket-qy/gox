package erx

type CallerInfo struct {
	Function string
	File     string
	Line     int
	Msg      string
}

type CtxErr struct {
	Code        Code
	CallerInfos []CallerInfo
	Cause       string // original error message
}

func (e *CtxErr) Error() string {
	return e.Code.Str()
}

func (e *CtxErr) SetCode(c Code) *CtxErr {
	if e == nil {
		return nil
	}

	e.Code = c
	return e
}
