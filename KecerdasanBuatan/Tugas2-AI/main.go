package main // inisialisasi nama package

import (
	"fmt" // fmt is a Go package that is used to format basic strings, values, inputs, and outputs. 
)

func main() { // buat fungsi main
	var NilaiAwal, Menu, NilaiAkhir int // menentukan tipe data pada tiap variable
	// Input nilai awal
	// dan menyimpan nilai awal pada variable NilaiAwal
	fmt.Print("Masukkan nilai awal: ")
	fmt.Scanln(&NilaiAwal)

	// Input nilai akhir
	fmt.Print("Masukkan nilai akhir: ")
	fmt.Scanln(&NilaiAkhir)

	// pilihan
	fmt.Println("Ketik 1 untuk mencari nilai ganjil, Ketik 2 Untuk mencari nilai genap")
	fmt.Print("1 atau 2? :")
	fmt.Scanln(&Menu)

	// jika user pilih 1 maka dilakukan perhitungan ganjil
	if Menu == 1 {
		fmt.Println("Bilangan ganjil antara", NilaiAwal, "dan", NilaiAkhir, "adalah:")
		for i := NilaiAwal; i <= NilaiAkhir; i++ {
			if i%2 != 0 {
				fmt.Println(i)
			}
		}
	} else if Menu == 2 {
		fmt.Println("Bilangan genap antara", NilaiAwal, "dan", NilaiAkhir, "adalah:")
		for i := NilaiAwal; i <= NilaiAkhir; i++ {
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	} else {
		fmt.Println("Menu tidak valid.")
	}
}