package counter

import "testing"

func TestCounter_NewCounter(t *testing.T) {
	t.Run("counter.NewCounter returns a pointer to a counter", func(t *testing.T) {
		t.Parallel()

		c, _ := NewCounter()

		if c == nil {
			t.Errorf("Expected a new counter, got %v", c)
		}
	})

	t.Run("counter.NewCounter accepts an initial value and sets it to the to the counter", func(t *testing.T) {
		t.Parallel()
		c, _ := NewCounter(WithInitialCount(5))

		got := c.Next()
		if got != 5 {
			t.Errorf("Expected 5, got %v", c.count)
		}
	})

	t.Run("counter.NewCounter doesn't accept negative initial values", func(t *testing.T) {
		t.Parallel()
		c, err := NewCounter(WithInitialCount(-1))

		if err != ErrNoNegativeValues {
			t.Errorf("Expected %v, got %v", ErrNoNegativeValues, err)
		}

		if c != nil {
			t.Errorf("Expected nil, got %v", c)
		}
	})
}

func TestCounter_Next(t *testing.T) {
	t.Run("counter.Next should return 0 on first call", func(t *testing.T) {
		t.Parallel()
		c, _ := NewCounter()

		got := c.Next()

		if got != 0 {
			t.Errorf("Expected 0, got %v", got)
		}
	})

	t.Run("counter.Next should return 0, 1, 2, 3 and so on on subsequent calls", func(t *testing.T) {
		t.Parallel()
		c, _ := NewCounter()

		for want := 0; want < 10; want++ {
			got := c.Next()
			if got != want {
				t.Errorf("Expected %v, got %v", want, got)
				return
			}
		}
	})
}
