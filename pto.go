package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	year := os.Args[1]
	month, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	holidays, err := fetchHolidays(year)
	if err != nil {
		log.Fatal(err)
	}

	printCalendar(strconv.Itoa(month), year)

	for _, holiday := range holidays {
		if int(holiday.Date.Month()) == month {
			printDateHeading(holiday.Date.Time, holiday.Name)
		}
	}
}

// TODO: cal version can vary between Linux distributions (with varying functionality). Alternatives?
func printCalendar(month string, year string) {
	out, err := exec.Command("cal", "-3", month, year).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("```\n%s```\n", out)
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
