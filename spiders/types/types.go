package types

import (
	"sync"
)
type Links struct {
	Title string            `json:"title"`
	Links map[string]string `json:"links"`
}

type Crawler interface {
	Launch(wg *sync.WaitGroup)
}