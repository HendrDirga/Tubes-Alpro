package main

import (
	"fmt"
	"strings"
)

type Petani struct {
	Nama        string
	Tanaman     string
	JumlahPanen int
	Lokasi      string
}

func inputData(petani *[]Petani, n *int) {
	fmt.Print("Masukkan jumlah petani: ")
	fmt.Scanln(n)
	for i := 0; i < *n; i++ {
		var p Petani
		fmt.Printf("Data petani ke-%d\n", i+1)
		fmt.Print("Nama: ")
		fmt.Scanln(&p.Nama)
		fmt.Print("Tanaman: ")
		fmt.Scanln(&p.Tanaman)
		fmt.Print("Jumlah Panen (kg): ")
		fmt.Scanln(&p.JumlahPanen)
		fmt.Print("Lokasi: ")
		fmt.Scanln(&p.Lokasi)
		*petani = append(*petani, p)
	}
}

func sortByPanenDescending(petani []Petani) {
	n := len(petani)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if petani[i].JumlahPanen < petani[j].JumlahPanen {
				petani[i], petani[j] = petani[j], petani[i]
			}
		}
	}
}

func sortByNamaAscending(petani []Petani) {
	n := len(petani)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if strings.ToLower(petani[i].Nama) > strings.ToLower(petani[j].Nama) {
				petani[i], petani[j] = petani[j], petani[i]
			}
		}
	}
}

func searchByNama(petani []Petani, target string) {
	found := false
	for _, p := range petani {
		if strings.EqualFold(p.Nama, target) {
			fmt.Println("Ditemukan:", p)
			found = true
		}
	}
	if !found {
		fmt.Println("Petani tidak ditemukan.")
	}
}

func searchByTanaman(petani []Petani, jenis string) {
	found := false
	for _, p := range petani {
		if strings.EqualFold(p.Tanaman, jenis) {
			fmt.Println("Ditemukan:", p)
			found = true
		}
	}
	if !found {
		fmt.Println("Tanaman tidak ditemukan.")
	}
}

func main() {
	var petani []Petani
	var n int

	inputData(&petani, &n)

	fmt.Println("\n--- Data Setelah Disorting (Panen Terbanyak) ---")
	sortByPanenDescending(petani)
	for _, p := range petani {
		fmt.Printf("%s - %s - %dkg - %s\n", p.Nama, p.Tanaman, p.JumlahPanen, p.Lokasi)
	}

	fmt.Println("\n--- Sorting Berdasarkan Nama ---")
	sortByNamaAscending(petani)
	for _, p := range petani {
		fmt.Printf("%s - %s - %dkg - %s\n", p.Nama, p.Tanaman, p.JumlahPanen, p.Lokasi)
	}

	var cari string
	fmt.Print("\nMasukkan nama petani untuk dicari: ")
	fmt.Scanln(&cari)
	searchByNama(petani, cari)

	fmt.Print("\nMasukkan jenis tanaman untuk dicari: ")
	fmt.Scanln(&cari)
	searchByTanaman(petani, cari)
}
