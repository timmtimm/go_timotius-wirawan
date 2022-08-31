# Time Complexity & Space Complexity

## Resume Materi

### Time Complexity

Time Complexity digunakan untuk mengestimasikan running time sebuah program. Complexity bisa dilihat dengan mengukur waktu maksimum pada primitive opeartions yang berjalan.

Perbandingan pada masing-masing time complexities
- Constant time (O(1))
    Hanya berjalan sekali
- Linear Time (O(n))
    Berjalan sebanyak n kali dimana biasanya memiliki karakteristik ada perulangan satu kali dengan kondisi terburuk sebanyak n kali
- Linear Time (O(n+m))
    Berjalan sebanyak n+m kali dimana biasanya memiliki karakteristik ada perulangan dua kali dengan kondisi terburuk sebanyak n+m kali
- Logarithmic Time (O(log n))
    Nilai dari n adalah setengah nilai dari iterasi sebelumnya dengan karakteristik perulangan sebanyak satu kali dengan kondisi setelah perulangan sebelumnya dibagi dua
- Quadratic Time (O(n^2))
    Berjalan sebanyak n kali n sehingga memiliki karakteristik ada perulangan dengan n kali di dalam perulangan dengan n kali (nested loop)
- Exponential and factorial time (O(n!) dan O(2^n))
   Exponential tumbuh dengan mengalikan dengan jumlah yang konstan sedangkan faktorial tumbuh dengan mengalikan dengan jumlah yang meningkat

### Time Limit

Time Limit ditetapkan pada online test yang dimana biasanya antara 1 sampai 10 detik.

- n <= 1.000.000 ekspektasi time complexity adalah O(n) atau O(n log n)
- n <= 10.000 ekspektasi time complexity adalah O(n^2)
- n <= 500 ekspekatis time comlexity adalah O(n^3)

### Space Complexity

Memory limit menyediakan informasi tentang ekspektasi space complextiy. Dengan menggunakan space complexity dapat mengetimasikan angka dari variabel yang dideklarasikan dari sebuah program/