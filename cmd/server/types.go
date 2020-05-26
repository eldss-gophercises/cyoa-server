package main

type story map[string]arc

type arc struct {
	Title   string   `json:"title,omitempty"`
	Story   []string `json:"story,omitempty"`
	Options []option `json:"options,omitempty"`
}

type option struct {
	Text  string `json:"text,omitempty"`
	ArcID string `json:"arc,omitempty"`
}
