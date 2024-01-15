package main

import (
	"fmt"
	"time"
)

// Орган - легені
func lungs(oxygenChannel chan<- bool) {
	for {
		time.Sleep(2 * time.Second) // Час на "вдих"
		oxygenChannel <- true       // Відправлення кисню до серця
		fmt.Println("Lungs: Oxygen sent to the heart.")
	}
}

// Орган - серце
func heart(oxygenChannel <-chan bool) {
	for {
		select {
		case <-oxygenChannel:
			fmt.Println("Heart: Oxygen received, continuing to beat.")
		case <-time.After(4 * time.Second): // Час очікування на кисень
			fmt.Println("Heart: No oxygen received, heart stopped!")
			return
		}
	}
}

func main() {
	oxygenChannel := make(chan bool)

	// Запуск горутин для легенів та серця
	go lungs(oxygenChannel)
	go heart(oxygenChannel)

	// Даємо час програмі для виконання (наприклад, 30 секунд)
	time.Sleep(30 * time.Second)
}
