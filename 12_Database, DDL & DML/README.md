# Database, DDL & DML

## Resume Materi

### Database

Database  adalah kumpulan data yang dikelola sedemikian rupa berdasarkan ketentuan tertentu yang saling berhubungan sehingga mudah dalam pengelolaannya.

Database relasional didasarkan pada model relasional, cara yang intuitif dan mudah untuk merepresentasikan data dalam tabel.

Berikut jenis-jenis relasi dalam database:
1. One to One
    Membuat hubungan antara dua tabel dimana satu baris dari tabel pertama hanya dapat dikaitkan dengan satu dan hanya satu catatan dari tabel kedua
2. One to Many
    Membuat hubungan antara dua tabel dimana setiap baris dari tabel pertama dapat dikaitkan dengan satu atau lebih baris dari tabel kedua, tetapi baris tabel kedua hanya dapat berhubungan dengan satu-satunya baris di tabel pertama
3. Many to Many
    Membuat hubungan antara dua tabel dimana setiap record dari tabel pertama dapat berhubungan dengan record apapun (atau tidak ada record) di tabel kedua. Demikian pula, setiap record dari tabel kedua juga dapat berhubungan dengan lebih dari satu record dari tabel pertama

### Data Manipulation Language(DML)

DML adalah perintah yang digunakan untuk memanipulasi data dalam tabel dari suatu database.

Berikut perintah-perintah yang dapat digunakan:
1. Insert
    Untuk input data ke table dengan sintaks sebagai berikut
    ```sql
    -- with column table
    INSERT INTO table_name (column1, column2, column3, ...)
    VALUES (value1, value2, value3, ...);

    -- without column table but the column must sequential
    INSERT INTO table_name
    VALUES (value1, value2, value3, ...);
    ```
2. Select
    Untuk menampilkan data pada tabel dengan sintaks sebagai berikut
    ```sql
    SELECT column1, column2, ... FROM table_name;
    ```
3. Update
    Untuk mengubah data pada tabel dengan sintaks sebagai berikut
    ```sql
    UPDATE table_name
    SET column1 = value1, column2 = value2, ...
    WHERE condition;
    ```
4. Delete
    Untuk menghapus data pada tabel dengan sintaks sebagai berikut
    ```sql
    DELETE FROM table_name WHERE condition;
    ```

### DML Statement

1. Like/Between
    Operator LIKE digunakan dalam klausa WHERE untuk mencari pola tertentu dalam kolom.
    ```sql
    SELECT column1, column2, ...
    FROM table_name
    WHERE columnN LIKE pattern;
    ```
2. And/Or
    Operator AND dan OR digunakan untuk memfilter record berdasarkan kondisi yang pada operator yang dipilih
    ```sql
    -- and
    SELECT column1, column2, ...
    FROM table_name
    WHERE condition1 AND condition2 AND condition3 ...;

    -- or
    SELECT column1, column2, ...
    FROM table_name
    WHERE condition1 OR condition2 OR condition3 ...;
    ```
3. Order By
    ORDER BY digunakan untuk mengurutkan kumpulan hasil dalam urutan menaik atau menurun. ORDER BY mengurutkan catatan dalam urutan menaik secara default. Untuk mengurutkan catatan dalam urutan menurun, gunakan DESC.
    ```sql
    SELECT column1, column2, ...
    FROM table_name
    ORDER BY column1, column2, ... ASC|DESC;
    ```
4. Limit
    LIMIT digunkan untuk membatasi jumlah row table yang akan ditampilkan
    ```sql
    SELECT column_name(s)
    FROM table_name
    WHERE condition
    LIMIT number;
    ```