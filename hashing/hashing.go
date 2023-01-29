package hashing

import (
	"crypto"
	_ "crypto/sha256"
	"fmt"
	"log"
	"os"
	"sync"
)

type ByteData struct {
	Path  string
	Bytes string
}

func GetBytesHashed(path string, group *sync.WaitGroup, sendTo chan ByteData) {
	defer group.Done()
	fmt.Println(path)
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Failed to read file data from: ", path)
	}
	hashed := HashSha256(data)
	sendTo <- ByteData{Path: path, Bytes: hashed}
}

func HashSha256(data []byte) string {
	hash := crypto.SHA256.New()
	hash.Write(data)
	finalHash := hash.Sum(nil)
	return string(finalHash)
}
