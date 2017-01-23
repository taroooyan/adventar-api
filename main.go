package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

type Adventar struct {
	Title        string
	Url          string
	Creator      string
	Description  string
	Entry_count  int
	Posted_count int
	Calendars    [25]Calendars
}

type Calendars struct {
	Date      int
	User      string
	Icon      string
	Comment   string
	Title     string
	Url       string
	Is_entry  bool
	Is_posted bool
}

func scraping(url string) (data Adventar) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("failed")
	}

	// Title
	doc.Find("h2").Each(func(_ int, s *goquery.Selection) {
		data.Title = s.Text()
	})

	// Description
	doc.Find(".mod-calendarDescription").Each(func(_ int, s *goquery.Selection) {
		description, _ := s.Html()
		data.Description = description
	})

	// Creator
	doc.Find(".mod-calendarHeader-meta").Each(func(_ int, s *goquery.Selection) {
		data.Creator = s.Find("span").Text()
	})

	// User
	doc.Find(".mod-calendar-cell").Each(func(_ int, s *goquery.Selection) {
		date := s.Find(".mod-calendar-date").Text()
		user := s.Find(".mod-calendar-user").Text()
		var icon string
		s.Find("img").Each(func(_ int, s *goquery.Selection) {
			url, _ = s.Attr("src")
			icon = url
		})
		dateI, _ := strconv.Atoi(date)
		data.Calendars[dateI-1].Date = dateI
		data.Calendars[dateI-1].User = user
		data.Calendars[dateI-1].Icon = icon
	})

	// Entry
	var entryCount int
	doc.Find(".is-entry").Each(func(i int, s *goquery.Selection) {
		date := s.Find(".mod-calendar-date").Text()
		dateI, _ := strconv.Atoi(date)
		data.Calendars[dateI-1].Is_entry = true
		entryCount = i
	})
	data.Entry_count = entryCount + 1

	// Posted
	var postedCount int
	doc.Find(".is-posted").Each(func(i int, s *goquery.Selection) {
		date := s.Find(".mod-calendar-date").Text()
		dateI, _ := strconv.Atoi(date)
		data.Calendars[dateI-1].Is_posted = true
		postedCount = i
	})
	data.Posted_count = postedCount + 1

	// comment
	doc.Find(".mod-entryList-comment").Each(func(i int, s *goquery.Selection) {
		data.Calendars[i].Comment = s.Text()
	})

	// title
	doc.Find(".mod-entryList-title").Each(func(i int, s *goquery.Selection) {
		data.Calendars[i].Title = s.Text()
	})

	// url
	doc.Find(".mod-entryList-url").Each(func(i int, s *goquery.Selection) {
		data.Calendars[i].Url = s.Text()
	})
	return
}

func main() {
	const url = "http://www.adventar.org/calendars/888"
	data := scraping(url)
	dataJson, err := json.Marshal(data)
	if err != nil {
		return
	}
	fmt.Println(string(dataJson))
}
