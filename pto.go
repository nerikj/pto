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
			printDateHeading(holiday.Date.Time, holiday.Name)
		}
	}
}

func printDateHeading(date time.Time, name string) {
	days := []string{
		"sön", "mån", "tis", "ons", "tors", "fre", "lör",
	}
	months := []string{
		"jan", "feb", "mar", "apr", "maj", "jun", "jul", "aug", "sep", "okt", "nov", "dec",
	}

	dateStr := fmt.Sprintf(
		"%s %d %s",
		days[date.Weekday()], date.Day(), months[date.Month()-1],
	)

	fmt.Printf("\n# %s (%s)\n", dateStr, name)
}
