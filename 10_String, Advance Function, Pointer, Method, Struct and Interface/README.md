# String, Advance Function, Pointer, Method, Struct and Interface

## Resume Materi

### String

Terdapat beberapa operasi yang bisa digunakan untuk memanipulasi string:
- menghitung panjang string dengan return integer panjang string (len)
    ```go
    fmt.Println(len("Hello"))
    ```
- membandingkan dua string (compare)
    ```go
    fmt.Println("abc" == "abd")
    ```
- melakukan pengecekan string apakah mengandung string yang dicari dengan return booelan (contains)
    ```go
    fmt.Println(strings.Contains("string", "str"))
    ```
- Mengambil substring dari sebuah string (substring)
    ```go
    string := "cat_and_dog"
    fmt.Println(string[8:len(string)])
    ```
- Melakukan penggantian string dengan string yang baru (replace)
    ```go
    string := "`12456789"
    fmt.Println(strings.Replace(string, "`", "0", -1))
    ```
- Melakukan penambahan string (insert)
    ```go
    string := "grn"
    fmt.Println(string[:2] + "ee" + string[2:])
    ```

### Pointer

Pointer adalah variabel yang menyimpan alamat memori dari variabel lain.

Ada dua jenis operator di pointer:

1. Operator Dereferencing (*)
    - Deklarasi variabel pointer
    - Akses Value yang disimpan dalam alamat memory

2. Operator Referencing (&)
    - Mengembalikan alamat memory dari variabel
    - Akses alamat memory dari variabel ke pointer

Sintaksis pointer
```go
var name string = "John"
var nameAddress *string = &name

fmt.Println(name) // John
fmt.Println(&name) // (memory address)
fmt.Println(*nameAddress) // John
fmt.Println(name) // (memory address)
```

### Interface

Interface adalah kumpulan metode yang diimplementasikan pada objek

Sintaksis interaface

```go
    /*
    type interface_name interface {
        method_name <return_type>
    }
    */
```