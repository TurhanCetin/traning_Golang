// 1 : int x, float y type conversion sample
// 2 : multiple assing sample x, y = y, x
// 3 : multiple assing sample x, y = y, x
// 4 : shadowing kavramı
// 5 : 40 as a string

// 1

/*package main

import "fmt"

func main() {

	x := 100

	var y float64

	y = float64(x) // bir şeyi float64 vb şeyleri çevirmek istiyorsak aktaracağımız değişken o tipte olmalı

	var z float64

	z = float64(x) // bir şeyi float64 vb şeyleri çevirmek istiyorsak aktaracağımız değişken o tipte olmalı

	fmt.Println(y)
	fmt.Println(z)

}*/

// 2

/*package main

import "fmt"

func main() {

	x := 100

	y := 200

	fmt.Println(x, " ", y)

	x, y = y, x // burada toplu değer aktarımı gerçekleştiriyoruz 1. parmaetre 1'e 2. parametre 2'ye

	fmt.Println(x, " ", y)
}*/

// 3 // utf8 uyumlu bir dildir latin alfabesi dışındaki alfabeler ile de değişkneler oluşturabilir.

/*package main

import "fmt"

func main() {

	yaş := 21

	fmt.Println(yaş) // burada hata vermiyecektir ancak türkçe değişkenler önerilen bir durum değildir.

	// istersek çince bile değşken oluşturabiliriz.
	// 名称 = name

	名称 := "Turhan"

	fmt.Println(名称)

}*/

// 4 Gölgeleme

/*package main

import "fmt"

func main() {
	x := 5

	if true {
		x := 10

		fmt.Println(x) // iki ayı isimli değişkeni farklı değerler girilecektir. ancak ikinci x ilk x değşkenini gölgelemiş oluyor bu da block Scope özelliğinden kaynaklanır Aslında yeni bir x değişkeni oluşturuyoruz
	}

	fmt.Println(x)
}*/

// 5

package main

import (
	"fmt"
	"strconv"
)

func main() { // 65 ASCII'da A 'ya denk gelmektedir.
	x := 65
	var y string

	y = string(x)

	fmt.Println(y) // peki bunu "65" olarak göstermek istersek napıcaz ?

	s := strconv.Itoa(x) //strconv pakedini kullanmalıyız

	fmt.Println(s)
}
