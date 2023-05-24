package main

import "fmt"

type ticketing struct {
	NamaPembeli string
	Jurusan     int
	JenisBus    int
	JumlahTiket int
	TotalBayar  int
}

const Diskon float64 = 0.85
const BandungSe, BandungAc, BandungNonAc int = 225000, 150000, 90000
const JakartaSe, JakartaAc, JakartaNonAc int = 200000, 120000, 70000

var t ticketing

func main() {
	fmt.Println("============= Agen P.O Bus JayaJaya-eSport ==============")
	fmt.Print("Masukkan nama pembeli: ")
	fmt.Scanln(&t.NamaPembeli)

	fmt.Print("Jurusan (1. Bandung 2. Jakarta) : ")
	fmt.Scanln(&t.Jurusan)
	fmt.Print("Jenis Bus (1. SE 2. AC 3.Non AC ) : ")
	fmt.Scanln(&t.JenisBus)
	fmt.Print("Jumlah Tiket : ")
	fmt.Scanln(&t.JumlahTiket)
	fmt.Println("=================================================")
	if t.Jurusan == 1 {
		if t.JenisBus == 1 {
			if t.JumlahTiket > 3 {
				t.TotalBayar = int(Diskon * float64(BandungSe) * float64(t.JumlahTiket))
			} else {
				t.TotalBayar = BandungSe * t.JumlahTiket
			}
			fmt.Println("Nama Pembeli:", t.NamaPembeli)
			fmt.Println("Jurusan : Bandung")
			fmt.Println("Jenis Bus : Kelas SE")
			fmt.Println("Bonus : 2x Snack dan 4x Makan")
			fmt.Println("Jumlah Tiket :", t.JumlahTiket)
			fmt.Println("Total Bayar : Rp.", t.TotalBayar)
		} else if t.JenisBus == 2 {
			if t.JumlahTiket > 3 {
				t.TotalBayar = int(Diskon * float64(BandungAc) * float64(t.JumlahTiket))
			} else {
				t.TotalBayar = BandungAc * t.JumlahTiket
			}
			fmt.Println("Nama Pembeli:", t.NamaPembeli)
			fmt.Println("Jurusan : Bandung")
			fmt.Println("Jenis Bus : Kelas AC")
			fmt.Println("Bonus : 2x Snack")
			fmt.Println("Jumlah Tiket:", t.JumlahTiket)
			fmt.Println("Total Bayar : Rp.", t.TotalBayar)

		} else if t.JenisBus == 3 {
			if t.JumlahTiket > 3 {
				t.TotalBayar = int(Diskon * float64(BandungNonAc) * float64(t.JumlahTiket))
			} else {
				t.TotalBayar = BandungNonAc * t.JumlahTiket
			}
			fmt.Println("Nama Pembeli:", t.NamaPembeli)
			fmt.Println("Jurusan : Bandung")
			fmt.Println("Jenis Bus : Kelas Non AC")
			fmt.Println("Bonus : Tidak Ada")
			fmt.Println("Jumlah Tiket:", t.JumlahTiket)
			fmt.Println("Total Bayar : Rp.", t.TotalBayar)
		} else {
			fmt.Println("Anda memilih Kategori yang salah! pilih 1,2 atau 3.")
		}

	} else if t.Jurusan == 2 {
		if t.JenisBus == 1 {
			if t.JumlahTiket > 3 {
				t.TotalBayar = int(Diskon * float64(JakartaSe) * float64(t.JumlahTiket))
			} else {
				t.TotalBayar = JakartaSe * t.JumlahTiket
			}
			fmt.Println("Nama Pembeli:", t.NamaPembeli)
			fmt.Println("Jurusan : Jakarta")
			fmt.Println("Jenis Bus : Kelas SE")
			fmt.Println("Bonus : 2x Snack dan 4x Makan")
			fmt.Println("Jumlah Tiket:", t.JumlahTiket)
			fmt.Println("Total Bayar:", t.TotalBayar)
		} else if t.JenisBus == 2 {
			if t.JumlahTiket > 3 {
				t.TotalBayar = int(Diskon * float64(JakartaAc) * float64(t.JumlahTiket))
			} else {
				t.TotalBayar = JakartaAc * t.JumlahTiket
			}
			fmt.Println("Nama Pembeli:", t.NamaPembeli)
			fmt.Println("Jurusan : Jakarta")
			fmt.Println("Jenis Bus : Kelas AC")
			fmt.Println("Bonus : 2x Snack")
			fmt.Println("Jumlah Tiket:", t.JumlahTiket)
			fmt.Println("Total Bayar:", t.TotalBayar)

		} else if t.JenisBus == 3 {
			if t.JumlahTiket > 3 {
				t.TotalBayar = int(Diskon * float64(JakartaNonAc) * float64(t.JumlahTiket))
			} else {
				t.TotalBayar = JakartaNonAc * t.JumlahTiket
			}
			fmt.Println("Nama Pembeli:", t.NamaPembeli)
			fmt.Println("Jurusan : Jakarta")
			fmt.Println("Jenis Bus : Kelas Non AC")
			fmt.Println("Bonus : Tidak Ada")
			fmt.Println("Jumlah Tiket:", t.JumlahTiket)
			fmt.Println("Total Bayar:", t.TotalBayar)
		} else {
			fmt.Println("Anda memilih Kategori Jenis Bus yang salah! pilih 1,2 atau 3.")
		}
	} else {
		fmt.Println("Anda memilih Kategori Jurusan yang salah!. Pilih 1 atau 2")
	}

}
