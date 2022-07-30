package main

import "fmt"

type rectangle struct {
	lenght float64
	width  float64
}

func (r rectangle) getLenght() float64 {
	return r.lenght
}

func (r rectangle) getWidth() float64 {
	return r.width
}

func (r *rectangle) setLenght(l1 float64) float64 {
	r.lenght = l1
	return l1
}

func (r *rectangle) setWidth(w1 float64) float64 {
	r.width = w1
	return w1
}

func (r rectangle) getArea() float64 {
	return r.lenght * r.width

}
func (r rectangle) getPerimeter() float64 {
	return 2 * (r.lenght + r.width)
}

func main() {

	rec1 := rectangle{
		lenght: 10,
		width:  20,
	}
	fmt.Println("the area of the rectangle is: ", rec1.getArea())
	fmt.Println("the perimeter of the rectangle is: ", rec1.getPerimeter())
	fmt.Println("the new length is: ", any(rec1.setLenght(30)))
	fmt.Println("the new area of the rectangle is: ", rec1.getArea())

}
