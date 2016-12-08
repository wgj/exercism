package twelve

import "fmt"

const testVersion = 1

// verse contains the ordinal number and song verse.
type verse struct {
	ordinal string
	verse   string
}

var stanzas []verse

func init() {
	stanzas = make([]verse, 12)
	stanzas[0].ordinal = "first"
	stanzas[0].verse = "a Partridge in a Pear Tree"
	stanzas[1].ordinal = "second"
	stanzas[1].verse = "two Turtle Doves"
	stanzas[2].ordinal = "third"
	stanzas[2].verse = "three French Hens"
	stanzas[3].ordinal = "fourth"
	stanzas[3].verse = "four Calling Birds"
	stanzas[4].ordinal = "fifth"
	stanzas[4].verse = "five Gold Rings"
	stanzas[5].ordinal = "sixth"
	stanzas[5].verse = "six Geese-a-Laying"
	stanzas[6].ordinal = "seventh"
	stanzas[6].verse = "seven Swans-a-Swimming"
	stanzas[7].ordinal = "eighth"
	stanzas[7].verse = "eight Maids-a-Milking"
	stanzas[8].ordinal = "ninth"
	stanzas[8].verse = "nine Ladies Dancing"
	stanzas[9].ordinal = "tenth"
	stanzas[9].verse = "ten Lords-a-Leaping"
	stanzas[10].ordinal = "eleventh"
	stanzas[10].verse = "eleven Pipers Piping"
	stanzas[11].ordinal = "twelfth"
	stanzas[11].verse = "twelve Drummers Drumming"
}

// Song iterates through Verse with all 12 lines.
func Song() string {
	var s string
	for i := 1; i <= 12; i++ {
		s = fmt.Sprintf("%s%s\n", s, Verse(i))
	}
	return s
}

// Verse interpolates "on the `ordinal` day of Christmas..."
func Verse(n int) string {
	s := fmt.Sprintf("On the %s day of Christmas my true love gave to me", stanzas[n-1].ordinal)
	for i := n; i > 0; i-- {
		// If this is the last line, prepend `and`, and exit the for-loop.
		if i == 1 && n != 1 {
			s = fmt.Sprintf("%s, and %s", s, stanzas[i-1].verse)
			break
		}
		// Append the verse.
		s = fmt.Sprintf("%s, %s", s, stanzas[i-1].verse)
	}
	// Append a period at the end of the sentence.
	s = fmt.Sprintf("%s.", s)
	return s
}
