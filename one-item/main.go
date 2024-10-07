package main

import (
	"fmt"
	"sort"
)

type Book struct {
	Title string
	Genre string
	Pages int
	Year  int
}

func knn(book Book, dataset []Book, k int) []Book {
	var recommendations []Book
	type Neighbor struct {
		Book     Book
		Distance float64
	}
	var neighbors []Neighbor
	for _, b := range dataset {
		if b.Genre == book.Genre {
			neighbors = append(neighbors, Neighbor{Book: b, Distance: float64((b.Pages-book.Pages)*(b.Pages-book.Pages) + (b.Year-book.Year)*(b.Year-book.Year))})
		}
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
		{"Livro A", "Ficção", 300, 2010},
		{"Livro B", "Fantasia", 150, 2012},
		{"Livro C", "Ficção", 100, 2018},
		{"Livro D", "Ficção", 400, 2008},
	}

	recommendedBooks := Book{Genre: "Fantasia", Pages: 250, Year: 2011}
	recommendations := knn(recommendedBooks, books, 2)
	fmt.Println("Recomendações de livros:")
	for _, book := range recommendations {
		fmt.Printf("- %s\n", book.Title)
	}
}
