package main

type Quiz struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"description"`
}

type Question struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Options       string `json:"options"`
	CorrectOption int    `json:"correct_option"`
	Quiz          int    `json:"quiz"`
	Points        int    `json:"points"`
}
