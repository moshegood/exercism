package twelve

import (
	"fmt"
	"strings"
)

var days = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var dailyGift = []string{
	"and a Partridge in a Pear Tree.",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Laping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

func Verse(day int) string {
	if day < 1 || day > 12 {
		return ""
	}
	onTheDay := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", days[day-1])
	gs := []string{}
	for d := day - 1; d >= 0; d-- {
		gs = append(gs, dailyGift[d])
	}
	gift := strings.Join(gs, ", ")
	if strings.HasPrefix(gift, "and ") {
		gift = gift[4:]
	}
	return onTheDay + gift
}

func Song() string {
	verses := make([]string, 12)
	for i := 0; i < 12; i++ {
		verses[i] = Verse(i + 1)
	}
	return strings.Join(verses, "\n")
}
