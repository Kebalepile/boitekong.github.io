package types

type Spider struct {
	Name           string
	AllowedDomains []string
	Shutdown       context.CancelFunc
}

type Links struct {
	Title string            `json:"title"`
	Links map[string]string `json:"links"`
}
