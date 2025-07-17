## database/sql

`database/sql` adalah interface generic untuk bekerja dengan database yang bersifat relasional di golang. Ini bukan driver database itu sendiri, melainkan sebuah standard interface yang bisa digunakan dengan driver-driver spesifik database.

## Arsitektur sederhana

```
+-----------------+
|     Aplikasi    |
|      Go App     |
+--------+--------+
         |
         | SQL Query
         v
+--------+--------+
| database/sql    | <== abstraction layer
+--------+--------+
         |
         +------------+---------------+
                      |               |
        +-------------v----+   +------v-------+
        | Driver MySQL      |   | Driver PG    |
        +-------------------+   +--------------+
```
- kita tidak perlu lagi menulis ulang logika koneksi, pooling, query.
- kode kita bisa berjalan di berbagai jenis database selama ada drivernya.
- dan manajemen koneksi yang aman efisien.
- serta mendukung prepared statement untuk mencegah terjadinya SQL injection

## Yang harus diperhatikan

Ketika bekerja dengan sql statement dan database, usahakan untuk
- gunakana prepared statement untuk menghindari terjadinya SQL injection.
- jangan lupa `defer baris.Close()` / `defer row.Close()` saat menggunakan query.
- logging error secara rinci untuk debug di environment production

untuk contohnya bisa lihat filenya [disini](databaseSql.go)
