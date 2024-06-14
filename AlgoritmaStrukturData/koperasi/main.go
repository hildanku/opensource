package main

import "fmt"

type Koperasi struct {
	NamaNasabah     string
	JenisNasabah    int
	JumlahPinjaman  int
	LamaPinjaman    int
	Bunga           float64
	TotalBunga      float64
	CicilanPerbulan int
	Hadiah          string
}

var k Koperasi

func main() {

	fmt.Println("=========== Koperasi JebolJaya ===========")
	fmt.Println("Selamat datang di Koperasi JebolJaya")
	fmt.Print("Masukan nama Nasabah : ")
	fmt.Scanln(&k.NamaNasabah)
	fmt.Print("Jenis Nasabah ( pilih '0' untuk Nasabah Baru, pilih '1' untuk Nasabah Lama ) : ")
	fmt.Scanln(&k.JenisNasabah)
	fmt.Print("Jumlah Pinjaman : ")
	fmt.Scanln(&k.JumlahPinjaman)
	fmt.Print("Lama Pinjaman : ")
	fmt.Scanln(&k.LamaPinjaman)

	if k.JumlahPinjaman > 50000000 {
		if k.LamaPinjaman == 12 {
			k.Bunga = 0.90
			k.TotalBunga = float64(k.JumlahPinjaman) * (k.Bunga / 100)
			k.CicilanPerbulan = k.JumlahPinjaman/k.LamaPinjaman + int(k.Bunga)
			if k.JenisNasabah == 0 {
				k.Hadiah = "Anda Mendapatkan Kalender"
			} else if k.JenisNasabah == 1 {
				k.Hadiah = "Anda Mendapatkan Payung"
			} else {
				k.Hadiah = "Masukan jenis nasabah dengan benar!"
			}
		} else if k.LamaPinjaman == 6 {
			k.Bunga = 0.90
			k.TotalBunga = float64(k.JumlahPinjaman) * (k.Bunga / 100)
			k.CicilanPerbulan = k.JumlahPinjaman/k.LamaPinjaman + int(k.Bunga)
			if k.JenisNasabah == 0 {
				k.Hadiah = "Anda Mendapatkan Kalender"
			} else if k.JenisNasabah == 1 {
				k.Hadiah = "Anda Mendapatkan Payung"
			} else {
				k.Hadiah = "Masukan jenis nasabah dengan benar!"
			}
		} else if k.LamaPinjaman == 1 {
			k.Bunga = 0.99
			k.TotalBunga = float64(k.JumlahPinjaman) * (k.Bunga / 100)
			k.CicilanPerbulan = k.JumlahPinjaman/k.LamaPinjaman + int(k.Bunga)
			if k.JenisNasabah == 0 {
				k.Hadiah = "Anda Mendapatkan Kalender"
			} else if k.JenisNasabah == 1 {
				k.Hadiah = "Anda Mendapatkan Payung"
			} else {
				k.Hadiah = "Masukan jenis nasabah dengan benar!"
			}
		}
	} else if k.JumlahPinjaman < 5000000 {
		if k.LamaPinjaman == 12 {
			k.Bunga = 0.95
			k.TotalBunga = float64(k.JumlahPinjaman) * (k.Bunga / 100)
			k.CicilanPerbulan = k.JumlahPinjaman/k.LamaPinjaman + int(k.Bunga)
			if k.JenisNasabah == 0 {
				k.Hadiah = "Anda Mendapatkan Kalender"
			} else if k.JenisNasabah == 1 {
				k.Hadiah = "Anda Mendapatkan Payung"
			} else {
				k.Hadiah = "Masukan jenis nasabah dengan benar!"
			}
		} else if k.LamaPinjaman == 6 {
			k.Bunga = 0.97
			k.TotalBunga = float64(k.JumlahPinjaman) * (k.Bunga / 100)
			k.CicilanPerbulan = k.JumlahPinjaman/k.LamaPinjaman + int(k.Bunga)
			if k.JenisNasabah == 0 {
				k.Hadiah = "Anda Mendapatkan Kalender"
			} else if k.JenisNasabah == 1 {
				k.Hadiah = "Anda Mendapatkan Payung"
			} else {
				k.Hadiah = "Masukan jenis nasabah dengan benar!"
			}
		} else if k.LamaPinjaman == 1 {
			k.Bunga = 0.99
			k.TotalBunga = float64(k.JumlahPinjaman) * (k.Bunga / 100)
			k.CicilanPerbulan = k.JumlahPinjaman/k.LamaPinjaman + int(k.Bunga)
			if k.JenisNasabah == 0 {
				k.Hadiah = "Anda Mendapatkan Kalender"
			} else if k.JenisNasabah == 1 {
				k.Hadiah = "Anda Mendapatkan Payung"
			} else {
				k.Hadiah = "Masukan jenis nasabah dengan benar!"
			}
		} else {
			fmt.Println("Masukan Total Pinjaman dengan Benar!")
		}
	}
	fmt.Println("========== Detail Nasabah ============")
	fmt.Println("Nama Nasabah : ", k.NamaNasabah)
	fmt.Println("Jenis Nasabah : ", k.JenisNasabah)
	fmt.Println("Jumlah Pinjaman : ", k.JumlahPinjaman)
	fmt.Println("Lama Pinjaman : ", k.LamaPinjaman)
	fmt.Println("Total bunga yang didapat : ", k.TotalBunga)
	fmt.Println("Total Cicilan : ", k.CicilanPerbulan)
	fmt.Println("Hadiah : ", k.Hadiah)
	fmt.Println("=====================================")
}
