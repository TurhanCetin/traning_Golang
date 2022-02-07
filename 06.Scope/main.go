package main

import "fmt"

var packVar = "Package Scope"

func main() { // JavaScripteki scope olayının aynısı func içerisinde tanımladığımız değişkeni içinde bulunduğu blok dışında kullanamayız

	if true {
		var blockVar = "Block Var"

		fmt.Println(blockVar)
	}

	// fmt.Println(blockVar) Burası bir hata verecektir çünlü tanımlandığı if bloğunun dışındadır. Undifened yani tanımlanmadı hatası alırız

	var funcVar = "Func Scope"

	fmt.Println(funcVar)

	fmt.Println(packVar) // Packvar paket scopuna ait olduğu için her yerde kullanabiliriz

	myFunc()

	var name = "Turhan"
	// name := "Fatma" bu kod satırı hata vericektir ":= öncesinde yeni bir değişken yok" ancak bunu şu şekilde kullanırsak
	name, surname := "Fatma", "Çetin" // burası hata vermiyecektir çünkü :='dan önce yeni bir değişken var ve diyoruzki bunlara şu değerleri ata yani := iki işi birden yapabiliyor yeni değişken oluşturup değer atabiliyor

	fmt.Println(name)
	fmt.Println(surname)

}

func myFunc() {
	// fmt.Println(funcVar) // buradada undefined hatası verecektir bloc scope olduğu için
	// fmt.Println(packVar) // packVar burada hata vermiyecektir çünkü package scope oluduğu için her yerde kullanabiliriz
}
