package main

import (
	"fmt"
	"reflect"
)

type ServiceLocator struct {
	services []any
	types    []reflect.Type
	values   []reflect.Value
}

func (s *ServiceLocator) Register(some any) {
	s.services = append(s.services, some)
	s.types = append(s.types, reflect.TypeOf(some))
	s.values = append(s.values, reflect.ValueOf(some))
}

func (s *ServiceLocator) Get(some any) bool {
	k := reflect.TypeOf(some).Elem()
	kind := k.Kind()
	if kind == reflect.Ptr {
		k = k.Elem()
		kind = k.Kind()
	}
	for i, t := range s.types {
		if kind == reflect.Interface && t.Implements(k) {
			reflect.Indirect(
				reflect.ValueOf(some),
			).Set(s.values[i])
			return true
		} else if kind == reflect.Struct && k.AssignableTo(t.Elem()) {
			reflect.ValueOf(some).Elem().Set(s.values[i])
			return true
		}
	}
	return false
}

type Concrete struct {
	Name string
}

func (c *Concrete) Do() {}

type Doer interface {
	Do()
}

func main() {
	l := ServiceLocator{}
	l.Register(&Concrete{"tomate"})

	var x Doer
	if l.Get(&x) {
		fmt.Println("by interface pointer ok")
		fmt.Println(x)
	}

	var z Doer
	if l.Get(&z) {
		fmt.Println("by interface ok")
		fmt.Println(z)
	}

	var y *Concrete
	if l.Get(&y) {
		fmt.Println("by struct pointer ok")
		fmt.Println(y)
	}
}
