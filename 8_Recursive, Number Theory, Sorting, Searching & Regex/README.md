# Recursive, Number Theory, Sorting, Searching & Regex

## Resume Materi

### Recursive

Recursive function merupakan fungsi yang mengulang atau menggunakan nilai dari hasil pemrosesan pada sebelumnya untuk menghitung nilai berikutnya. Fungsi rekursif biasa digunakan pada pemrosesan nilai yang prosesnya sama sehingga menyederhakan proses logika pada algoritma. Untuk cara mengimplementasikannya sendiri yaitu cukup memanggil kembali fungsi tersebut pada pemrosesan yang algoritmanya sama.

Contoh sederhana adalah pada studi kasus Factorial
```go
func factorial(n int) int {
    if (n == 1) {
        return 1
    } else {
        return n * factorial(n-1)
    }
}
```

### Number Theory

Number Theory adalah cabang ilmu matematika yang mempelajari tentang bilangan (integers).

Adapun topik-topik yang berkaitan dengan teori bilangan antara lain:
- Faktorial
- Bilangan Prima
- Greatest Common Divisor (GCD)
- Least Common Multiple (LCM)

### Searching

Searching adalah proses dari mencari nilai yang dibutuhkan dari nilai-nilai yang ada dalam sebuah list.

Adapun metode searching antara lain:
- Linier Search (O(n))
- Builtins Search