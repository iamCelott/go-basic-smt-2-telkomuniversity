package main

import "fmt"

const nMax = 7919

type Buku struct {
	ID        string
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    int
}

type DaftarBuku [nMax]Buku

/*
I.S. sejumlah n data buku telah siap pada masukan
F.S. pustaka berisi n data buku
*/
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("\nData Buku %d\n", i+1)

		fmt.Print("ID        : ")
		fmt.Scan(&(*pustaka)[i].ID)

		fmt.Print("Judul     : ")
		fmt.Scan(&(*pustaka)[i].Judul)

		fmt.Print("Penulis   : ")
		fmt.Scan(&(*pustaka)[i].Penulis)

		fmt.Print("Penerbit  : ")
		fmt.Scan(&(*pustaka)[i].Penerbit)

		fmt.Print("Eksemplar : ")
		fmt.Scan(&(*pustaka)[i].Eksemplar)

		fmt.Print("Tahun     : ")
		fmt.Scan(&(*pustaka)[i].Tahun)

		fmt.Print("Rating    : ")
		fmt.Scan(&(*pustaka)[i].Rating)
	}
}

/*
I.S. array pustaka berisi n data buku dan belum terurut
F.S. menampilkan buku dengan rating tertinggi
*/
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada data buku")
		return
	}

	maxIdx := 0

	for i := 1; i < n; i++ {
		if pustaka[i].Rating > pustaka[maxIdx].Rating {
			maxIdx = i
		}
	}

	fmt.Println("\n=== BUKU TERFAVORIT ===")
	fmt.Println("Judul    :", pustaka[maxIdx].Judul)
	fmt.Println("Penulis  :", pustaka[maxIdx].Penulis)
	fmt.Println("Penerbit :", pustaka[maxIdx].Penerbit)
	fmt.Println("Tahun    :", pustaka[maxIdx].Tahun)
	fmt.Println("Rating   :", pustaka[maxIdx].Rating)
}

/*
I.S. array pustaka berisi n data buku
F.S. array terurut menurun berdasarkan rating
Metode: Insertion Sort
*/
func UrutBuku(pustaka *DaftarBuku, n int) {

	for i := 1; i < n; i++ {

		temp := (*pustaka)[i]

		j := i - 1

		for j >= 0 && (*pustaka)[j].Rating < temp.Rating {
			(*pustaka)[j+1] = (*pustaka)[j]
			j--
		}

		(*pustaka)[j+1] = temp
	}
}

/*
I.S. pustaka sudah terurut berdasarkan rating
F.S. menampilkan maksimal 5 buku dengan rating tertinggi
*/
func Cetak5Teratas(pustaka DaftarBuku, n int) {

	fmt.Println("\n=== 5 JUDUL DENGAN RATING TERTINGGI ===")

	batas := 5

	if n < 5 {
		batas = n
	}

	for i := 0; i < batas; i++ {
		fmt.Printf("%d. %s (Rating %d)\n",
			i+1,
			pustaka[i].Judul,
			pustaka[i].Rating)
	}
}

/*
I.S. pustaka sudah terurut menurun berdasarkan rating
F.S. menampilkan salah satu buku dengan rating r
Metode: Binary Search
*/
func CariBuku(pustaka DaftarBuku, n int, r int) {

	left := 0
	right := n - 1

	for left <= right {

		mid := (left + right) / 2

		if pustaka[mid].Rating == r {

			fmt.Println("\n=== BUKU DITEMUKAN ===")
			fmt.Println("ID        :", pustaka[mid].ID)
			fmt.Println("Judul     :", pustaka[mid].Judul)
			fmt.Println("Penulis   :", pustaka[mid].Penulis)
			fmt.Println("Penerbit  :", pustaka[mid].Penerbit)
			fmt.Println("Tahun     :", pustaka[mid].Tahun)
			fmt.Println("Eksemplar :", pustaka[mid].Eksemplar)
			fmt.Println("Rating    :", pustaka[mid].Rating)

			return
		}

		// karena urut MENURUN
		if r > pustaka[mid].Rating {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	fmt.Println("\nTidak ada buku dengan rating seperti itu")
}

func main() {

	var pustaka DaftarBuku
	var n int
	var ratingCari int

	fmt.Print("Jumlah buku: ")
	fmt.Scan(&n)

	DaftarkanBuku(&pustaka, n)

	CetakTerfavorit(pustaka, n)

	UrutBuku(&pustaka, n)

	Cetak5Teratas(pustaka, n)

	fmt.Print("\nRating yang dicari: ")
	fmt.Scan(&ratingCari)

	CariBuku(pustaka, n, ratingCari)
}
