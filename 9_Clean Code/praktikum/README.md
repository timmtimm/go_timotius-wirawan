# Praktikum Clean Code

## Problem 1 - Analysis

Berapa banyak kekurangan dalam penulisan kode tersebut?

Terdapat 3 kekurangan utama

Bagian mana saja terjadi kekurangan tersebut?

- penamaan struct menggunakan PascalCase
- penamaan variable dan function menggunakan camelCase
- penamaan variable tidak deskriptif

Tuliskan alasan dari tiap kekurangan tersebut!

Untuk melihat alasan lebih lanjut bisa dilihat pada [problem1.go](./problem1.go)

## Problem 2 - Rewrite

```go
package main

/*
- penamaan struct menggunakan PascalCase
- penamaan variable menggunakan camelCase
*/
type Kendaraan struct {
   totalRoda       int
   kecepatanPerJam int
}

/*
- penamaan struct menggunakan PascalCase
- penamaan variable menggunakan camelCase
*/
type Mobil struct {
   kendaraan
}

/*
- m tidak deskriptif lebih baik diganti mobil
- nama function menggunakan camelCase
*/
func (mobil *mobil) berjalan() {
   mobil.tambahKecepatan(10)
}

/*
- m tidak deskriptif lebih baik diganti mobil
- nama function menggunakan camelCase
- nama variable menggunakan camelCase
*/
func (mobil *mobil) tambahKecepatan(kecepatanBaru int) {
   mobil.kecepatanPerJam = mobil.kecepatanperjam + kecepatanBaru
}

func main(){
	// nama variable menggunakan camelCase
	mobilCepat := mobil{}
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	// nama variable menggunakan camelCase
	mobilLamban := mobil{}
	mobilLamban.berjalan()
}
```