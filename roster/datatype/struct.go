package datatype

import "time"

type Person struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Age        uint8     `json:"age"`
	Gender     uint8     `json:"gender"`
	Birth      time.Time `json:"birth"`
	BirthPlace string    `json:"birthPlace"`
	GroupTags  []string  `json:"groupTags"`
	HistroyId  string    `json:"historyId"`
}

type Record struct {
	Id      string    `json:"id"`
	Time    time.Time `json:"time"`
	Content string    `json:"content"`
	Comment string    `json:"comment"`
}

type History struct {
	Id      string   `json:"id"`
	Records []Record `json:"records"`
}

type Groups struct {
	Groups []GroupTag `json:"groups"`
}

type GroupTag struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Members []uint8 `json:"members"`
}
