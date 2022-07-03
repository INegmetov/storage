package main

import (
	"fmt"
	"log"

	"github.com/inegmetov/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello"))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello wold!", file)
}
