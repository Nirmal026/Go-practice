package main

import (
	"bufio"
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func main() {
	// key := []byte("")
	key := uuid.New().String()
	fmt.Print("Input a text : ", key)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	alpabets := scanner.Text()
	c := EncryptAES([]byte(key), alpabets)
	fmt.Println(alpabets)
	fmt.Println(c)
	DecryptAES([]byte(key), c)
}
func EncryptAES(key []byte, plaintext string) string {
	c, err := aes.NewCipher(key)
	CheckError(err)
	out := make([]byte, len(plaintext))
	c.Encrypt(out, []byte(plaintext))
	return hex.EncodeToString(out)
}
func DecryptAES(key []byte, ct string) {
	ciphertext, _ := hex.DecodeString(ct)
	c, err := aes.NewCipher(key)
	CheckError(err)
	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)
	s := string(pt[:])
	fmt.Println("DECRYPTED:", s)
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
