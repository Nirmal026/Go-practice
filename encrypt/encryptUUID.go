package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pborman/uuid"
)

func generator() ([]byte, error) {
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Print(err)
	}
	return nonce, nil
}

func encrypt(key string, text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Print(err)
	}
	plaintext := []byte(text)
	nonce, err := generator()
	if err != nil {
		fmt.Print(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Print(err)
	}
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

func decrypt(key string, text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Print(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
	}
	in, _ := hex.DecodeString(text)
	size := gcm.NonceSize()
	nonce, ciphertext := in[:size], in[size:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Print(err)
	}
	return string(plaintext[:]), nil
}

func main() {
	fmt.Print("Input a text : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	alpabets := scanner.Text()
	uuid1 := uuid.NewRandom()
	fmt.Print(uuid1)
	key := strings.Replace(uuid1.String(), "-", "", -1)
	ciphertext, _ := encrypt(key, alpabets)
	fmt.Println(ciphertext)
	plaintext, _ := decrypt(key, ciphertext)
	fmt.Printf("Decrypt: %s\n", plaintext)
}
