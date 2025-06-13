package main

import (
	"fmt"
	"os"
	"strings"
)

const MAX = 10

type Tanaman struct {
	Nama string
}

type Lokasi struct {
	Desa string
}

type Petani struct {
	Nama        string
	Tanaman     Tanaman
	JumlahPanen int
	Lokasi      Lokasi
}

var dataPetani [MAX]Petani

func inputData(jumlah *int) {
	fmt.Print("Masukkan jumlah petani: ")
	_, err := fmt.Scanln(jumlah)
	if err != nil || *jumlah <= 0 || *jumlah > MAX {
		fmt.Println("Jumlah tidak valid. Harus berupa angka 1 hingga", MAX)
		os.Exit(1)
	}

	for i := 0; i < *jumlah; i++ {
		var p Petani
		fmt.Printf("\nData Petani ke-%d\n", i+1)
		fmt.Print("Nama: ")
		fmt.Scanln(&p.Nama)
		fmt.Print("Tanaman: ")
		fmt.Scanln(&p.Tanaman.Nama)
		fmt.Print("Jumlah Panen (kg): ")
		fmt.Scanln(&p.JumlahPanen)
		fmt.Print("Lokasi (Desa): ")
		fmt.Scanln(&p.Lokasi.Desa)
		dataPetani[i] = p
	}
}

func tampilkanData(jumlah int) {
	fmt.Println("--- Data Petani ---")
	for i := 0; i < jumlah; i++ {
		p := dataPetani[i]
		fmt.Printf("%s - %s - %dkg - %s\n", p.Nama, p.Tanaman.Nama, p.JumlahPanen, p.Lokasi.Desa)
	}
}

func selectionSortNama(jumlah int) {
	for i := 0; i < jumlah-1; i++ {
		min := i
		for j := i + 1; j < jumlah; j++ {
			if strings.ToLower(dataPetani[j].Nama) < strings.ToLower(dataPetani[min].Nama) {
				min = j
			}
		}
		dataPetani[i], dataPetani[min] = dataPetani[min], dataPetani[i]
	}
}

func insertionSortPanenDesc(jumlah int) {
	for i := 1; i < jumlah; i++ {
		key := dataPetani[i]
		j := i - 1
		for j >= 0 && dataPetani[j].JumlahPanen < key.JumlahPanen {
			dataPetani[j+1] = dataPetani[j]
			j--
		}
		dataPetani[j+1] = key
	}
}

func linearSearchNama(jumlah int, target string) {
	found := false
	for i := 0; i < jumlah; i++ {
		if strings.EqualFold(dataPetani[i].Nama, target) {
			fmt.Println("Ditemukan (Linear Search):", dataPetani[i])
			found = true
		}
	}
	if !found {
		fmt.Println("Petani tidak ditemukan.")
	}
}

func binarySearchNama(jumlah int, target string) {
	selectionSortNama(jumlah)
	kiri, kanan := 0, jumlah-1
	found := false
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		compare := strings.Compare(strings.ToLower(dataPetani[tengah].Nama), strings.ToLower(target))
		if compare == 0 {
			fmt.Println("Ditemukan (Binary Search):", dataPetani[tengah])
			found = true
			kiri = kanan + 1
		} else if compare < 0 {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	if !found {
		fmt.Println("Petani tidak ditemukan (Binary Search).")
	}
}

func updateData(jumlah int) {
	var namaCari string
	fmt.Print("\nMasukkan nama petani yang ingin diubah: ")
	fmt.Scanln(&namaCari)

	selectionSortNama(jumlah)
	kiri, kanan := 0, jumlah-1
	ketemu := -1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		compare := strings.Compare(strings.ToLower(dataPetani[tengah].Nama), strings.ToLower(namaCari))
		if compare == 0 {
			ketemu = tengah
			kiri = kanan + 1 // keluar dari loop
		} else if compare < 0 {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	if ketemu != -1 {
		fmt.Println("Data ditemukan. Masukkan data baru:")
		fmt.Print("Nama: ")
		fmt.Scanln(&dataPetani[ketemu].Nama)
		fmt.Print("Tanaman: ")
		fmt.Scanln(&dataPetani[ketemu].Tanaman.Nama)
		fmt.Print("Jumlah Panen (kg): ")
		fmt.Scanln(&dataPetani[ketemu].JumlahPanen)
		fmt.Print("Lokasi (Desa): ")
		fmt.Scanln(&dataPetani[ketemu].Lokasi.Desa)
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func deleteData(jumlah *int) {
	var namaHapus string
	fmt.Print("\nMasukkan nama petani yang ingin dihapus: ")
	fmt.Scanln(&namaHapus)

	selectionSortNama(*jumlah)
	kiri, kanan := 0, *jumlah-1
	ketemu := -1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		compare := strings.Compare(strings.ToLower(dataPetani[tengah].Nama), strings.ToLower(namaHapus))
		if compare == 0 {
			ketemu = tengah
			kiri = kanan + 1
		} else if compare < 0 {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	if ketemu != -1 {
		for i := ketemu; i < *jumlah-1; i++ {
			dataPetani[i] = dataPetani[i+1]
		}
		*jumlah--
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func main() {
	var jumlah int
	inputData(&jumlah)

	selesai := false
	for !selesai {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1. Tampilkan Data")
		fmt.Println("2. Update Data")
		fmt.Println("3. Delete Data")
		fmt.Println("4. Urutkan Nama (Ascending)")
		fmt.Println("5. Urutkan Jumlah Panen (Descending)")
		fmt.Println("6. Cari Nama (Linear Search)")
		fmt.Println("7. Cari Nama (Binary Search)")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var pilih int
		fmt.Scanln(&pilih)

		if pilih == 0 {
			selesai = true
		} else if pilih == 1 {
			tampilkanData(jumlah)
		} else if pilih == 2 {
			updateData(jumlah)
		} else if pilih == 3 {
			deleteData(&jumlah)
		} else if pilih == 4 {
			selectionSortNama(jumlah)
			tampilkanData(jumlah)
		} else if pilih == 5 {
			insertionSortPanenDesc(jumlah)
			tampilkanData(jumlah)
		} else if pilih == 6 {
			var cari string
			fmt.Print("Masukkan nama: ")
			fmt.Scanln(&cari)
			linearSearchNama(jumlah, cari)
		} else if pilih == 7 {
			var cari string
			fmt.Print("Masukkan nama: ")
			fmt.Scanln(&cari)
			binarySearchNama(jumlah, cari)
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}

	fmt.Println("Program selesai.")
}
