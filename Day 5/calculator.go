package main

import "fmt"

func main() {

	var scan int

	fmt.Println("Pilih Calculator yang diinginkan")
	fmt.Println("1.tambah")
	fmt.Println("2.kurang")
	fmt.Println("3.kali")
	fmt.Println("4.bagi")

	fmt.Scan(&scan, "Masukkan Inputan : ")

	switch {
	case scan == 1:

		var a int
		var b int
		var c int

		fmt.Println("masukkan int ke 1")
		fmt.Scan(&a)
		fmt.Println("masukkan int ke 2")
		fmt.Scan(&b)

		c = a + b

		fmt.Println(c)

	case scan == 2:

		var a int
		var b int
		var c int

		fmt.Println("masukkan int ke 1")
		fmt.Scan(&a)
		fmt.Println("masukkan int ke 2")
		fmt.Scan(&b)

		c = a - b

		fmt.Println(c)

	case scan == 3:

		var a int
		var b int
		var c int

		fmt.Println("masukkan int ke 1")
		fmt.Scan(&a)
		fmt.Println("masukkan int ke 2")
		fmt.Scan(&b)

		c = a * b

		fmt.Println(c)

	case scan == 4:

		var a int
		var b int
		var c int

		fmt.Println("masukkan int ke 1")
		fmt.Scan(&a)
		fmt.Println("masukkan int ke 2")
		fmt.Scan(&b)

		c = a / b

		fmt.Println(c)

	default:

		fmt.Println("input salah")
	}

}
