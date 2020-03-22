package controller

import (
	"fmt"
	"github.com/era/covid-telegram/api"
)

func Instructions() string {
	return "/country {country}: searches stats for {country}\n" +
		   "/all: returns all total cases, recovery, and deaths.\n"
}

func Country(country string) string {
	r, err := api.ByCountry(country)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("%s:\nCases: %d\nDeaths: %d\nRecovered: %d\n" +
		"New Cases: %d\nNew Deaths: %d", r.Name, r.Cases, r.Deaths, r.Recovered,
		r.TodayCases, r.TodayDeaths)
}

func All() string {
	r, err := api.GetAll()
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("Cases: %d\nDeaths: %d\nRecovered: %d\n",
		r.Cases, r.Deaths, r.Recovered)
}