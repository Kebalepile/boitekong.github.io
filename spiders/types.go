package spiders


type Links struct {
	Title string            `json:"title"`
	Links map[string]string `json:"links"`
}
