package main

import (
	"encryption"
	"fmt"

	"golang.org/x/crypto/argon2"
)


func main() {
	fmt.Println("Encrypting")

	password := []byte("some password")
	salt  := []byte("somesalt")

	time := uint32(1000)
	memory := uint32( 1*1024)
	threads := uint8(1)
	keyLen := uint32(32)

	key := argon2.IDKey(password, salt, time, memory, threads, keyLen)

	message := "super secret message"

	encryptedMessage, err := encryption.EncryptMessage(key, message)
	if err != nil {
		fmt.Errorf("encryption error: %v", err)
	}

	decryptedMessage, err := encryption.DecryptMessage(key, encryptedMessage)
	if err != nil {
		fmt.Errorf("decryption error: %v", err)
	}

	fmt.Printf("Message: %v\n", message)
	fmt.Printf("Encrypted Message: %v\n", encryptedMessage)
	fmt.Printf("Decrypted Message: %v\n", decryptedMessage)
}
