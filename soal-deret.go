/*
SOAL DERET

Buatlah algoritma untuk menampilkan N buah bilangan bulat yang pertama, dimana N positif integer dan merupakan input dari user

____________________________________________________________________________
Spesifikasi
Input: N integer positif
Proses: menampilkan nilai iterasi
output: 1 2 3 ... N
____________________________________________________________________________
Jenis perulangan
	a. Kondisi perulangan: iterasi <= N
	b. Stop kondisi: > N
	c. Jumlah iterasi: N
____________________________________________________________________________
*/

/*
{WHILE - KONDISI PERULANGAN}
Program deret
Kamus
	N, iterasi: integer
Algoritma
1	output("N= ")
2	input(N)
3	iterasi <- 1
4	while iterasi <= N do
5		output(iterasi)
6		iterasi <- iterasi + 1
7	endwhile
8	output("done")
*/

package main

import "fmt"

func main() {
	var N, iterasi int
	fmt.Print("N= ")
	fmt.Scanln(&N)
	iterasi = 1
	for iterasi <= N {
		fmt.Print(iterasi)
		iterasi = iterasi + 1
	}
	fmt.Println("done")
}