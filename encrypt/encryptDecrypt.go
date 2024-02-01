package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

// main function
func main() {
	key := []byte("")
	fmt.Print("Input a text : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	alpabets := scanner.Text()
	text, err := Encrypt([]byte(alpabets), key)
	if err != nil {
		fmt.Println(err)
	}
	// printing encrypted text
	fmt.Println("Encrypted : ", base64.StdEncoding.EncodeToString(text))
	plainText, err := Decrypt(text, key)
	if err != nil {
		fmt.Println(err)
	}
	// printing decrypted text
	fmt.Println("Decrypted : ", string(plainText))
}

// function Encrypt for encryption
func Encrypt(plaintext []byte, key []byte) (ciphertext []byte, err error) {
	data := sha256.Sum256(key)
	block, err := aes.NewCipher(data[:])
	if err != nil {
		fmt.Println(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		fmt.Println(err)
	}
	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// function Decrypt for decryption
func Decrypt(ciphertext []byte, key []byte) (plaintext []byte, err error) {
	data := sha256.Sum256(key)
	block, err := aes.NewCipher(data[:])
	if err != nil {
		fmt.Println(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
	}
	if len(ciphertext) < gcm.NonceSize() {
		fmt.Println("Error found")
	}
	return gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
}
