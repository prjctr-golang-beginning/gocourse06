package pkg

// Clinic - структура, що містить мапу пацієнтів і м'ютекс для синхронізації доступу
type Clinic struct {
	patients map[string]Patient
}

// NewClinic створює нову клініку з порожньою мапою пацієнтів
func NewClinic() *Clinic {
	return &Clinic{
		patients: make(map[string]Patient),
	}
}

// AddPatient додає нового пацієнта у мапу
func (c *Clinic) AddPatient(p Patient) {
	c.patients[p.ID] = p
}

// GetPatient повертає пацієнта за ID
func (c *Clinic) GetPatient(id string) (Patient, bool) {
	p, exists := c.patients[id]
	return p, exists
}
