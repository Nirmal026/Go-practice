package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Input a text : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	alpabets := scanner.Text()
	h := sha256.New()
	h.Write([]byte(alpabets))
	bs := h.Sum(nil)
	fmt.Println(alpabets)
	fmt.Printf("%x\n", bs)
	// rawDecodedText, err := base64.StdEncoding.DecodeString(string(bs))
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Decoded text: %s\n", rawDecodedText)
}
