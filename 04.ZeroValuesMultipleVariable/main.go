package main

import "fmt"

func main() {

	var name string // var - name of variable - static types

	name = "Turhan"

	var surname = "Çetin" // typesafe bir dil değildir. ancak okunabilirlik açısından type belirtmek önemli olabilir.

	name2 := "Fatma" // short declaration yani sadece değişken adı ile değişken tanımlammıza olanak sağlar

	// Bu şekil üç tür değişken tanımlama olanağımız var birde hepsi var ile tanımlanacak bir değişken dizimiz var ise şu şekildede tanımlayabiliriz.
	var (
		username string = "Mösyö"
		userage  int    = 21

		student bool    = true
		height  float32 = 72.5
	)

	// Değişken dizimizin birde bu şekilde tanımlayabiliriz

	// var username, userage, student, height = "Mösyö", 21 , true , 72.5

	// Short declaration şeklinde multiple variable tanımlayabiliriz.

	// username, userage, studen, height, := "Mösyö", 21, true, 72.5

	var age int

	age = 21

	var weight float32 = 72.5

	var zeroValues int // bu değişkenin zero values özelliğinden dolayı println ile yazdırdığımızda 0 değerini verecektir. Bu durum diğer değişken türleri içinde geçerlidir

	var zeroValues2 bool // zero values deniliyor ancak aslında default değer diyebiliriz bu yüzden boolean bir değişken türünde olduğu zaman bize vereceği çıktı false'dur

	var zeroValues3 string // string türünde ise bu durum " " bu şekilde ifade ettiğimiz boş bir string değeridir zero values

	fmt.Println(zeroValues)
	fmt.Println(zeroValues2)
	fmt.Println(zeroValues3)

	fmt.Println(name)
	fmt.Println(surname)
	fmt.Println(age)
	fmt.Println(name2)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", surname) // "%T" yanına yazılan değişkenin veri tipini söyler \n ise veri tipini bir alt satıra yazsın diye kullandığımız bir parametredir.
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", weight)

	fmt.Println(username)
	fmt.Println(userage)
	fmt.Println(student)
	fmt.Println(height)

}
