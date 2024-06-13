package Common

type CombineErrors struct {
	Errs []error
}

func NewCombineErrors() CombineErrors {
	return CombineErrors{
		Errs: []error{},
	}
}

func (c *CombineErrors) AddError(err error) {
	if err != nil {
		c.Errs = append(c.Errs, err)
	}
}
