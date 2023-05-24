package main

import "fmt"

type ticketing struct {
	NamaPembeli  string
	KodeJurusan  int
	Jurusan      string
	KodeJenisBus int
	JenisBus     string
	JumlahTiket  int
	TotalBayar   int
	Bonus        string
}

const Diskon float64 = 0.85
const BandungSe, BandungAc, BandungNonAc int = 225000, 150000, 90000
const JakartaSe, JakartaAc, JakartaNonAc int = 200000, 120000, 70000

var t ticketing

func main() {
	for {

		fmt.Println("========= Agen P.O Bus JayaJaya-Esport ==========")
		fmt.Print("Masukkan nama pembeli (atau tekan 'q' untuk keluar) : ")
		fmt.Scanln(&t.NamaPembeli)

		if t.NamaPembeli == "q" {
			break
		}

		fmt.Print("Jurusan (1. Bandung 2. Jakarta) : ")
		fmt.Scanln(&t.KodeJurusan)
		fmt.Print("Jenis Bus (1. SE 2. AC 3.Non AC ) : ")
		fmt.Scanln(&t.KodeJenisBus)
		fmt.Print("Jumlah Tiket : ")
		fmt.Scanln(&t.JumlahTiket)

		if t.KodeJurusan == 1 {
			t.Jurusan = "Bandung"
			if t.KodeJenisBus == 1 {
				t.JenisBus = "Kelas SE"
				if t.JumlahTiket > 3 {
					t.TotalBayar = int(Diskon * float64(BandungSe) * float64(t.JumlahTiket))
					t.Bonus = "Anda mendapatkan 2x Snack dan 4x Makan"
				} else {
					t.TotalBayar = BandungSe * t.JumlahTiket
					t.Bonus = "Tidak Ada"
				}
			} else if t.KodeJenisBus == 2 {
				t.JenisBus = "Kelas AC"
				if t.JumlahTiket > 3 {
					t.TotalBayar = int(Diskon * float64(BandungAc) * float64(t.JumlahTiket))
					t.Bonus = "Anda mendapatkan 2x Snack"
				} else {
					t.TotalBayar = BandungAc * t.JumlahTiket
					t.Bonus = "Tidak ada"
				}
			} else if t.KodeJenisBus == 3 {
				t.JenisBus = "Kelas Non AC"
				if t.JumlahTiket > 3 {
					t.TotalBayar = int(Diskon * float64(BandungNonAc) * float64(t.JumlahTiket))
					t.Bonus = "Tidak ada"
				} else {
					t.TotalBayar = BandungNonAc * t.JumlahTiket
					t.Bonus = "Tidak ada"
				}
			} else {
				fmt.Println("Anda memilih Kategori yang salah! pilih 1,2 atau 3.")
			}

		} else if t.KodeJurusan == 2 {
			t.Jurusan = "Jakarta"
			if t.KodeJenisBus == 1 {
				t.JenisBus = "Kelas SE"
				if t.JumlahTiket > 3 {
					t.TotalBayar = int(Diskon * float64(JakartaSe) * float64(t.JumlahTiket))
					t.Bonus = "Anda mendapatkan 1x Snack dan 2x Makan"
				} else {
					t.TotalBayar = JakartaSe * t.JumlahTiket
					t.Bonus = "Tidak ada"
				}
			} else if t.KodeJenisBus == 2 {
				t.JenisBus = "Kelas AC"
				if t.JumlahTiket > 3 {
					t.TotalBayar = int(Diskon * float64(JakartaAc) * float64(t.JumlahTiket))
					t.Bonus = "Anda mendapatkan 1x Snack"
				} else {
					t.TotalBayar = JakartaAc * t.JumlahTiket
					t.Bonus = "Tidak ada"
				}

			} else if t.KodeJenisBus == 3 {
				t.JenisBus = "Kelas Non AC"
				if t.JumlahTiket > 3 {
					t.TotalBayar = int(Diskon * float64(JakartaNonAc) * float64(t.JumlahTiket))
					t.Bonus = "Tidak ada"
				} else {
					t.TotalBayar = JakartaNonAc * t.JumlahTiket
					t.Bonus = "Tidak ada"
				}
			} else {
				fmt.Println("Anda memilih Kategori Jenis Bus yang salah! pilih 1,2 atau 3.")
			}
		} else {
			fmt.Println("Anda memilih Kategori Jurusan yang salah!. Pilih 1 atau 2")
		}
		fmt.Println("=================================================")
		fmt.Println("Nama Pembeli:", t.NamaPembeli)
		fmt.Println("Jurusan : ", t.Jurusan)
		fmt.Println("Jenis Bus : ", t.JenisBus)
		fmt.Println("Bonus : ", t.Bonus)
		fmt.Println("Jumlah Tiket:", t.JumlahTiket)
		fmt.Println("Total Bayar : Rp.", t.TotalBayar)
		fmt.Println("Terimakasih telah membeli Tiket di Agen P.O Bus JayaJaya-Esport")
		fmt.Println("=================================================")
	}
}
