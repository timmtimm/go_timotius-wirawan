# Basic Programming

## Resume Materi

### Variable
Variable digunakan untuk menyimpan informasi dalam program komputer, mereka memberikan nama deskriptif dan mereka memiliki tipe data.

Tipe data yang tersedia dalam Golang:
- Boolean
- Numeric
    - Integer
        - uintt8
        - uint16
        - unit32
        - uint64
        - int8
        - int16
        - int32
        - int64
        - int
        - rune
        - uint
        - byte
    - Float
        - float32
        - float64
    - Complex
        - complex64
        - complex128
- String

Untuk cara deklarasi variabel ada dua yaitu:
- Long
    ```go
    var name string
    // var <variable_name> <type_data>

    var name string = "moryku"
    // var <variable_name> <type_data> = <value>

    var name, gender string
    // var <list_variable_name> <type data>

    var name, gender string  = "moryku", "L"
    // var <list_variable_name> <type data> = <value>

    const pi float64 = 3.14
    // const <variable_name> <type_data> = <value>

    const (
        pi float64 = 3.14
        pi2
        age = 10
    )
    // const (<list_variable_name> <type data (optional)> = <value>)
    ```
- Short
    ```go
    name := "moryku"
    // <variable_name> = <value>
    ```

### Operator
Berikut operasi yang dapat dilakukan Golang
| Jenis         | Operator                                       |
|---------------|------------------------------------------------|
| Arithmetic    | `+, -, *, /, %, ++, --`                        |
| Comparasion   | `==, !=, >, <, >=, <=`                         |
| Logical       | `&&, \|\|, !`                                  |
| Bitwise       | `&, \|, ^, <<, >>`                             |
| Assignment    | `=, +=, -=, *=, /=, &=, <<=, >>=, &=, ^=, \|=` |
| Miscellaneous | `&, *`                                         |

### Control Structuring
Untuk melakukan percabangan dapat menggunakan if dan switching sedangkan untuk melakukan perulangan dapat menggunakan for.

Sintaksis percabangan
``` go
// sintaks if
if condition {
    code
}

if (condition) {
    code
}

if condition1 && condition2 {
    code
}

if declare_variable; condition {
    code
}

// sintaks if else
if condition {
    code
} else {
    code
}

// sintaks if, if else, else
if condition {
    code
} else if condition {
    code
} else {
    code
}

// sintaks switch
switch variable {
    case condition:
        code
    default:
        code
}
```

Sintaksis perulangan
```go
// sintaks for
for intiate_variable; conditon; change_value {
    code
}

// sintaks while
for condition {
    code
}
```
