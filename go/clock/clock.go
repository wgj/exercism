// Package clock uses Clock struct to represent time.
package clock

import "fmt"

const testVersion = 4

// Clock struct represents a 24-hour clock
type Clock struct {
	hour   int
	minute int
}

// New returns a Clock struct with the initialized time.
func New(hour, minute int) Clock {
	c := Clock{}
	c.hour = hour
	c = c.Add(minute)

	return c
}

// String implements Stringer interface.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add returns a new Clock struct with updated minutes. Subtraction is done by adding negative integers.
func (c Clock) Add(minutes int) Clock {
	c.minute += minutes

	// If 'minute' is over an hour, increment 'hour' respectively.
	if c.minute >= 60 {
		c.hour += c.minute / 60
		c.minute = c.minute % 60
	}

	// If 'minute' is under an hour, add the negative hour(s) to 'hour', and add the negative minutes and 60 to 'minute'
	if c.minute < 0 {
		c.hour += c.minute/60 + (-1)
		c.minute = c.minute%60 + 60
	}

	// If 'hour' is negative, add the negative number to the positive number, less the amount over -24. This was a tricky one https://github.com/golang/go/issues/448 https://golang.org/ref/spec#Arithmetic_operators
	if c.hour < 0 {
		c.hour = c.hour%24 + 24
	}

	// If 'hour' is over a day, update 'hour', but drop whatever would have been added to days.
	if c.hour >= 24 {
		c.hour = c.hour % 24
	}

	return c
}
