package counter

import (
	"io"

	"github.com/isaporiti/tpg-tools/interval"
)

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

func WithInterval(i interval.Interval) option {
	return func(c *counter) error {
		c.interval = i
		return nil
	}
}
