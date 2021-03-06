// Leap stub file

// The package name is expected by the test program.
package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// It's good style to write a comment here documenting IsLeapYear.
func IsLeapYear(year int) bool {
    switch {
        case year % 400 == 0: { return true  }
        case year % 100 == 0: { return false }
        case year %   4 == 0: { return true  }
    }
    return false
}
