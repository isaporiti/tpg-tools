package counter

import "io"

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

func WithWriter(w io.Writer) option {
	return func(c *counter) error {
		c.writer = w
		return nil
	}
}
