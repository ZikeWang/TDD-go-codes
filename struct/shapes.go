package main

import "math"

//Rectangle defines the width and hight of a rectangle
type Rectangle struct {
	width  float64
	height float64
}

//Circle defines the radius of a circle
type Circle struct {
	radius float64
}

//Triangle defines the base and height of a triangle
type Triangle struct {
	base   float64
	height float64
}

//Shape defines the interface, only the variable that has the Area method whose return type is float64 will satisfy the interface.
type Shape interface {
	Area() float64
}

//Perimeter calculates the perimeter of a rectangle which has specific width and height
func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.width + rectangle.height) * 2
}

/*
//Area calculates the area of a rectangle which has specific width and height
func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}
*/

//refactoring function Area() into method of structs above

//Area is a method of strcut Rectangle
func (r Rectangle) Area() float64 { //it's common to use the first character of receiverType as receiverName
	return r.width * r.height
}

//Area is a method of strcut Circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

//Area is method of struct Triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}
