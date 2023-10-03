package types

import (
	"sync"
)

// Used in Govpage spider
type Links struct {
	Title     string     `json:"title"`
	BlogPosts []BlogPost `json:"blogPosts"`
}

// Used in Govpage spider
type BlogPost struct {
	Href       string   `json:"href"`
	Title      string   `json:"title"`
	Content    []string `json:"content"`
	PostedDate string   `json:"date"`
	Iframe     string   `json:"iframe"`
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
	Apply         string `json:"apply"`
	JobTitle      string `json:"jobTitle"`
	JobSpecFields string `json:"jobSpecFields"`
	Details       string `json:"details"`
	Province      string `json:"province"`
	ExpiryDate    string `json:"expiryDate"`
}

// Used in Heitha spider
type HeithaJobs struct {
	Title string    `json:"title"`
	Links []JobPost `json:"links"`
}

// Used in ProPersonnel spider
type ProJobPost struct {
	JobTitle      string            `json:"jobTitle"`
	StartDate     string            `json:"startDate"`
	VacancyType   string            `json:"vacancyType"`
	JobSpecFields string            `json:"jobSpecFields"`
	Location      map[string]string `json:"location"`
	Details       string            `json:"details"`
	Contact       string            `json:"contact"`
	Apply         string            `json:"apply"`
}

// Used in ProPersonnel spider
type ProPersonnelJobs struct {
	Title string       `json:"title"`
	Links []ProJobPost `json:"links"`
}
