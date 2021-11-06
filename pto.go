package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	year := os.Args[2]
	month, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	holidays, err := fetchHolidays(year)
	if err != nil {
		log.Fatal(err)
	}

	for _, holiday := range holidays {
		if int(holiday.Date.Month()) == month {
			date := swedishDate(holiday.Date.Time)
			fmt.Printf("\n# %s (%s)\n", date, holiday.Name)
		}
	}
}

func swedishDate(t time.Time) string {
	days := []string{
		"sön", "mån", "tis", "ons", "tors", "fre", "lör",
	}

	months := []string{
		"jan", "feb", "mar", "apr", "maj", "jun", "jul", "aug", "sep", "okt", "nov", "dec",
	}

	date := fmt.Sprintf(
		"%s %d %s",
		days[t.Weekday()], t.Day(), months[t.Month()-1],
	)

	return date
}
