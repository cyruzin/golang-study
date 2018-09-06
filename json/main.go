package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Result struct
type Result []*Movie

// Movie struct
type Movie struct {
	ID   int        `json:"id"`
	Name string     `json:"name"`
	Year int        `json:"year"`
	Crew *MovieCrew `json:"crew"`
}

// MovieCrew Struct
type MovieCrew struct {
	Cast     []string `json:"cast"`
	Director []string `json:"director"`
	Writer   []string `json:"writer"`
}

func main() {

	r := Result{
		&Movie{
			ID:   1,
			Name: "Elite Squad",
			Year: 2007,
			Crew: &MovieCrew{
				Cast: []string{
					"Wagner Moura",
					"Milhem Cortaz",
					"André Ramiro",
					"Caio Junqueira",
					"Fábio Lago",
					"Maria Ribeiro",
					"Fernanda Machado",
					"Fernanda de Freitas",
				},
				Director: []string{
					"José Padilha",
				},
				Writer: []string{
					"José Padilha",
					"Luiz Eduardo Soares",
					"Bráulio Mantovani",
					"Rodrigo Pimentel",
				},
			},
		},
		&Movie{
			ID:   2,
			Name: "Elite Squad: The Enemy Within",
			Year: 2010,
			Crew: &MovieCrew{
				Cast: []string{
					"Wagner Moura",
					"Milhem Cortaz",
					"André Ramiro",
					"André Mattos",
					"Irandir Santos",
					"Maria Ribeiro",
					"Tainá Müller",
					"Juliana Schalch",
				},
				Director: []string{
					"José Padilha",
				},
				Writer: []string{
					"José Padilha",
					"Bráulio Mantovani",
					"Rodrigo Pimentel",
				},
			},
		},
	}

	data, err := json.MarshalIndent(r, "", "\t")

	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%s", data)

}
