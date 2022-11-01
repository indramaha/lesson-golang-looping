/*
SOAL BILANGAN

Buatlah algoritma yang menerima input berupa bilangan bulat positif, kemudian proses input berhenti apabila input adalah negatif, kemudian tampitlkan banyaknya bilangan positif

______________________________________________________________________
| Contoh 1:				| Contoh 2:									|
| <-- input start -->	| <-- input start -->						|
| 12					| -12										|
| 3						| <-- input end -->							|
| 21					| bilangan positif = 0						|
| 0						|											|
| 5						|											|
| 74					|											|
| -19					|											|
| <-- input end -->		|											|
|Bilangan positif = 6	|											|
______________________________________________________________________

______________________________________________________________________
| Spesifikasi														|
| Input : bilangan bulat posiif										|
| Proses: 															|
|		- menerima input (berulang)									|
|		- menghitung banyaknya bilangan positif						|
| Output: banyak bilangan positif									|
______________________________________________________________________
| Jenis perulangan													|
| 	a. Kondisi perulangan: bilangan = iteger positif				|
|	b. Stop kondisi: bilangan = integer negatif						|
|	c. Jumlah iterasi: tidak diketahui								|
______________________________________________________________________
*/

/*
Program bilangan
Kamus
	nilai, positif : integer
Algoritma
1	positif <- 0
2	input(nilai)
3	while nilai >= 0 do
4		positif <- positif + 1
5		input(nilai)
6	endwhile
7	output("Bilangan positif: ", positif)
*/

package main

import "fmt"

func main() {
	var nilai, positif int
	positif = 0
	fmt.Scanln(&nilai)
	// WHILE nilai >= 0 DO
	for nilai >= 0 {
		positif = positif + 1
		fmt.Scanln(&nilai)
	}
	fmt.Println("Bilangan Positif: ", positif)
}
