package hashing

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type ByteData struct {
	Path  string
	Bytes string
}

func GetBytes(path string, group *sync.WaitGroup, sendTo chan ByteData) {
	defer group.Done()
	fmt.Println(path)
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Failed to read file data from: ", path)
	}
	sendTo <- ByteData{Path: path, Bytes: string(data)}
}
