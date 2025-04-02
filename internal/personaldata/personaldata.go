package personaldata

import "fmt"

// Ниже создайте структуру Personal
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Ниже создайте метод Print()
func (z Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", z.Name, z.Weight, z.Height)
}
