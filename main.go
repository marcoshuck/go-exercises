package main

import "fmt"

type car struct {
	Trademark string
	Owner     string
	Seats     int
	Mileage   float64
}

// get Trademark
func (c car) getTrademark() {
	fmt.Println("Car's make is:", c.Trademark)
}

func (c car) getOwner() {
	fmt.Println("Car belongs to:", c.Owner)
}
func (c car) getSeats() {
	fmt.Println("It has:", c.Seats, "seats")
}
func (c car) getMileage() {
	fmt.Println("with:", c.Mileage, "miles")
}

func (c *car) setOwner(s string) {
	c.Owner = s
	fmt.Println("the new owner is: ", s)
}

func main() {
	c1 := car{
		Trademark: "Bentley",
		Owner:     "Santiago",
		Seats:     2,
		Mileage:   10000,
	}
	c2 := car{
		Trademark: "Mercedes",
		Owner:     "Marcos",
		Seats:     4,
		Mileage:   20000,
	}

	c1.getTrademark()
	c1.getSeats()
	c1.getOwner()
	c1.getMileage()
	c2.setOwner("Santy")

}
