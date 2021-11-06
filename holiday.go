package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type holiday struct {
	Date holidayDate `json:"date"`
	Name string      `json:"localName"`
}

type holidayDate struct {
	time.Time
}

func (h *holidayDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	(*h).Time = t
	return nil
}

func fetchHolidays(year string) ([]holiday, error) {
	url := fmt.Sprintf("https://date.nager.at/api/v3/PublicHolidays/%s/SE", year)

	body, err := httpGet(url)
	if err != nil {
		return nil, err
	}

	holidays := []holiday{}
	err = json.Unmarshal(body, &holidays)
	if err != nil {
		return nil, err
	}

	return holidays, nil
}
