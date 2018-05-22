package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
)

// Decrypt
func decrypt(cipherstring string, keystring string) string {
	ciphertext := []byte(cipherstring)
	key := []byte(keystring)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("Text is too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext)
}

// Encrypt
func encrypt(plainstring, keystring string) string {
	plaintext := []byte(plainstring)
	key := []byte(keystring)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return string(ciphertext)
}

// WriteToFile
func writeToFile(data, file string) {
	ioutil.WriteFile(file, []byte(data), 777)
}

// ReadFromFile
func readFromFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	return data, err
}

// Cifrar archivo ...
func CifrarArchivo(file, key string) {
	content, err := readFromFile(file)
	if err != nil {
		panic(err)
	}
	encrypted := encrypt(string(content), key)
	writeToFile(encrypted, file+".enc")
}

// Descifrar archivo ...
func DescifrarArchivo(file, key string) {
	content, err := readFromFile(file)
	if err != nil {
		panic(err)
	}
	decrypted := decrypt(string(content), key)
	writeToFile(decrypted, file[:len(file)-4])
}
