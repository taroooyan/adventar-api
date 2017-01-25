package main

import (
  "testing"
)

func TestIsErrorStatus(t *testing.T) {
  // normal url
  url := "http://www.adventar.org/calendars/888"
  result := IsErrorStatus(url)
  if false != result {
    t.Errorf("result is %v", result)
    t.Fail()
  }

  // error url
  url = "http://www.adventar.org/calendars/88"
  result = IsErrorStatus(url)
  if true != result {
    t.Errorf("result is %v", result)
    t.Fail()
  }
}

func TestIsErrorNumber(t *testing.T) {
  // normal
  number := "10"
  result := IsErrorNumber(number)
  if false != result {
    t.Errorf("number is %s %v", number, result)
    t.Fail()
  }

  // illegal
  number = "123ab"
  result = IsErrorNumber(number)
  if true != result {
    t.Errorf("number is %s %v", number, result)
    t.Fail()
  }
}

func TestScraping(t *testing.T) {
  url := "http://www.adventar.org/calendars/888"
  result := Scraping(url)
  title := "Aizu Advent Calendar 2015"
  if result.Title != title {
    t.Errorf("Return title is %s", result.Title)
    t.Fail()
  }
}
