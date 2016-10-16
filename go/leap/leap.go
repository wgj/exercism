// leap checks for leap years.
package leap

const testVersion = 2

// IsLeapYear returns true if `year` is a leap year, and false if not. A year
// is a leap year if it is divisible by 4, except when divisible by 100,
// unless divisible by 400.
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
