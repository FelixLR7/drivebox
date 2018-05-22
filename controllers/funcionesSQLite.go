package controllers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

const (
	DB_NAME = "sqlite3"
	DB_HOST = "database/BBDD.db"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// createDirIfNotExist ...
func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

// HashPassword ...
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash ...
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Insertar Usuario ...
func InsertarUsuario(email, pass string) {
	db, _ := sql.Open(DB_NAME, DB_HOST)
	stmt, err := db.Prepare("INSERT INTO users (email, password) values(?,?)")
	checkErr(err)

	hash, _ := HashPassword(pass)
	stmt.Exec(email, hash)
	checkErr(err)
	defer db.Close()

	createDirIfNotExist("files/" + email)
}

// Eliminar Usuario ...
func EliminarUsuario(email string) {
	db, _ := sql.Open(DB_NAME, DB_HOST)
	stmt, err := db.Prepare("DELETE FROM users WHERE email = ?")
	checkErr(err)

	stmt.Exec(email)
	checkErr(err)
}

// Listar Usuarios ...
func ListarUsuarios() {
	database, _ := sql.Open(DB_NAME, DB_HOST)

	rows, _ := database.Query("SELECT * FROM users")
	var email string
	var password string
	for rows.Next() {
		rows.Scan(&email, &password)
		fmt.Println(email + " - " + password)
	}
}

// Insertar Archivos ...
func InsertarArchivo(nombre, email string) {
	url := "files/" + email + "/" + nombre

	db, _ := sql.Open(DB_NAME, DB_HOST)
	stmt, err := db.Prepare("INSERT INTO archivos (nombre, url, emailuser) values(?,?,?)")
	checkErr(err)

	stmt.Exec(nombre, url, email)
	checkErr(err)
}

// Eliminar Archivo ...
func EliminarArchivo(archivo, email string) {
	db, _ := sql.Open(DB_NAME, DB_HOST)
	stmt, err := db.Prepare("DELETE FROM archivos WHERE nombre = ? and emailuser = ?")
	checkErr(err)

	stmt.Exec(archivo, email)
	checkErr(err)
}

// Listar Archivos ...
func ListarArchivos(useremail string) {
	database, _ := sql.Open(DB_NAME, DB_HOST)

	rows, _ := database.Query("SELECT * FROM archivos WHERE emailuser = '" + useremail + "';")
	var nombre string
	var url string
	var email string
	for rows.Next() {
		rows.Scan(&nombre, &url, &email)
		fmt.Println(nombre + " - " + url + " - " + email)
	}
}
