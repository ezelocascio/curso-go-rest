package main

//Movie struct
type Movie struct {
	Name     string `json:"name"`
	Year     int    `json: "year"`
	Director string `json: "director"`
}

//Movies array
type Movies []Movie
