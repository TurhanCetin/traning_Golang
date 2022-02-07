// 1-) studentName --> Jhon Doe, grade --> 77, isPassed --> true değişkenleri
// 3 farklı yöntem ile oluşturunuz

package main

import "fmt"

func main() { // !!! ÖNEMLİ NOT = var ile değişkne tanımlayacaksa func dışında tanımlayabiliriz ancak short declaration ile tanımlamak istersek izin vermeyecektir ve func dışında tanımladığımız değişkeni kerana yazırdımak istediğimizde gene fmt'yi func içinde kullanmak zorundayız hata verecektir.
	var studentName string = "Jhon Doe"

	grade := 77

	var isPassed bool = true

	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)

	// 2-) yukarıda belirtilen değişkenleri tek satır içerisinde tanımlayınız.

	// studentName, grade, isPassed := "Jhon Doe", 77, true

	// 3-) "Declaration", "Assign", "Initialization", "Initial Value" kavramlarını açıklayınız. (Terminoloji)
	// 4-) "Statically Typed" vs "Dynamically Typed" ifadelerini GO ve Python üzerinden gösteriniz.

	// 5-) ":=" vs "=" aradaki farkı gösteriniz, double declaration

	name := "Turhan" // BU değişken türü belirtmeden var parametresi kullanmadan değişken tanımlamadır Short declaration'dır.
	fmt.Println(name)
	name = "Çetin" // Bu ise oluşturduğumuz değişkene veri aktarmamızı sağlar.
	fmt.Printf(name)
}
