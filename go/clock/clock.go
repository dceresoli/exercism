// Clock stub file

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

import "fmt"

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.

type Clock struct {
    h, m int
}

func New(hour, minute int) Clock {
    hour += minute/60
    minute %= 60
    if minute < 0 { minute += 60; hour -= 1 }
    hour %= 24
    if hour < 0 { hour += 24 }
    return Clock{hour, minute}
}

func (self Clock) String() string {
    return fmt.Sprintf("%02d:%02d", self.h, self.m)
}

func (self Clock) Add(minutes int) Clock {
    return New(self.h, self.m+minutes)
}

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.
