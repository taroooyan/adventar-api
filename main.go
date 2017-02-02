package main

import (
	"encoding/json"
	"fmt"
	"github.com/taroooyan/goquery"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"net/http"
	"strconv"
	"strings"
)

type Adventar struct {
	Is_error     bool
	Title        string
	Url          string
	Creator      string
	Description  string
	Entry_count  int
	Posted_count int
	Calendars    [25]Calendar
}

type Calendar struct {
	Date      int
	User      string
	Icon      string
	Comment   string
	Title     string
	Url       string
	Is_entry  bool
	Is_posted bool
}

func IsErrorStatus(url string, r *http.Request) bool {
	// it is urlfetch instead of http for GAE.
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)

	res, err := client.Get(url)
	if err != nil || res.StatusCode != 200 {
		return true
	}
	return false
}

func Scraping(url string, r *http.Request) (data Adventar) {
	data.Url = url

	if IsErrorStatus(url, r) {
		data.Is_error = true
		return
	}

	// usually argument of goguery.NewDocument is only url(string) created by PuerkitoBio but I want to use GAE. So I use taroooyan/goquery package that is rewrite code to use GAE.
	doc, err := goquery.NewDocument(url, r)
	if err != nil {
		data.Is_error = true
		return
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

	// Calendars.Date, Calendars.User, Calendars.Icon
	doc.Find(".mod-calendar-cell").Each(func(_ int, s *goquery.Selection) {
		d := s.Find(".mod-calendar-date").Text()
		date, _ := strconv.Atoi(d)
		user := s.Find(".mod-calendar-user").Text()
		s.Find("img").Each(func(_ int, s *goquery.Selection) {
			icon, _ := s.Attr("src")
			data.Calendars[date-1].Icon = icon
		})
		data.Calendars[date-1].Date = date
		data.Calendars[date-1].User = user
	})

	// Calendars.Is_entry, Entry_count
	var entryCount int
	doc.Find(".is-entry").Each(func(c int, s *goquery.Selection) {
		d := s.Find(".mod-calendar-date").Text()
		date, _ := strconv.Atoi(d)
		data.Calendars[date-1].Is_entry = true
		entryCount = c
	})
	data.Entry_count = entryCount + 1

	// Calendars.Is_posted, Posted_count
	var postedCount int
	doc.Find(".is-posted").Each(func(c int, s *goquery.Selection) {
		d := s.Find(".mod-calendar-date").Text()
		date, _ := strconv.Atoi(d)
		data.Calendars[date-1].Is_posted = true
		postedCount = c
	})
	data.Posted_count = postedCount + 1

	// Calendars.Comment
	doc.Find(".mod-entryList-comment").Each(func(c int, s *goquery.Selection) {
		dateId, _ := s.Attr("data-reactid")
		d := strings.Split(strings.Split(dateId, "-")[2], ".")[0]
		date, _ := strconv.Atoi(d)
		data.Calendars[date-1].Comment = s.Text()
	})

	// Calendars.Title
	doc.Find(".mod-entryList-title").Each(func(c int, s *goquery.Selection) {
		dateId, _ := s.Attr("data-reactid")
		d := strings.Split(strings.Split(dateId, "-")[2], ".")[0]
		date, _ := strconv.Atoi(d)
		data.Calendars[date-1].Title = s.Text()
	})

	// Calendars.Url
	doc.Find(".mod-entryList-url").Each(func(c int, s *goquery.Selection) {
		dateId, _ := s.Attr("data-reactid")
		d := strings.Split(strings.Split(dateId, "-")[2], ".")[0]
		date, _ := strconv.Atoi(d)
		data.Calendars[date-1].Url = s.Text()
	})
	return
}

func IsErrorNumber(number string) bool {
	if _, err := strconv.Atoi(number); err == nil {
		return false
	}
	return true
}

func CreateData(w http.ResponseWriter, r *http.Request) {
	const baseUrl = "http://www.adventar.org/calendars/"
	number := strings.Split(r.URL.Path, "/")[1]
	if IsErrorNumber(number) {
		fmt.Fprintf(w, "Request number error")
		return
	}

	data := Scraping(baseUrl+number, r)
	dataJson, err := json.Marshal(data)
	if err != nil {
		return
	}
	fmt.Fprintf(w, string(dataJson))
}

func init() {
	http.HandleFunc("/", CreateData)
	http.ListenAndServe(":80", nil)
}
