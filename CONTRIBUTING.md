# Contributing

## Daftar Isi
- [Contributor](#contributor)
- [Kategori yang tersedia](#kategori-yang-tersedia)
  - [Basic](#pembelajaran-dasar-basic-learning)
  - [Algoritma](#algoritma-algorithm)
  - [Utility](#komponen-fungsi-kegunaan-utilities)
- [Standar Penulisan](#standar-penulisan)
  - [Tentang gaya PSR-12](#gaya-penulisan-psr12)
  - [Pedoman penulisan untuk basic](#pedoman-penulisan-lainnya-untuk-pembelajaran-dasar-basic-learning)
  - [Pedoman penulisan untuk utility](#pedoman-penulisan-lainnya-untuk-fungsi-kegunaan-utilities)
  - [Pedoman penulisan untuk algorithm](#pedoman-penulisan-lainnya-untuk-algoritma-algorithm)
- [Unit Testing](#testing)
- [Pull Request](#pull-request)

## Contributor

Kami sangat senang Anda dapat ikut berkontribusi dalam implementasi algoritma, struktur data atau memperbaiki error serta merilis implementasi materi basic learning Golang. Semua boleh ikut berkontribusi walaupun hal kecil dengan ketentuan sebagai berikut:

- Hasil pekerjaan Anda adalah buatan Anda sendiri dan tidak ada hak cipta dari orang lain, jika kami menemukan kesamaan maka kami tidak `merged`.
- Hasil kerja Anda akan berlisensi [MIT](LICENSE) ketika permintaan pull Anda sudah di merged
- Hasil kerja Anda wajib mengikuti standar dan style koding dari kami (lihat bagian [Standar penulisan](#standar-penulisan))
- Jika dibutuhkan, hanya gunakan library package dari composer, sebelumnya silahkan diskusi di [issue](https://github.com/bellshade/Golang/issues)

## Kategori yang tersedia

### Pembelajaran Dasar (_Basic Learning_)

Basic learning adalah kumpulan demonstrasi kode materi pemrograman Go beserta artikel penjelasan yang lengkap, mudah di mengerti dan menarik secara visual. Dirangkai dengan alur yang runtut mulai dasar sehingga bisa dijadikan referensi untuk kalian yang belajar pemrograman Go.

Untuk berkontribusi di repositori bellshade Golang bagian Basic ini kamu harus mengikuti urutan materi yang telah ditentukan. Daftar materi yang tersedia bisa kamu jelajahi di [Issues tag materi](https://github.com/bellshade/Golang/issues). Didalam issue itu kami juga menuliskan rekomendasi tentang konten yang ada dalam materi tersebut, namun kamu masih dapat menambah materi konten jika dirasa perlu. 

Semua implementasi kode untuk Basic Learning bisa ditemukan di `/basics`  <br/>
dengan struktur direktori sebagai berikut <br/>

```
📦 /
 ┣📦basics
 ┃ ┣ 📂x_nama_materi
 ┃ ┃ ┣ 📜1_sub_materi.go
 ┃ ┃ ┣ 📜2_sub_materi.go
 ┃ ┃ ┣ 📜3_sub_materi.go
 ┃ ┃ ┣  ....
 ┃ ┃ ┗ 📜README.md
```

> Artikel materi ditulis pada `README.md` <br/>


### Algoritma (_Algorithm_)

Algoritma adalah langkah-langkah untuk menyelesaikan suatu pekerjaan dimana terdiri dari 3 bagian utama, yaitu:

- Input/masukan, sebelum menjalankan sebuah algoritma maka hal yang pertama harus dilakukan adalah menerima masukan, input dapat berasal dari pengguna ataupun dari langkah sebelumnya.
- Proses, bagian utama dari algoritma yang melakukan pengolahan input yang akan menghasilkan output.
- Output/keluaran, output adalah hasil dari bagian proses, output ini juga bisa digunakan untuk langkah selanjutnya (jika masih ada).

Algoritma harus dikemas sedemikian rupa sehingga memudahkan pembaca untuk memasukkannya ke dalam program yang lebih besar.

Semua implementasi kode untuk Algorithm bisa ditemukan di `/algorithms` dengan bentuk struktur seperti berikut

```
📦/
 ┣📦algorithms
 ┃ ┣ 📂NamaKategoriAlgoritma
 ┃ ┃ ┣ 📂NamaKasusAlgoritma
 ┃ ┃ ┃ ┣ 📜FileAlgoritmanya.go
 ┃ ┃ ┃ ┗ 📜README.md
 ┃ ┃ ┗ 📜README.md
```
> Artikel penjelasan algoritma ditulis pada `README.md` <br/>

### Komponen / Fungsi kegunaan _(Utilities)_
Kategori utilities adalah kumpulan referensi dan contoh penggunaan komponen-komponen / fitur-fitur umum yang mungkin diperlukan oleh para web developer untuk diterapkan pada project aplikasi webnya. Didalam kategori ini, kalian dapat menjelajahi berbagai macam kebutuhan umum untuk pengembangan web umum dan mempelajari bagaimana cara mengimplementasikannya.

Utility ini adalah sesuatu yang bukan esensial dalam pengembangan web namun terkadang dibutuhkan. Bisa juga dibilang Utility ini adalah peralatan / komponen yang bisa melengkapi sebuah fitur dalam aplikasi berbasis web. Contohnya adalah seperti: Manipulasi Gambar, DateTime, Mencetak ke printer dot matrix dan sebagainya. 

Semua tentang Utility dapat ditemukan pada direktori `/utilities`<br/>
dengan bentuk struktur yang mirip dengan `algorithm`

```
📦/
 ┣📦utilities
 ┃ ┣ 📂NamaKategoriUtility
 ┃ ┃ ┣ 📂NamaUtility
 ┃ ┃ ┃ ┣ 📜demo_utility_1.go
 ┃ ┃ ┃ ┣ 📜demo_utility_2.go
 ┃ ┃ ┃ ┣ 📜 .....
 ┃ ┃ ┃ ┗ 📜README.md
 ┃ ┃ ┗ 📜README.md
 ```
> Artikel penjelasan utility ditulis pada `README.md` <br/>

## Standar Penulisan

### Gaya Penulisan 


### Pedoman Penulisan Lainnya untuk Pembelajaran Dasar (_Basic Learning_)

Selain menggunakan gaya penulisan PSR-12 untuk implementasi kode pada Basic Learning. Ada ketentuan konsistensi yang harus Anda terapkan yaitu:

- Setiap materi pembelajaran berada pada folder tersendiri yang dinamai dengan materi yang berkaitan
- Folder materi memiliki penomoran diawal sebagai _prefix_
- Artikel penjelasan ditulis di `README.md` pada setiap folder materi. jadi kamu perlu membuat `README.md`nya
- Penamaan variabel dan fungsi menggunakan `snake_case`

### Pedoman Penulisan Lainnya untuk fungsi kegunaan _(Utilities)_

Pedoman penulisan untuk utilities sama persis seperti materi pembelajaran dasar _(basic)_ dengan pengecualian, penamaan variabel bebas asalkan tetap konsisten diseluruh satu _utility_.

### Pedoman Penulisan Lainnya untuk Algoritma (_Algorithm_)



## Percobaan (_Testing_)

Testing hanya dilakukan untuk Algorithm dan tidak diperlukan untuk Basic Learning ataupun Utility. Testing algoritma diperlukan agar suatu algoritma berjalan sesuai dengan yang di harapkan. Kami juga menerapkan unit testing ini pada _Continuous Integration_ kami agar selalu dijaga keandalan kodenya.


## Pull Request

**_Pull Request_ yang Baik**

- Lakukan fork pada repository kami
- usahakan Anda membuat branch baru

```bash
git checkout -b <branch_name>
git add .
git commit -m "add: menambahkan algoritma baru"
```

- Lakukan push ke branch Anda dan kemudian *open *pull request\*\*

**Saran Pesan Commit:**

- `add` Untuk menambah algoritma atau tambahan lainnya
- `fix` Untuk mengubah algoritma yang sudah ada atau memperbaiki
- `docs` Untuk mengubah atau membuat dokumentasi
- `style` Untuk mengubah atau memperbaiki style kode untuk contohnya bisa dilihat pada commit yang diatas

**_Pull Request_ `merged` jika:**

- Mengikuti standar dan arahan dari `CONTRIBUTING.md`
- Lulus test dan cek dari beberapa test yang sudah kami siapkan

**Tambahan:**

- Jika ada kendala atau masalah dalam _pull request_, Anda bisa laporkan masalah pada [issue](https://github.com/bellshade/Golang/issues)
- Jika ada tes yang tidak lewat atau gagal, kami akan mengecek kembali perubahan.

Untuk _pull request_ kami sarankan untuk menjelaskan secara detail yang Anda ubah atau tambahkan, dan bersikap sopan, serta selalu berterima kasih, itu salah satu bentuk tata krama yang baik terhadap sesama contributor dan programmer lainnya.

Terima kasih sudah berkontribusi di **Bellshade/Golang**
