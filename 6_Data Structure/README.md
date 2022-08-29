# Data Structure

## Resume Materi

### Array

Array adalah struktur data yang berisi sekelompok elemen, dan dapat mencakup satu jenis variabel dengan ukuran alokasi tetap. Tipe data yang berbeda dapat ditangani sebagai elemen dalam array seperti numerik, string, dan boolean.

Sintaksis Array
```go
// initiate array
var countries [2] string

// assign value on array
countries[0] = "India"
countries[1] = "Canada"

// intiate array with value
odd_numbers := [5]int{1, 3, 5, 7, 9} // 1, 3, 5, 7, 9
odd_numbers := [5]int{1, 3, 5} // 1, 3, 5, 0, 0

// initiate array with value on spesific index
even_numbers := [5]int{1: 2, 2: 4} // 0, 2, 4, 0, 0
```

### Slice

Slice adalah struktur data yang berisi group of elements dimana dapat berisi satu jenis variabel seperti array tetapi memiliki dynamic allocation size

Sintaksis Slice
``` go
// create slice from array
var primes = [5]int{2, 3, 5, 7 ,11}
var part_primes = []int := primes[1:4] // 3, 5, 7

// long declaration
var even_numbers = []int // no value
var even_numbers = []int{1, 3, 5, 7, 9} // with value

// short declaration
odd_numbers := []int // no value
odd_numbers := []int{1, 3, 5, 7, 9} // with value

// using make function
var primes = make([]int, 5, 10)

// append
var colors = []string{"red", "green", "yellow"}
colors = append(colors, "purple")

// copy
copied_colors := make([]string, 10)
```

### Map
Map adalah struktur data yang menyimpan data dalam bentuk pasangan key dan value di mana setiap key adalah unik

Sintaks Map
```go
// long declaration
var salary = map[string]int{} // no value
var salary_a = map[string]int{"umam": 1000, "iswanul": 2000} // with value

// short declaration
salary_b := map[string]int{}

// using make function
salary_c := make(map[string]int)
salary_c["doe"] = 7000
```