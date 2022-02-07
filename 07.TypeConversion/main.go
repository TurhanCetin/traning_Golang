package main

import "fmt"

func main() {

	x := 10 // typesafe bir dildir ancak belirmesekde değişkene verdiğimiz değer ile bir type belirleyecektir.

	y := 10.0

	fmt.Printf("%v %T\n", x, x) // %v yanındaki değişkenin veri tipini verir
	fmt.Printf("%v %T\n", y, y)

	//fmt.Println(x+y) // aynı türde veriler olmadıkları için hata verecektir bu işlemleri yapabilmek için tür dönüşümü yapmalıyız

	// TYPE CONVERSİON tür dönüşümü
	// type(value) = int(y)

	fmt.Println(x + int(y)) // int(y) sadece o işlem yapıldığı anda int çevirir daha sonrasında kendi türüne dönecektir.(geçiçi)

	// Flaout64,int8 ve int16 gibi değişkenlerdede Type conversion geçerlidir.

	// var x int16 = 128

	// var y float64 = 128.0

	// y = int16(y);

	// fmt.Println(x + y)

}
