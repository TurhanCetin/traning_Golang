package main

import "fmt"

func main() {

	var name string // var - name of variable - static types

	name = "Turhan"

	var surname = "Çetin" // typesafe bir dil değildir. ancak okunabilirlik açısından type belirtmek önemli olabilir.

	var age int

	age = 21

	var weight float32 = 72.5

	fmt.Println(name)
	fmt.Println(surname)
	fmt.Println(age)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", surname) // "%T" yanına yazılan değişkenin veri tipini söyler \n ise veri tipini bir alt satıra yazsın diye kullandığımız bir parametredir.
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", weight)

}
