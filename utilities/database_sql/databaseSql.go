package databasesql

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	id    int
	nama  string
	email string
}

func buatTabel(db *sql.DB) {
	queryDatabaseTabel := `
	CREATE TABLE IF NOT EXISTS users (
	    id INT AUTO_INCREMENT PRIMARY KEY,
	    nama VARCHAR(50),
	    email VARCHAR(100)
	);
	`

	_, err := db.Exec(queryDatabaseTabel)
	if err != nil {
		log.Fatal("error saat membuat tabel, informasi: ", err)
	}

	fmt.Println("tabel `users` sudah dibuat atau sudah ada")
}

func insertDataUser(db *sql.DB, nama, email string) {
	queryData := "INSERT INTO users (nama, email) VALUES (?, ?)"
	hasil, err := db.Exec(queryData, nama, email)
	if err != nil {
		log.Fatal("error pada saat insert data user, informasi: ", err)
	}

	id, _ := hasil.LastInsertId()
	fmt.Printf("insert data dengan id: %d\n", id)
}

func ambilSemuaDataUser(db *sql.DB) []User {
	baris, err := db.Query("SELECT id, nama, email FROM users")
	if err != nil {
		log.Fatal("error fetching data user, informasi: ", err)
	}

	defer baris.Close()

	var userData []User

	for baris.Next() {
		var u User
		err := baris.Scan(&u.id, &u.nama, &u.email)
		if err != nil {
			log.Fatal("error scanning dari baris, informasi: ", err)
		}
		userData = append(userData, u)
	}
	return userData
}

func updateUserEmail(db *sql.DB, id int, emailBaru string) {
	queryData := "UPDATE users SET email = ? WHERE id = ?"
	_, err := db.Exec(queryData, emailBaru, id)
	if err != nil {
		log.Fatal("error update user, informasi: ", err)
	}
	fmt.Printf("update email untuk user id %d\n", id)
}

func deleteUser(db *sql.DB, id int) {
	query := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal("error delete data user, informasi: ", err)
	}
	fmt.Printf("menghapus data user dengan id: %d\n", id)
}

func main() {
	// format data source name atau DSN
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal("error waktu buka database, informasi: ", err)
	}

	defer db.Close()

	// ping database untuk pastikan koneksi berhasil
	err = db.Ping()
	if err != nil {
		log.Fatal("database ada masalah, informasi: ", err)
	}
	fmt.Println("koneksi ke database berjalan dengan baik")

	// membuat tabel jika tidak ada
	buatTabel(db)
	insertDataUser(db, "Mamad", "mamad@email.com")
	insertDataUser(db, "hiruzen", "beban@email.com")

	// ambil semua data user dan tampilin
	userData := ambilSemuaDataUser(db)
	for _, user := range userData {
		fmt.Printf("ID: %d, Nama: %s, Email: %s\n", user.id, user.nama, user.email)
	}

	// update data user
	updateUserEmail(db, 1, "nongpal@email.com")
	// delete data user
	deleteUser(db, 2)
}
