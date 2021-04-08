package main

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
	"fmt"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func main() {
	plainText := []byte("My Name Is Yuki Yada")

	keyText := "astaxie12798akljzmknm.ahkjkljl;k" // aes key

	c, err := aes.NewCipher([]byte(keyText))
	if err != nil {
		log.Fatal("Error: NewCipher(%d bytes) = %s", len(keyText), err)
	}

	cfb := cipher.NewCFBEncrypter(c,commonIV)
	cipherText := make([]byte,len(plainText))
	cfb.XORKeyStream(cipherText,plainText)
	fmt.Printf("%s => %x \n",plainText,cipherText)

	cfbdec := cipher.NewCFBDecrypter(c,commonIV)
	plainTextCopy := make([]byte,len(plainText))
	cfbdec.XORKeyStream(plainTextCopy,cipherText)
	fmt.Printf("%x => %s \n",cipherText,plainTextCopy)
}
