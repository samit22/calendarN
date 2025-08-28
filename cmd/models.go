package cmd

type Row struct {
	Day   int  `json:"day,omitempty"`
	Today bool `json:"today,omitempty"`
	Blank bool `json:"blank"`
}

type Calendar struct {
	Rows  [][]Row `json:"rows"`
	Year  int     `json:"year"`
	Month int     `json:"month"`
	Days  int     `json:"days"`
}
