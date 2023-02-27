package main

type Thread struct {
	No     int    `json:"no"`
	Sticky int    `json:"sticky"`
	Closed int    `json:"closed"`
	Now    string `json:"now"`
	Name   string `json:"name"`
	Sub    string `json:"sub"`
}

type Page struct {
	Page    int      `json:"page"`
	Threads []Thread `json:"threads"`
}

type General struct {
	ParsedAt int64             `json:"parsed_at"`
	Board    string            `json:"board"`
	General  map[string]string `json:"general"`
}
