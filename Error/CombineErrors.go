package Error

type CombineErrors struct {
	Errs []error
}

func NewErrors() CombineErrors {
	return CombineErrors{
		Errs: []error{},
	}
}

func (c *CombineErrors) Add(err error) {
	if err != nil {
		c.Errs = append(c.Errs, err)
	}
}
