# Join, Union, Agregasi, Subquery & Function

## Resume Materi

## Join

JOIN digunakan untuk menggabungkan baris dari dua atau lebih tabel, berdasarkan kolom terkait di antara mereka.

Berikut adalah berbagai jenis JOIN dalam SQL
1. (INNER) JOIN
    Mengembalikan rekaman yang memiliki nilai yang cocok di kedua tabel
    ![alt text](./screenshots/inner_join.png "inner_join")
    ```sql
    SELECT column_name(s)
    FROM table1
    INNER JOIN table2
    ON table1.column_name = table2.column_name;
    ```
2. LEFT (OUTER) JOIN
    Mengembalikan semua record dari tabel kiri, dan record yang cocok dari tabel kanan
    ![alt text](./screenshots/left_join.png "left_join")
    ```sql
    SELECT column_name(s)
    FROM table1
    LEFT JOIN table2
    ON table1.column_name = table2.column_name;
    ```
3. RIGHT (OUTER) JOIN
    Mengembalikan semua record dari tabel kanan, dan record yang cocok dari tabel kiri
    ![alt text](./screenshots/right_join.png "right_join")
    ```sql
    SELECT column_name(s)
    FROM table1
    RIGHT JOIN table2
    ON table1.column_name = table2.column_name;
    ```
4. FULL (OUTER) JOIN
    Mengembalikan semua catatan ketika ada kecocokan di tabel kiri atau kanan
    ![alt text](./screenshots/full_join.png "full_join")
    ```sql
    SELECT column_name(s)
    FROM table1
    FULL OUTER JOIN table2
    ON table1.column_name = table2.column_name
    WHERE condition;
    ```

### Union

UNION digunakan untuk menggabungkan kumpulan hasil dari dua atau lebih pernyataan SELECT.

Berikut karakterisitik UNION:
1. Setiap SELECT dalam UNION harus memiliki jumlah kolom yang sama
2. Kolom juga harus memiliki tipe data yang serupa
3. Kolom dalam setiap pernyataan SELECT juga harus dalam urutan yang sama

Sintaks UNION
```sql
SELECT column_name(s) FROM table1
UNION
SELECT column_name(s) FROM table2;
```

Operator UNION hanya memilih nilai yang berbeda secara default. Untuk mengizinkan nilai duplikat, gunakan UNION ALL
```sql
SELECT column_name(s) FROM table1
UNION ALL
SELECT column_name(s) FROM table2;
```

### Function

FUNCTION digunakan untuk membuat fungsi tersimpan dan fungsi yang ditentukan pengguna. Untuk isi FUNCTION sendiri adalah seperangkat pernyataan SQL yang melakukan beberapa operasi dan mengembalikan nilai tunggal.

Sintaks FUNCTION
```sql
DELIMITER $$

CREATE FUNCTION function_name(
    param1,
    param2,…
)
RETURNS datatype
[NOT] DETERMINISTIC
BEGIN
 -- statements
END $$

DELIMITER ;
```