package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Animal описує структуру даних для тварини
type Animal struct {
	ID     int
	Health int
	Hunger int
	Mood   int
}

// Enclosure описує статус вольєра
type Enclosure struct {
	ID     int
	IsOpen bool
}

// Feeder описує статус кормушки
type Feeder struct {
	ID      int
	IsEmpty bool
}

// Генерує тестові дані для тварин
func generateAnimals(n int) []Animal {
	var animals []Animal
	for i := 0; i < n; i++ {
		animal := Animal{
			ID:     i,
			Health: rand.Intn(100),
			Hunger: rand.Intn(100),
			Mood:   rand.Intn(100),
		}
		animals = append(animals, animal)
	}
	return animals
}

// Генерує тестові дані для вольєрів
func generateEnclosures(n int) []Enclosure {
	var enclosures []Enclosure
	for i := 0; i < n; i++ {
		enclosure := Enclosure{
			ID:     i,
			IsOpen: rand.Intn(2) == 1,
		}
		enclosures = append(enclosures, enclosure)
	}
	return enclosures
}

// Генерує тестові дані для кормушок
func generateFeeders(n int) []Feeder {
	var feeders []Feeder
	for i := 0; i < n; i++ {
		feeder := Feeder{
			ID:      i,
			IsEmpty: rand.Intn(2) == 1,
		}
		feeders = append(feeders, feeder)
	}
	return feeders
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Генеруємо тестові дані
	animals := generateAnimals(10)
	enclosures := generateEnclosures(5)
	feeders := generateFeeders(3)

	// Виводимо згенеровані дані
	fmt.Println("Animals:", animals)
	fmt.Println("Enclosures:", enclosures)
	fmt.Println("Feeders:", feeders)
}
