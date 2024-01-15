package pkg

import "sync"

// Clinic - структура, що містить мапу пацієнтів і м'ютекс для синхронізації доступу
type Clinic struct {
	patients map[string]Patient
	mu       sync.RWMutex
}

// NewClinic створює нову клініку з порожньою мапою пацієнтів
func NewClinic() *Clinic {
	return &Clinic{
		patients: make(map[string]Patient),
	}
}

// AddPatient додає нового пацієнта у мапу
func (c *Clinic) AddPatient(p Patient) {
	c.mu.Lock() // show code without mutex and with --race flag to demonstrate race condition
	defer c.mu.Unlock()
	c.patients[p.ID] = p
}

// GetPatient повертає пацієнта за ID
func (c *Clinic) GetPatient(id string) (Patient, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	p, exists := c.patients[id]
	return p, exists
}

func (c *Clinic) Patients() map[string]Patient {
	return c.patients
}

func (c *Clinic) ForcedLock() {
	c.mu.Lock()
}
