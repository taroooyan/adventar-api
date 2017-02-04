package main

import (
	"google.golang.org/appengine/aetest"
	"testing"
)

func TestIsErrorStatus(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	// normal url
	req1, err := inst.NewRequest("GET", "/888", nil)
	if err != nil {
		t.Fatalf("Failed to create req1: %v", err)
	}

	url := "http://www.adventar.org/calendars/888"
	result := IsErrorStatus(url, req1)
	if true == result {
		t.Errorf("result is %v", result)
		t.Fail()
	}

	// error url
	req2, err := inst.NewRequest("GET", "/88", nil)
	if err != nil {
		t.Fatalf("Failed to create req2: %v", err)
	}

	url = "http://www.adventar.org/calendars/88"
	result = IsErrorStatus(url, req2)
	if false == result {
		t.Errorf("result is %v", result)
		t.Fail()
	}
}

func TestIsErrorNumber(t *testing.T) {
	// normal
	number := "10"
	result := IsErrorNumber(number)
	if true == result {
		t.Errorf("number is %s %v", number, result)
		t.Fail()
	}

	// illegal
	number = "123ab"
	result = IsErrorNumber(number)
	if false == result {
		t.Errorf("number is %s %v", number, result)
		t.Fail()
	}
}

func TestScraping(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/888", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	url := "http://www.adventar.org/calendars/888"
	result := Scraping(url, req)
	title := "Aizu Advent Calendar 2015"
	if result.Title != title {
		t.Errorf("Return title is %s", result.Title)
		t.Fail()
	}
}
