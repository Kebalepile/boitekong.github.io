package types

import (
	"sync"
)

// Used in Govpage spider
type Links struct {
	Title string            `json:"title"`
	Links map[string]string `json:"links"`
}

// Used in package main main.go
type Crawler interface {
	// initiate the Spider instant
	// Configers chromedp options such as headless flag userAgent & window size
	// Creates Navigates to the allowed domain to crawl
	Launch(wg *sync.WaitGroup)
}

// Used in Heitha spider
type JobPost struct {
	Href          string `json:"href"`
	JobTitle      string `json:"jobTitle"`
	IndustryTitle string `json:"industryTitle"`
	Bullets       string `json:"bullets"`
	Province      string `json:"province"`
	ExpiryDate    string `json:"expiryDate"`
}

// Used in Heitha spider
type HeithaJobs struct {
	Title string    `json:"title"`
	Links []JobPost `json:"links"`
}
