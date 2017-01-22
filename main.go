package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type Adventar struct {
	Title        string
	Url          string
	Creator      string
	Description  string
	Entry_count  int
	Posted_count int
	Calendars    Calendars
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

func scraping(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("failed")
	}

	// Title
	var title string
	var creator string
	doc.Find("h2").Each(func(_ int, s *goquery.Selection) {
		title = s.Text()
	})
	// Description
	doc.Find(".mod-calendarDescription").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})
	// Creator
	doc.Find(".mod-calendarHeader-meta").Each(func(_ int, s *goquery.Selection) {
		creator = s.Find("span").Text()
	})
	fmt.Println("Title:", title)
	fmt.Println("Creator:", creator)
	// Entry
	var entry_count int
	doc.Find(".is-entry").Each(func(c int, s *goquery.Selection) {
		date := s.Find(".mod-calendar-date").Text()
		user := s.Find(".mod-calendar-user").Text()
		fmt.Println(date, user)
		entry_count = c
	})
	fmt.Println(entry_count + 1)
	// Posted
	var posted_count int
	doc.Find(".is-posted").Each(func(c int, s *goquery.Selection) {
		date := s.Find(".mod-calendar-date").Text()
		user := s.Find(".mod-calendar-user").Text()
		fmt.Println(date, user)
		posted_count = c
	})
	fmt.Println(posted_count + 1)

	// comment
	doc.Find(".mod-entryList-comment").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
	// title
	doc.Find(".mod-entryList-title").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
	// url
	doc.Find(".mod-entryList-url").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}

func main() {
	const url = "http://www.adventar.org/calendars/888"
	scraping(url)
}
