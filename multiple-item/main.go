package main

import (
	"fmt"
	"math"
	"sort"
)

type Book struct {
	Title  string
	Genres []string
	Pages  int
	Year   int
}

func genreSimilarity(genres1, genres2 []string) float64 {
	intersection := 0
	set := make(map[string]bool)
	for _, genre := range genres1 {
		set[genre] = true
	}

	for _, genre := range genres2 {
		if set[genre] {
			intersection++
		}
	}

	union := len(genres1) + len(genres2) - intersection
	if union == 0 {
		return 0
	}
	return float64(intersection) / float64(union)
}

func distance(book1, book2 Book) float64 {
	pageDiff := book1.Pages - book2.Pages
	yearDiff := book1.Year - book2.Year
	genreSim := genreSimilarity(book1.Genres, book2.Genres)
	return math.Sqrt(float64(pageDiff*pageDiff+yearDiff*yearDiff)) * (1.0 - float64(genreSim))
}

func knn(book Book, dataset []Book, k int) []Book {
	var recommendations []Book
	type Neighbor struct {
		Book     Book
		Distance float64
	}
	var neighbors []Neighbor
	for _, b := range dataset {
		distance := distance(book, b)
		neighbors = append(neighbors, Neighbor{Book: b, Distance: distance})
	}

	sort.Slice(neighbors, func(i, j int) bool {
		return neighbors[i].Distance < neighbors[j].Distance
	})

	for i := 0; i < k && i < len(neighbors); i++ {
		recommendations = append(recommendations, neighbors[i].Book)
	}

	return recommendations

}

func main() {
	books := []Book{
		{"Livro A", []string{"Ficção", "Aventura"}, 300, 2010},
		{"Livro B", []string{"Fantasia", "Ação"}, 150, 2012},
		{"Livro C", []string{"Ciência", "Ficção"}, 220, 2018},
		{"Livro D", []string{"Ficção", "Mistério"}, 240, 2008},
		{"Livro E", []string{"Aventura", "Fantasia"}, 320, 2020},
	}

	recommendBook := Book{Genres: []string{"Ficção", "Suspense"}, Pages: 250, Year: 2011}
	recommendations := knn(recommendBook, books, 2)
	fmt.Println("Recomendações de livros:")
	for _, book := range recommendations {
		fmt.Printf("- %s\n", book.Title)
	}
}
