package main

import "fmt"

func main() {
	//(Print - Println) -Printf

	fmt.Println("Merhaba") // Burada conver edilcek bir veri olmadığı için direk merhabayı ekran basacaktır tek fark println'nin  bir alt satıra inecek olmasıdır.
	fmt.Print("Merhaba")
	fmt.Printf("Merhaba")

	name := "Turhan"

	fmt.Print(name) // burada şu şekil bir çıktı olacaktır TurhanTurhan bir alt satırda Turhan yazacaktır bu da println'nin bir alt satıra inmesidir.
	fmt.Println(name)
	fmt.Print(name)

	fmt.Print("Benim adım", name)
}
