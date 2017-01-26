package twelve

import (
	"strings"
)

const testVersion = 1

type eachday struct {
	day  string
	gift string
}

var days = []eachday{
	{"first", "a Partridge in a Pear Tree"},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"}}

func Verse(day int) string {
	day--
	songday := "On the " + days[day].day + " day of Christmas my true love gave to me, "
	if day == 0 {
		songday += days[day].gift
	} else {
		for i := day; i > 0; i-- {
			songday += days[i].gift + ", "
		}
		songday += "and " + days[0].gift
	}

	return songday + "."
}

func Song() string {
	var song [12]string
	for i := 1; i <= 12; i++ {
		song[i-1] = Verse(i)
	}
	return strings.Join(song[:], "\n") + "\n"
}
