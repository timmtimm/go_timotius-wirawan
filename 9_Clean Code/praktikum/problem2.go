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