package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

// I.S. Data tersedia dalam piranti masukan
// F.S. Array t berisi sejumlah n karakter yang dimasukkan user,
//      Proses input selama karakter bukanlah '.' dan n <= NMAX
func isiArray(t *tabel, n *int) {
	var ch string
	*n = 0

	fmt.Print("Teks : ")

	for *n < NMAX {
		fmt.Scan(&ch)

		if ch == "." {
			break
		}

		t[*n] = rune(ch[0])
		*n++
	}
}

// I.S. Terdefinisi array t yang berisi sejumlah n karakter
// F.S. n karakter dalam array muncul di layar
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

// I.S. Terdefinisi array t yang berisi sejumlah n karakter
// F.S. Urutan isi array t terbalik
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	var balik tabel
	for i := 0; i < n; i++ {
		balik[i] = t[i]
	}
	balikanArray(&balik, n)
	for i := 0; i < n; i++ {
		if t[i] != balik[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Println(`Akhiri teks dengan simbol "."`)
	isiArray(&tab, &m)
	fmt.Println("Palindrom ?", palindrom(tab, m))
	balikanArray(&tab, m)
	fmt.Print("Reverse teks : ")
	cetakArray(tab, m)
}