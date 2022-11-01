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