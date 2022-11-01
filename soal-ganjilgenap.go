/*
Soal Ganjil-Genap

Buatlah sebuah program yang menerima input berupa dua bilangan bulat positif secara berulang. satu dari bilangan tersebut harus bernilai ganjil dan lainnya adalah genap. Tidak diperbolehkan sama sama genap atau ganjil. Tampilkan hasil dari modulo antara perkalian kedua bilangan tersebut dan penjumlahannya. Contoh (4 x 5) mod (4 + 1) = 2, untuk bilangn 4 dan 5. Program akan berhenti apabila bilangan yang diinputkan adalah ganjil semua atau genap semua. Tampilkan seperti contoh berikut:

______________________________________________________________________
| 4 5
| 2
| 9 2
| 7
| 11 7
| pasangan bilangan ganjil? true
______________________________________________________________________

______________________________________________________________________
| Spesifikasi
| Input : dua buah bilangan bil1 dan bil2 (berulang)
| Proses: menghitung (bil1 x bil2) mod (bil1 + bil2)
			Perulangan berhenti apabila bil1 dan bil2
			sama sama ganjil atau genap
| Output: hasil perhitungan dan bil1 dan bil2 ganjil?
______________________________________________________________________
| Bilangan dikatakan genap apabila dibagi dua hasilnya tidak bersisa,
| atau tidak berkoma, sedangkan ganjil sebaliknya.
| Contoh : 4 dikatakan genap karena di bagi 2 hasilnya adalah 2
|			(0 dibelakang koma diabaikan)
			11 dikatakan ganjil karena kalau di bagi 2 hasilnya mengandung koma
| Secara matematik genap apabila bilangan mod 2 = 0. dan sebaliknay ganjil.
______________________________________________________________________
*/

/*
Program Ganjilgenap
Kamus
	Status	: boolean
	bil1, bil2 : integer
Algoritma
1	input(bil1, bil2)
2	status <- (bil1 mod 2 = 0 and bil2 mod 2 = 0) OR
			  (bil1 mod 2 != 0 and bil2 mod 2 != 0)
3	while not status do
4		output( (bil1 x bil2) mod (bil1 + bil2) )
5		input(bil1, bil2)
6		status <- (bil1 mod 2 = 0 and bil2 mod 2 = 0) OR
				  (bil1 mo 2 != 0 and bil2 mod 2 != 0)
7		endwhile
8		output("Pasangan bilangan ganjil?", bil1 mod 2 != 0 and bil1 mod 2 != 0)
*/

package main

import "fmt"

func main() {
	var status bool
	var bil1, bil2 int

	fmt.Scan(&bil1, &bil2)
	status = (bil1%2 == 0 && bil2%2 == 0) || (bil1%2 != 0 && bil2%2 != 0)

	for !status {
		fmt.Println((bil1 * bil2) % (bil1 + bil2))
		fmt.Scan(&bil1, &bil2)
		status = (bil1%2 == 0 && bil2%2 == 0) || (bil1%2 != 0 && bil2%2 != 0)
	}
	fmt.Println("Pasangan bilangan ganjil?", bil1%2 != 0 && bil1%2 != 0)
}
