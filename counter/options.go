package counter

import "io"

type option func(*counter) error

func WithInitialCount(value uint) option {
	return func(c *counter) error {
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
