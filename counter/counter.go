package counter

import "errors"

type counter struct {
	count int
}

func NewCounter(options ...option) (*counter, error) {
	c := &counter{}
	for _, opt := range options {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

type option func(*counter) error

func WithInitialCount(value int) option {
	return func(c *counter) error {
		if value < 0 {
			return ErrNoNegativeValues
		}

		c.count = value
		return nil
	}
}

var ErrNoNegativeValues = errors.New("no negative values allowed")

func (c *counter) Next() (next int) {
	next = c.count
	c.count++
	return next
}
