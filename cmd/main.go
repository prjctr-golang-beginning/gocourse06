package main

import (
	"fmt"
	. "gocourse06/pkg"
	"sync"
	"time"
)

func processBloodType(bloodType string, ch chan<- string, sec int) {
	// Тут може бути якась складна логіка обробки
	time.Sleep(time.Duration(sec+1) * time.Second)
	ch <- fmt.Sprintf("Processed blood type: %s", bloodType)
}

func main() {
	// Основи Горутин: Паралельна обробка даних пацієнтів
	c := NewClinic()
	patients := []Patient{
		{"1", "John Doe", 30, "A+"},
		{"2", "Jane Smith", 25, "O-"},
		{"3", "Will Smith", 50, "C-"},
		// Додайте більше пацієнтів тут
	}

	var wg sync.WaitGroup
	for _, p := range patients {
		wg.Add(1)
		go func(p Patient) {
			defer wg.Done()
			c.AddPatient(p)
			fmt.Printf("Patient added: %s\n", p.Name)
		}(p)
	}
	wg.Wait()
	fmt.Printf("%d patients added to the clinic.\n", len(c.Patients()))

	// Синхронізація Горутин: Координація між медичними відділами
	bloodTypes := []string{"A+", "O-", "B+", "AB+"}
	ch := make(chan string, len(bloodTypes))

	for i, bt := range bloodTypes {
		go processBloodType(bt, ch, i)
	}

	for range bloodTypes {
		fmt.Println(<-ch)
	}

	// Використання select/case для паралельної обробки медичних запитів
	addPatientChan := make(chan Patient)
	getPatientChan := make(chan string)
	quitChan := make(chan bool)

	// Горутина для обробки медичних запитів
	go func() {
		for {
			select {
			case p := <-addPatientChan:
				c.AddPatient(p)
			case id := <-getPatientChan:
				patient, exists := c.GetPatient(id)
				if exists {
					fmt.Printf("Patient retrieved: %s, Age: %d, BloodType: %s\n", patient.Name, patient.Age, patient.BloodType)
				} else {
					fmt.Printf("Patient not found with ID: %s", id)
					return
				}
			case <-quitChan:
				fmt.Println("Clinic closed.")
				return
			}
		}
	}()

	// Симуляція додавання пацієнтів
	addPatientChan <- Patient{"1", "Mike Doe", 35, "A+"}
	addPatientChan <- Patient{"2", "Mika Smith", 20, "O-"}

	//c.ForcedLock() // Example for deadlock

	addPatientChan <- Patient{"3", "Mika Smith", 33, "AO-"}

	// Симуляція запиту на отримання даних пацієнта
	getPatientChan <- "1"
	getPatientChan <- "3" // ID, який не існує

	// Дайте час на обробку запитів
	time.Sleep(3 * time.Second)

	// Закриття клініки
	quitChan <- true

	// Дайте час на завершення обробки
	time.Sleep(1 * time.Second)
}
