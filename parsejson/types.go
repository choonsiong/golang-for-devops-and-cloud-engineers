package main

type Page struct {
	Name string `json:"page"`
}

type Words struct {
	Page
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func (w Words) GetResponse() string {
	return w.Name
}

type Occurrence struct {
	Page
	Words map[string]int `json:"words"`
}

func (o Occurrence) GetResponse() string {
	return o.Name
}
