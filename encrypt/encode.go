package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	// String to encode
	fmt.Print("Input a text : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	alpabets := scanner.Text()
	encodedStr := base64.StdEncoding.EncodeToString([]byte(alpabets))
	fmt.Println("Encoded:", encodedStr)
	decodedStr, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decoded:", string(decodedStr))
}
