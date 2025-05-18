package main

import "fmt"

type Dog struct {
	Name  string
	Age   int
	Breed string
}

func NewDog(name string, age int, breed string) *Dog {
	return &Dog{
		Name:  name,
		Age:   age,
		Breed: breed,
	}
}

func (d *Dog) GetAge() int {
	return d.Age
}

func (d *Dog) SetAge(newAge int) {
	d.Age = newAge
}

func (d *Dog) Bark() {
	fmt.Printf("%s говорит: Гав-гав!\n", d.Name)
}

func main() {
	myDog := NewDog("Бобик", 3, "Дворняжка")

	fmt.Printf("Имя: %s, Возраст: %d, Порода: %s\n", myDog.Name, myDog.GetAge(), myDog.Breed)
	
	myDog.SetAge(4)
	fmt.Printf("Новый возраст: %d\n", myDog.GetAge())
	
	myDog.Bark()
}