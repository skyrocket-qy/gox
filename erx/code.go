package erx

type Code interface {
	Str() string
}

var _ Code = CodeImp("")

type CodeImp string

const (
	ErrUnknown CodeImp = "500.0000"
)

func (c CodeImp) Str() string {
	return string(c)
}
